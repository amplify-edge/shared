package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ncabatoff/yurt/binaries"
	"github.com/ncabatoff/yurt/consul"
	"github.com/ncabatoff/yurt/nomad"

	"github.com/hashicorp/go-sockaddr"
	"github.com/hashicorp/vault/sdk/helper/certutil"
	"github.com/ncabatoff/yurt"
	"github.com/ncabatoff/yurt/pki"
	"github.com/ncabatoff/yurt/runenv"
	"github.com/ncabatoff/yurt/runner"
	"gopkg.in/yaml.v2"
)

type yurtConfig struct {
	DataDir         string   `yaml:"data_dir,omitempty"`
	TLS             bool     `yaml:"tls,omitempty"`
	NetworkCIDR     string   `yaml:"network_cidr,omitempty"`
	ConsulServerIPs []string `yaml:"consul_server_ips,omitempty"`
	ConsulBin       string   `yaml:"consul_bin,omitempty"`
	NomadBin        string   `yaml:"nomad_bin,omitempty"`
	CACertFile      string   `yaml:"ca_cert_file,omitempty"`
	serverIP        string
	network         sockaddr.SockAddr
	TLSConfig       *pki.TLSConfigPEM
}

func (c *yurtConfig) IsConsulServer() bool {
	for _, ip := range c.ConsulServerIPs {
		if c.serverIP == ip {
			return true
		}
	}
	return false
}

