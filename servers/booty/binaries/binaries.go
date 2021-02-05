// binaries fetches "binaries" from URLs, decompresses/extracts them if the
// URL points to an archive like a zip, and provides a path to the binary
// named after the package.
// The binary should be either the last element of the URL path for non-archives,
// or live somewhere within the archive.
package binaries

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"text/template"

	"github.com/hashicorp/go-getter"
)

type URLHelper struct {
	urlTemplate    *template.Template
	urlSumTemplate *template.Template
}

func NewURLHelper(mainURL, sumURL string) (*URLHelper, error) {
	var err error
	u := &URLHelper{
		urlTemplate:    template.New("url"),
		urlSumTemplate: template.New("sumurl"),
	}

	u.urlTemplate, err = u.urlTemplate.Parse(mainURL)
	if err != nil {
		return nil, fmt.Errorf("bad mainURL template: %v", err)
	}

	u.urlSumTemplate, err = u.urlSumTemplate.Parse(sumURL)
	if err != nil {
		return nil, fmt.Errorf("bad sumURL template: %v", err)
	}

	return u, nil
}

const hashicorpURLTemplateBase = "https://releases.hashicorp.com/{{ .Package }}/{{ .Version }}/"
const hashicorpURLTemplate = hashicorpURLTemplateBase + "{{ .Package }}_{{ .Version }}_{{ .OS }}_{{ .Arch }}.zip"
const hashicorpURLSumTemplate = hashicorpURLTemplateBase + "{{ .Package }}_{{ .Version }}_SHA256SUMS"
const prometheusURLTemplateBase = "https://github.com/prometheus/{{ .Package }}/releases/download/v{{ .Version }}/"
const prometheusURLTemplate = prometheusURLTemplateBase + "{{ .Package }}-{{ .Version }}.{{ .OS }}-{{ .Arch }}.tar.gz"
const prometheusURLSumTemplate = prometheusURLTemplateBase + "sha256sums.txt"

var hashicorpURLHelper, prometheusURLHelper *URLHelper

var Default Manager

func init() {
	u, err := NewURLHelper(hashicorpURLTemplate, hashicorpURLSumTemplate)
	if err != nil {
		panic(err.Error())
	}
	hashicorpURLHelper = u

	u, err = NewURLHelper(prometheusURLTemplate, prometheusURLSumTemplate)
	if err != nil {
		panic(err.Error())
	}
	prometheusURLHelper = u

	Default, err = NewDownloadManager(filepath.Join(os.TempDir(), "yurt/binaries"))
	if err != nil {
		log.Fatal(err)
	}
}

type registryEntry struct {
	name    string
	version string
	from    *URLHelper
}

func registry() map[string]registryEntry {
	// yurt-run is "fetchable", but it will be built locally,
	// so we don't have a registry entry for it
	return map[string]registryEntry{
		"nomad": {
			name:    "nomad",
			version: "0.12.5",
			from:    hashicorpURLHelper,
		},
		"consul": {
			name:    "consul",
			version: "1.8.4",
			from:    hashicorpURLHelper,
		},
		"vault": {
			name:    "vault",
			version: "1.5.4",
			from:    hashicorpURLHelper,
		},
		"prometheus": {
			name:    "prometheus",
			version: "2.22.0",
			from:    prometheusURLHelper,
		},
		"consul_exporter": {
			name:    "consul_exporter",
			version: "0.7.1",
			from:    prometheusURLHelper,
		},
		"node_exporter": {
			name:    "node_exporter",
			version: "1.0.1",
			from:    prometheusURLHelper,
		},
	}
}

type Manager interface {
	Get(packageName string) (string, error)
	GetOSArch(packageName, os, arch, version string) (string, error)
}

type EnvPathManager struct {
}

func (e EnvPathManager) Get(packageName string) (string, error) {
	return exec.LookPath(packageName)
}

