package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
)

func main() {
	aeswrapper.MakeSecretDir(consts.SecretDir)

	plainText := "a"
	key := "#secret-that-has-to-be-32-byt!"

	fmt.Println(plainText)
	fmt.Println(key)

	aeswrapper.EncryptAES(consts.EncryptionFile, []byte(plainText), []byte(key))

	message := aeswrapper.DecryptAES(consts.EncryptionFile, []byte(key))
	fmt.Println(string(message))

	hashMessage := aeswrapper.HashBytes([]byte(plainText))
	stringHash := aeswrapper.ByteToString(hashMessage[:])

	fmt.Println(stringHash)
}
