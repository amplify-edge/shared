package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	mainAssets "github.com/amplify-cms/shared/tool/sdk/bs-embed-ex/maintemplate/assets"
	mod01Assets "github.com/amplify-cms/shared/tool/sdk/bs-embed-ex/mod01/assets"
	"github.com/gorilla/mux"
)

// PageData structure
type PageData struct {
	Title       string
	Heading     string
	Description string
}

const (
	errMsg       = "error getting %s from %s static assets: %v\n"
	errParseMsg  = "error parsing %s to %s: %v\n"
	errListAsset = "error list assets for specified directory: %s, doesn't exists"
	httpPort     = 8080
)

// listAsset will list all embedded assets in the directory recursively
// naively however since it's a poc.
// 4. Config Time
func listAsset(dir string) ([]string, error) {
	// this is ugly but it's just proof of concept
	if strings.Contains(dir, "maintemplate") {
		return mainAssets.AssetNames(), nil
		// return RecursiveListFiles("maintemplate/static/")
	} else if strings.Contains(dir, "mod01") {
		return mod01Assets.AssetNames(), nil
	} else {
		return nil, errors.New(fmt.Sprintf(errListAsset, dir))
	}
}

// loadAsset will unpack embedded asset and unpack it to given directory
// parameter 1: assetDir is the name of the asset directory
// parameter 2: target is the target directory to unpack
// you can easily create a new handle func to test this.
// 5. Run Time
func loadAsset(assetDir, target string) error {
	switch assetDir {
	case "maintemplate":
		return mainAssets.RestoreAssets(assetDir, target)
	case "mod01":
		return mod01Assets.RestoreAssets(assetDir, target)
	default:
		return errors.New(fmt.Sprintf("%s not found", assetDir))
	}
}

func main() {

	// hello := pkg.Hello()
	//
	// fmt.Println(hello)

	// Reflect on bindata
	mainTemplateIndex, err := mainAssets.Asset("maintemplate/static/index.html")
	if err != nil {
		log.Fatalf(errMsg, "index.html", "maintemplate", err)
	}

	mod01Index, err := mod01Assets.Asset("mod01/static/index.html")
	if err != nil {
		log.Fatalf(errMsg, "index.html", "mod01", err)
	}

	// Template
	mainTmpl := template.Must(template.New("maintemplate").Parse(string(mainTemplateIndex)))
	mod01Tmpl := template.Must(template.New("mod01").Parse(string(mod01Index)))

	mainData := PageData{
		Title:       "Maintemplate Bindata",
		Heading:     "Hello from Maintemplate index.html",
		Description: "Hello Maintemplate Bindata",
	}

	mod01Data := PageData{
		Title:       "Mod01 Bindata",
		Heading:     "Hello from Mod01 index.html",
		Description: "Hello Mod01 Bindata",
	}

	router := mux.NewRouter()

	router.HandleFunc("/list/{param}", func(w http.ResponseWriter, r *http.Request) {
		par := mux.Vars(r)
		assetsList, err := listAsset(par["param"])
		if err != nil {
			log.Print(err)
		}
		respString := strings.Join(assetsList, "\n")
		w.Write([]byte(respString))
	})

	// Handle function
	router.HandleFunc("/maintemplate", func(w http.ResponseWriter, r *http.Request) {
		// data := PageData{
		// 	Title:       "The easiest way to embed static files into a binary file",
		// 	Heading:     "This is easiest way",
		// 	Description: "My life credo is 'If you can not use N, do not use N'.",
		// }

		// Execute template with data
		if err := mainTmpl.Execute(w, mainData); err != nil {
			log.Fatalf(errParseMsg, mainData, "maintemplate/static/index.html", err)
		}
	})

	router.HandleFunc("/mod01", func(w http.ResponseWriter, r *http.Request) {
		if err := mod01Tmpl.Execute(w, mod01Data); err != nil {
			log.Fatalf(errParseMsg, mod01Data, "mod01/static/index.html", err)
		}
	})

	log.Printf("Serving http server on port %d", httpPort)
	// Start server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router); err != nil {
		log.Fatal(err)
	}
}
