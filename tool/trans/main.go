package main

import (
	"flag"
	"fmt"
	"go.amplifyedge.org/shared-v2/tool/trans/lib"
	"go.amplifyedge.org/shared-v2/tool/trans/lib/gsheets"
	"log"
)

const (
	defaultGoogleCredPath = "./client_secret.json"
)

var (
	googCredPath string
)

func main() {

	flag.StringVar(&googCredPath, "c", defaultGoogleCredPath, "path to google service account client_secret.json")
	flag.Parse()

	cfg, err := lib.NewConfigFromFile("./config.json")
	if err != nil {
		log.Fatalf("unable to read config from path")
	}

	c, err := gsheets.NewClient(googCredPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	locName := c.WorksheetName()
	fmt.Println("Querying gsheet: ", locName)

	cells, err := c.Localizations()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cells: %v\n", cells)

	res, err := c.LastIdx()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response: %v", res)
}
