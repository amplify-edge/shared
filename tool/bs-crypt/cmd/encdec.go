package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.amplifyedge.org/shared-v2/tool/bs-crypt/lib"
	"os"
)

var (
	srcDir string
	dstDir string
)

const (
	defaultSrcDir       = "./plain"
	defaultDstDir       = "./encrypted"
	defaultEncryptUsage = "encrypt -s <SRC> -d <DEST>"
	defaultDecryptUsage = "decrypt -s <SRC> -d <DEST>"
)

func EncryptCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   defaultEncryptUsage,
		Short: defaultEncryptUsage,
	}
	rootCmd.PersistentFlags().StringVarP(&srcDir, "src", "s", defaultSrcDir, "source directory for encryption / decryption")
	rootCmd.PersistentFlags().StringVarP(&dstDir, "dest", "d", defaultDstDir, "destination directory for encryption / decryption")
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		passwd := os.Getenv("BS_CRYPT_PASSWORD")
		if passwd == "" {
			return fmt.Errorf("BS_CRYPT_PASSWORD is empty")
		}
		return lib.EncryptAllFiles(srcDir, dstDir, passwd)
	}
	return rootCmd
}

func DecryptCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   defaultDecryptUsage,
		Short: defaultDecryptUsage,
	}
	rootCmd.PersistentFlags().StringVarP(&srcDir, "src", "s", defaultSrcDir, "source directory for encryption / decryption")
	rootCmd.PersistentFlags().StringVarP(&dstDir, "dest", "d", defaultDstDir, "destination directory for encryption / decryption")
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		passwd := os.Getenv("BS_CRYPT_PASSWORD")
		if passwd == "" {
			return fmt.Errorf("BS_CRYPT_PASSWORD is empty")
		}
		return lib.DecryptAllFiles(srcDir, dstDir, passwd)
	}
	return rootCmd
}
