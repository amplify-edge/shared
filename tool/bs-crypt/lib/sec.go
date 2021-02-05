package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	memCost       = 64 * 1024
	blockSize     = 8
	parallelCount = 1
	kLength       = 32
)

func Encrypt(key, data []byte) ([]byte, error) {
	key, salt, err := setKey(key, nil)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = append(ciphertext, salt...)
	return ciphertext, nil
}

func Decrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-kLength:], data[:len(data)-kLength]
	key, _, err := setKey(key, salt)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
func setKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, kLength)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	key, err := scrypt.Key(password, salt, memCost, blockSize, parallelCount, kLength)
	if err != nil {
		return nil, nil, err
	}
	return key, salt, nil
}

func EncryptAllFiles(srcdir, destdir, password string) error {
	_, err := os.Stat(destdir)
	if err != nil {
		err = os.MkdirAll(destdir, 0755)
		if err != nil {
			return err
		}
	}
	fi, err := os.Stat(srcdir)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return fmt.Errorf("requires a directory, not a file path")
	}
	transformFileFunc := func(fpath string) *string {
		fname := fmt.Sprintf("%s.%s", filepath.Base(fpath), "sec")
		dirName := filepath.Dir(fpath)
		joined := filepath.Join(dirName, fname)
		return &joined
	}
	return walkerFunc(srcdir, destdir, password, Encrypt, transformFileFunc)
}

func DecryptAllFiles(srcdir, destdir, password string) error {
	_, err := os.Stat(destdir)
	if err != nil {
		err = os.MkdirAll(destdir, 0755)
		if err != nil {
			return err
		}
	}
	fi, err := os.Stat(srcdir)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return fmt.Errorf("requires a directory, not a file path")
	}
	transformFileFunc := func(fpath string) *string {
		fname := filepath.Base(fpath)
		dirName := filepath.Dir(fpath)
		fsplit := strings.Split(fname, ".")
		if len(fsplit) > 2 {
			var newName []string
			lastIdx := len(fsplit) - 1
			for i := 0; i < lastIdx; i++ {
				newName = append(newName, fsplit[i])
			}
			newFileName := strings.Join(newName, ".")
			destName := filepath.Join(dirName, newFileName)
			return &destName
		}
		return nil
	}
	return walkerFunc(srcdir, destdir, password, Decrypt, transformFileFunc)
}

func walkerFunc(srcdir, destdir, password string, dataFunc func([]byte, []byte) ([]byte, error), transformFileNameFunc func(string) *string) error {
	return filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		fileRelPath, err := filepath.Rel(srcdir, path)
		if err != nil {
			return err
		}
		destFileName := filepath.Join(destdir, fileRelPath)
		if info.IsDir() {
			err = os.MkdirAll(destFileName, 0755)
			if err != nil {
				return fmt.Errorf("unable to create directory in %s => %v", destFileName, err)
			}
		} else {
			// Read the file
			fileByte, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("unable to read file at path: %s => %v", path, err)
			}
			cipherByte, err := dataFunc([]byte(password), fileByte)
			if err != nil {
				return err
			}
			transformed := transformFileNameFunc(destFileName)
			if transformed != nil {
				err = ioutil.WriteFile(*transformed, cipherByte, 0644)
				if err != nil {
					return fmt.Errorf("unable to write encrypted file to %s => %v", destFileName, err)
				}
			}
		}
		return nil
	})
}
