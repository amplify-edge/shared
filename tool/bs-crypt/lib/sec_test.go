package lib_test

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/amplify-cms/shared/tool/bs-crypt/lib"
	"github.com/stretchr/testify/require"
)

const (
	testPlainDir     = "./plain"
	testEncryptedDir = "./encrypted"
	testDecryptDir   = "./decrypted"
)

func init() {
	_ = os.MkdirAll(testEncryptedDir, 0755)
	_ = os.Mkdir(testDecryptDir, 0755)
}

func TestAll(t *testing.T) {
	t.Run("test encrypt & decrypt", testEncryptDecrypt)
	t.Run("test encrypt decrypt all files in dir", testEncDecAll)
}

func testEncryptDecrypt(t *testing.T) {
	var (
		password = []byte("mydyingbride")
		data     = []byte("some stuff in here")
	)
	ciphertext, err := lib.Encrypt(password, data)
	require.NoError(t, err)

	t.Logf("ciphertext: %s\n", base64.RawStdEncoding.EncodeToString(ciphertext))
	plaintext, err := lib.Decrypt(password, ciphertext)
	require.NoError(t, err)

	t.Logf("plaintext: %s\n", plaintext)

	require.Equal(t, plaintext, data)
}

func testEncDecAll(t *testing.T) {
	var password = "trazynInfinitum"
	err := lib.EncryptAllFiles(testPlainDir, testEncryptedDir, password)
	require.NoError(t, err)

	err = lib.DecryptAllFiles(testEncryptedDir, testDecryptDir, password)
	require.NoError(t, err)
}
