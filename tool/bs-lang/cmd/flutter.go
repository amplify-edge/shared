package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"go.amplifyedge.org/shared-v2/tool/bs-lang/services"
)

var (
	dir        *string
	template   *string
	prefixName *string
	languages  *string
	full       *bool
	cacheFile  *string
)

// flutterCmd represents the flutter command
var flutterCmd = &cobra.Command{
	Use:   "flutter",
	Short: "Generate json and arb files.",
	Long:  `Allows to generate json and arb files already translated any languages wanted.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if *template == "" {
			return services.GenerateMultiLanguagesArbFilesFromJSONFiles(*dir, *prefixName, "json", "arb", *full)
		}
		return services.GenerateMultiLanguageFilesFromTemplate(*template, *dir, *prefixName, "json", getLanguages(*languages, ","), *full, *cacheFile)
	},
}

func init() {
	RootCmd.AddCommand(flutterCmd)
	dir = flutterCmd.Flags().StringP("dir", "d", ".", "Directory where to out and look for files.")
	template = flutterCmd.Flags().StringP("template", "t", "", "Template file path to generate multi languages files.")
	prefixName = flutterCmd.Flags().StringP("prefix", "p", "", "The prefix to add for each file generated.")
	languages = flutterCmd.Flags().StringP("languages", "l", "en,fr,es,de,it,ur,tr", "Languages list separated by coma.")
	full = flutterCmd.Flags().BoolP("full", "f", false, "Get full detailed out file example to generate json file without arb tags.")
	cacheFile = flutterCmd.Flags().StringP("cache", "c", filepath.Join(os.TempDir(), "transcache.json"), "Cache file location")
}

func getLanguages(languages, sep string) []string {
	return strings.Split(languages, sep)
}