func (e EnvPathManager) GetOSArch(packageName, os, arch, version string) (string, error) {
	return "", fmt.Errorf("GetOSArch not implemented for path-based binary manager")
}

var _ Manager = EnvPathManager{}

type DownloadManager struct {
	l       sync.Mutex
	cache   map[string]string
	workDir string
}

var _ Manager = &DownloadManager{}

func NewDownloadManager(workDir string) (*DownloadManager, error) {
	m := &DownloadManager{
		workDir: workDir,
		cache:   make(map[string]string),
	}
	if err := os.MkdirAll(m.workDir, 0755); err != nil {
		return nil, err
	}
	return m, nil
}

// dldirToBinary takes as input dldir, a directory that go-getter wrote to,
// and the name of an upstream package defined in the registry package variable
// (e.g. "consul").
// Returns the path to an executable file that matches packageName
// found somewhere under dldir.
func dldirToBinary(dldir, packageName string) (string, error) {
	fis, err := ioutil.ReadDir(dldir)
	if err != nil {
		return "", err
	}

	// There can be as many leading nested directories as you like.
	if len(fis) == 1 && fis[0].IsDir() {
		return dldirToBinary(filepath.Join(dldir, fis[0].Name()), packageName)
	}

	// As soon as we find a non-directory or multiple files in a directory,
	// we better find an executable file matching packageName in this directory
	// or there's going to be an error.
	for _, fi := range fis {
		if fi.Name() == packageName && fi.Mode().IsRegular() && (fi.Mode()&0111) == 0111 {
			return filepath.Join(dldir, packageName), nil
		}
	}

	return "", fmt.Errorf("didn't find %s under %s", packageName, dldir)
}

func (m *DownloadManager) Get(packageName string) (string, error) {
	return m.GetOSArch(packageName, runtime.GOOS, runtime.GOARCH, "")
}

func (m *DownloadManager) GetOSArch(packageName, os, arch, version string) (string, error) {
	m.l.Lock()
	defer m.l.Unlock()

	if binPath, ok := m.cache[packageName+":"+version]; ok {
		return binPath, nil
	}

	var binPath string
	var err error

	if packageName == "yurt-run" {
		binPath, err = m.buildLocalBin(packageName, os, arch)
	} else {
		binPath, err = m.Fetch(packageName, os, arch, version)
	}

	if err != nil {
		return "", err
	}
	m.cache[packageName+":"+version] = binPath
	return binPath, nil
}

