package main

import (
	"fmt"
	"log"
	"os"

	"go.amplifyedge.org/shared-v2/tool/bs-crypt/lib"
	"github.com/spf13/cobra"
)

const (
	defaultSrcDir = "./plain"
	defaultDstDir = "./encrypted"
	defaultUsage  = "bs-crypt <encrypt | decrypt> --src <SRCPATH> --dest <DESTPATH>"
	commandName   = "bs-crypt"

	defaultEncryptUsage = "encrypt -s <SRC> -d <DEST>"
	defaultDecryptUsage = "decrypt -s <SRC> -d <DEST>"
)

var (
	srcDir string
	dstDir string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   defaultUsage,
		Short: defaultUsage,
		Long:  defaultUsage,
	}
	rootCmd.AddCommand(encryptCmd(), decryptCmd())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func encryptCmd() *cobra.Command {
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

func decryptCmd() *cobra.Command {
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
