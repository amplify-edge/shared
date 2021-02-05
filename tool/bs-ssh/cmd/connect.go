package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connects to a server using ssh.",
	Long:  `The ssh key is used as identity to connect to a ssh server, so you can manage servers and upload files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connect called")
	},
}
