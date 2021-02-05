package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagLetters int
	langFlagLetters  string
)

//flagOS = flag.String("os", runtime.GOOS, "override default OS")
//flagArch    = flag.String("arch", runtime.GOARCH, "override default arch")

var getbinariesCmd = &cobra.Command{
	Use:   "getbinaries",
	Short: "Gtes the binaries.",
	Long:  `All software has dependencies. This is Booty's`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting binaries ...")

		/*

			// OS and Arch can be passed in, so that you can test this for any OS.
			*flagWorkDir = os.TempDir()


				binmgr, err := binaries.NewDownloadManager(*flagWorkDir)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				path, err := binmgr.GetOSArch(args[0], *flagOS, *flagArch, *flagVersion)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				fmt.Println(path)
		*/
	},
}

func init() {
	rootCmd.AddCommand(getbinariesCmd)

	getbinariesCmd.Flags().IntVarP(
		&Count, "count", "c", 0,
		"A count of random letters",
	)
	getbinariesCmd.MarkFlagRequired("count")

	getbinariesCmd.Flags().StringVarP(
		&Lang, "lang", "l", "en",
		"A language. Optional",
	)
}
