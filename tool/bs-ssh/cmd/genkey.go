package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genkeyCmd)
}

var genkeyCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a ssh key.",
	Long:  `Ssh Keys are used to prove your identity and connect to systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gen called")
	},
}
