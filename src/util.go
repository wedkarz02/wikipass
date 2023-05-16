package src

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func EncryptAES(plainText []byte, key []byte) []byte {
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

	return gcm.Seal(nonce, nonce, plainText, nil)
}

func DecryptAES(cipherText []byte, key []byte) []byte {
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while initializing the cipher: ", err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while initializing GCM: ", err)
	}

	nonceSize := gcm.NonceSize()

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
