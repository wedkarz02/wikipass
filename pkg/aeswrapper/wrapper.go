package aeswrapper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"os"
)

func MakeSecretDir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)

		if err != nil {
			log.Fatalln("[ERROR]: Something went wrong while creating a directory \"secret\": ")
		}
	}
}

func EncryptAES(fileName string, plainText []byte, key []byte) {
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

	// writeToFile(gcm.Seal(nonce, nonce, plainText, nil))
	err = os.WriteFile(fileName, gcm.Seal(nonce, nonce, plainText, nil), 0644)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while writing to the file: ", err)
	}
}

func DecryptAES(fileName string, key []byte) []byte {
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
		log.Fatalln("[ERROR]: Something went wrong while decrypting: ", err)
	}

	return plainText
}
