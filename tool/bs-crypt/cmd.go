package main

import (
	"go.amplifyedge.org/shared-v2/tool/bs-crypt/cmd"
	"log"

	"github.com/spf13/cobra"
)

const (
	defaultUsage = "bs-crypt <encrypt | decrypt> --src <SRCPATH> --dest <DESTPATH>"
	commandName  = "bs-crypt"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   defaultUsage,
		Short: defaultUsage,
		Long:  defaultUsage,
	}
	rootCmd.AddCommand(cmd.EncryptCmd(), cmd.DecryptCmd())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
