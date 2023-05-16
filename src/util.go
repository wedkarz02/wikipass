package src

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func EncryptAES(plainText []byte, key []byte) {
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
		log.Fatalln("[ERROR]: Something went wrong while seeding the nounce: ", err)
	}

	fmt.Println(string(gcm.Seal(nonce, nonce, plainText, nil)))
}
