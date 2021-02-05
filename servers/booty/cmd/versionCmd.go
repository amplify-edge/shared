package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Booty",
	Long:  `All software has versions. This is Booty's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