// Fetch fetches the packageName based on its registry entry,
// if it's not already present on disk with the correct checksum.
// Then it extracts the archive and finds the binary with the same name as packageName.
// Returns the absolute path (located under m.workDir) where the binary was found.
func (m *DownloadManager) Fetch(packageName, osName, arch, version string) (string, error) {
	workdir := m.workDir
	o, ok := registry()[packageName]
	if !ok {
		return "", fmt.Errorf("unknown package name %q", packageName)
	}

	if version == "" {
		version = o.version
	}

	var sumURL bytes.Buffer
	err := o.from.urlSumTemplate.Execute(&sumURL, struct {
		Package string
		Version string
	}{
		Package: packageName,
		Version: version,
	})
	if err != nil {
		return "", err
	}

	var sourceURL bytes.Buffer
	err = o.from.urlTemplate.Execute(&sourceURL, struct {
		Package string
		Version string
		OS      string
		Arch    string
	}{
		Package: packageName,
		Version: version,
		OS:      osName,
		Arch:    arch,
	})
	if err != nil {
		return "", err
	}

	sourceURLParsed, err := url.Parse(sourceURL.String())
	if err != nil {
		return "", err
	}

	localPackage := filepath.Join(workdir, packageName, filepath.Base(sourceURLParsed.Path))
	beforeStat, err := os.Stat(localPackage)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	// First download the package archive file, using a checksum URL to validate
	// its contents.  This also allows us to skip the download if the file
	// already exists with the valid checksum.
	client := &getter.Client{
		Src:           sourceURL.String() + "?checksum=file:" + sumURL.String(),
		Dst:           localPackage,
		Mode:          getter.ClientModeFile,
		Decompressors: map[string]getter.Decompressor{},
	}
	if err := client.Get(); err != nil {
		return "", fmt.Errorf("go-getter error: %w", err)
	}
	afterStat, err := os.Stat(localPackage)
	if err != nil {
		return "", err
	}

	// Now that we know the correct file exists on disk, we could just extract it.
	// But we could also allow go-getter to do the work of figuring out how.
	// Unless of course the archive file is unchanged and we see an existing
	// extract dir, in which case we do nothing.
	packageExtract := filepath.Join(workdir, packageName, version)
	_, err = os.Stat(localPackage)
	_, err2 := os.Stat(packageExtract)
	if err == nil && err2 == nil && beforeStat != nil && beforeStat.ModTime().Equal(afterStat.ModTime()) {
		return dldirToBinary(packageExtract, packageName)
	}

	// If we reached this point we might have re-downloaded something due to a
	// checksum issue, so don't use any files that might have come from prev ver.
	if err = os.RemoveAll(packageExtract); err != nil {
		return "", err
	}

	// Extract to a temp folder so that we can detect previous errors, i.e. partial extracts
	packageExtractTmp := packageExtract + ".tmp"
	if err = os.RemoveAll(packageExtractTmp); err != nil {
		return "", err
	}

	client = &getter.Client{
		Src:  localPackage,
		Dst:  packageExtractTmp,
		Mode: getter.ClientModeDir,
	}
	if err := client.Get(); err != nil {
		return "", fmt.Errorf("go-getter error: %w", err)
	}
	err = os.Rename(packageExtractTmp, packageExtract)

	return dldirToBinary(packageExtract, packageName)
}

// Work upwards through the directory tree starting at the current directory,
// stopping when a directory named ".git" and a file named "go.mod" exists.
// Intended for use in tests.
func (m *DownloadManager) projectRoot() (string, error) {
	isProjectRoot := func(dir string) (bool, error) {
		s, err := os.Stat(filepath.Join(dir, ".git"))
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
		if !s.IsDir() {
			return false, nil
		}

		s, err = os.Stat(filepath.Join(dir, "go.mod"))
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
		if s.IsDir() {
			return false, nil
		}
		return true, nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for dir[len(dir)-1] != filepath.Separator {
		found, err := isProjectRoot(dir)
		if err != nil {
			return "", err
		}
		if found {
			return dir, nil
		}

		dir = filepath.Dir(dir)
	}

	return "", fmt.Errorf("no project root found (looked for .git dir and go.mod file)")
}

func (m *DownloadManager) buildLocalBin(name, osname, arch string) (string, error) {
	proot, err := m.projectRoot()
	if err != nil {
		return "", err
	}

	fqfn, err := exec.LookPath("go")
	if err != nil {
		return "", fmt.Errorf("can't find 'go' in path: %v", err)
	}

	dest := filepath.Join(m.workDir, name)
	cmd := exec.Command(fqfn, "build", "-o", dest)
	cmd.Dir = filepath.Join(proot, "cmd", name)
	for _, e := range os.Environ() {
		switch {
		case strings.HasPrefix(e, "GOOS="):
		case strings.HasPrefix(e, "GOARCH="):
		case strings.HasPrefix(e, "GO111MODULE="):
		case strings.HasPrefix(e, "GOFLAGS="):
		case strings.HasPrefix(e, "CGO_ENABLED="):
		default:
			cmd.Env = append(cmd.Env, e)
		}
	}
	cmd.Env = append(cmd.Env, "GOOS="+osname, "GOARCH="+arch, "GO111MODULE=on", "GOFLAGS=-mod=")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Print("running command:", cmd)

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return dest, nil
}