func main() {
	var (
		flagConfigFile  = flag.String("config-file", "", "optional config file")
		flagConsulBin   = flag.String("consul-bin", "", "path to Consul binary, will download if empty")
		flagConsulIPs   = flag.String("consul-server-ips", "", "comma-separated list of consul server IPs")
		flagData        = flag.String("data", "/var/yurt", "directory to store state")
		flagNetworkCIDR = flag.String("network-cidr", "", "network cidr, optional if consul-server-ips are on a /24")
		flagNomadBin    = flag.String("nomad-bin", "", "path to Nomad binary, will download if empty")
		flagTLS         = flag.Bool("tls", false, "enable TLS authentication")
		flagVaultAddr   = flag.String("vault-addr", "", "vault address for TLS cert gen, put token in $VAULT_TOKEN")
		// restart policy
	)
	flag.Parse()
	noArgsGiven := yurtConfig{DataDir: "/var/yurt", ConsulServerIPs: []string{}}
	yc := &yurtConfig{
		ConsulBin:       *flagConsulBin,
		ConsulServerIPs: strings.Split(*flagConsulIPs, ","),
		DataDir:         *flagData,
		NetworkCIDR:     *flagNetworkCIDR,
		NomadBin:        *flagNomadBin,
		TLS:             *flagTLS,
	}

	switch {
	case *flagConfigFile == "":
	case reflect.DeepEqual(*yc, noArgsGiven):
		var err error
		yc, err = loadConfigFile(*flagConfigFile)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("cannot provide other arguments along with -config-file")
	}

	if yc.NetworkCIDR == "" {
		// assume it's a /24 if not specified
		last := strings.LastIndexByte(yc.ConsulServerIPs[0], '.')
		if last == -1 {
			log.Fatalf("bad consul ip: %q", yc.ConsulServerIPs[0])
		}

		yc.NetworkCIDR = yc.ConsulServerIPs[0][:last] + ".0/24"
	}

	if netSA, err := sockaddr.NewSockAddr(yc.NetworkCIDR); err != nil {
		log.Fatalf("bad cidr %q, err=%w", yc.NetworkCIDR, err)
	} else {
		yc.network = netSA
	}

	for _, ip := range yc.ConsulServerIPs {
		ipSA, err := sockaddr.NewSockAddr(ip)
		if err != nil {
			log.Fatalf("bad consul ip %q, err=%w", ip, err)
		}
		if !yc.network.Contains(ipSA) {
			log.Fatalf("consul ip %s is not contained in network %s", ipSA, yc.network)
		}
	}

	ifAddrs, err := sockaddr.GetAllInterfaces()
	if err != nil {
		log.Fatalf("error listing interfaces: %v", err)
	}
	for _, ifAddr := range ifAddrs {
		if yc.network.Contains(ifAddr.SockAddr) {
			yc.serverIP = sockaddr.ToIPv4Addr(ifAddr.SockAddr).NetIP().String()
			log.Print(yc.serverIP)
		}
	}
	if yc.serverIP == "" {
		log.Fatalf("network interface for network_cidr %s not found", yc.NetworkCIDR)
	}

	if *flagVaultAddr != "" {
		if err := yc.setupTLS(*flagVaultAddr, yc.serverIP); err != nil {
			log.Fatal(err)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e, err := runenv.NewExecEnv(ctx, "yurt", *flagData, 16000, binaries.Default)
	if err != nil {
		log.Fatalf("error creating env: %v", err)
	}

	// TODO should we handle restarts, or rely on OS?
	e.Go(runConsul(ctx, e, yc).Wait)
	e.Go(runNomad(ctx, e, yc).Wait)

	if err := e.Wait(); err != nil {
		log.Fatal(err)
	}
}

func loadConfigFile(path string) (*yurtConfig, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var c yurtConfig
	if err := yaml.Unmarshal(contents, &c); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	return &c, nil
}

func (c *yurtConfig) setupTLS(vaultAddr, myIP string) error {
	caFile := c.CACertFile
	if caFile == "" {
		caFile = filepath.Join(c.DataDir, "ca.pem")
	}

	contents, err := ioutil.ReadFile(caFile)
	switch {
	case err == nil:
		_, err := certutil.ParsePEMBundle(string(contents))
		if err == nil {
			c.CACertFile = caFile
			return nil
		}
		log.Printf("error parsing CA file %s: %v", caFile, err)
	case errors.Is(err, os.ErrNotExist):
	default:
		log.Printf("error reading CA file %s: %v", caFile, err)
	}

	ca, err := pki.NewExternalCertificateAuthority(vaultAddr, os.Getenv("VAULT_TOKEN"))
	if err != nil {
		return fmt.Errorf("error setting up external certificate authority: %w", err)
	}

	cert, err := ca.ConsulServerTLS(context.Background(), myIP, "168h")
	if err != nil {
		return fmt.Errorf("error generating Consul server certificate for ip=%v: %w", myIP, err)
	}

	certFile := filepath.Join(c.DataDir, "consul.pem")
	err = ioutil.WriteFile(certFile, []byte(cert.Cert), 0644)
	if err != nil {
		return fmt.Errorf("error writing cert: %v", err)
	}

	keyFile := filepath.Join(c.DataDir, "consul-key.pem")
	err = ioutil.WriteFile(keyFile, []byte(cert.PrivateKey), 0644)
	if err != nil {
		return fmt.Errorf("error writing cert key: %v", err)
	}

	err = ioutil.WriteFile(caFile, []byte(cert.CA), 0644)
	if err != nil {
		return fmt.Errorf("error writing CA: %v", err)
	}

	c.TLSConfig = cert
	return nil
}

func runConsul(ctx context.Context, e runenv.Env, yc *yurtConfig) runner.Harness {
	myName, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname: %v", err)
	}

	// TODO add tls when used
	command := consul.NewConfig(yc.IsConsulServer(), yc.ConsulServerIPs, yc.TLSConfig)
	r, err := e.Run(ctx, command, yurt.Node{
		Name: myName,
	})
	if err != nil {
		log.Fatal(err)
	}

	return r
}

func runNomad(ctx context.Context, e runenv.Env, yc *yurtConfig) runner.Harness {
	myName, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname: %v", err)
	}

	expect := 0
	if yc.IsConsulServer() {
		expect = len(yc.ConsulServerIPs)
	}
	command := nomad.NewConfig(expect, fmt.Sprintf("127.0.0.1:%d", consul.DefPorts().HTTP), yc.TLSConfig)
	r, err := e.Run(ctx, command, yurt.Node{
		Name: myName,
	})
	if err != nil {
		log.Fatal(err)
	}

	return r
}
