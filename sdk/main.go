package main

import (
	"os"
	/*
		"github.com/amplify-cms/sdk/cmd/yurt-run"
		"github.com/amplify-cms/sdk/cmd/yurt-cluster"
	*/
	yurt "github.com/amplify-cms/sdk/cmd/yurt-run"
)

func main() {
	resp := yurt.Execute(os.Args[1:])

	if resp.Err != nil {
		if resp.IsUserError() {
			resp.Cmd.Println("")
			resp.Cmd.Println(resp.Cmd.UsageString())
		}
		os.Exit(-1)
	}

}
