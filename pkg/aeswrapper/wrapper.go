package aeswrapper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func CheckIfExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func RmDir(path string) {
	err := os.RemoveAll(path)

	if err != nil {
		log.Fatalln("[ERROR]: Directory deletion failed: ", err)
	}
}

func InitIV(fileName string, size uint8) {
	iv := make([]byte, size)

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalln("[ERROR]: Generation of the IV failed: ", err)
	}

	err := os.WriteFile(fileName, iv, 0644)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while writing to the file: ", err)
	}
}

func InitSecretDir(dirPath string, ivPath string, ivSize uint8) {
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dirPath, os.ModePerm)

		if err != nil {
			log.Fatalln("[ERROR]: Something went wrong while creating a directory \"secret\": ")
		}

		InitIV(ivPath, ivSize)
	}
}

func InitAuth(fileName string, key []byte, size int) {
	randBytes := make([]byte, size)

	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		log.Fatalln("[ERROR]: Generation of the Auth file failed: ", err)
	}

	strRandBytes := []string{string(randBytes)}
	EncryptAES(fileName, strRandBytes, key)
}

func TestAuth(fileName string, key []byte) bool {
	_, err := DecryptAES(fileName, key)
	return err == nil
}

func StringsToByte(str []string) []byte {
	result := []byte{}

	for _, word := range str {
		row := []byte(word)
		row = append(row, byte('\n'))
		result = append(result, row...)
	}

	return result
}

func ByteToStrings(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func EncryptAES(fileName string, data []string, key []byte) {
	byteData := StringsToByte(data)
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while initializing the cipher: ", err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while initializing GCM: ", err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln("[ERROR]: Something went wrong while seeding the nonce: ", err)
	}

	err = os.WriteFile(fileName, gcm.Seal(nonce, nonce, byteData, nil), 0644)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while writing to the file: ", err)
	}
}

func DecryptAES(fileName string, key []byte) ([]string, error) {
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while initializing the cipher: ", err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while initializing GCM: ", err)
	}

	nonceSize := gcm.NonceSize()

	cipherText, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while reading the file: ", err)
	}

	if len(cipherText) < nonceSize {
		log.Fatalln("[ERROR]: Cipher length is shorter than the nonce!")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		// The tested key is incorrect -> return an empty slice and an error
		return []string{}, err
	}

	return ByteToStrings(plainText), nil
}
