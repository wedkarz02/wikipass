package aeswrapper

import (
	"crypto/sha512"
	"encoding/hex"
	"log"
	"os"
	"wikipass/pkg/consts"
)

func HashBytes(plainText []byte) [64]byte {
	return sha512.Sum512(plainText)
}

func ByteToString(hexSum []byte) string {
	return hex.EncodeToString(hexSum)
}

func ReadIV(fileName string) []byte {
	iv, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalln("[ERROR]: IV file reading failed: ", err)
	}

	return iv
}

func GenKey(masterPassword string) []byte {
	hashedPassword := HashBytes([]byte(masterPassword))

	iv := ReadIV(consts.IVFile)
	key := make([]byte, 32)

	if len(iv) != len(key) {
		log.Fatalln("[ERROR]: len(iv) and len(key) do not match!")
	}

	for i := range iv {
		key[i] = hashedPassword[i] ^ iv[i]
	}

	return key
}
