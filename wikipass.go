package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
)

const encryptionFile = "./pkg/aeswrapper/passwd.aes"

func main() {
	plainText := "Hello, world!"
	key := "#secret-that-has-to-be-32-bytes!"

	fmt.Println(plainText, key)
	aeswrapper.EncryptAES(encryptionFile, []byte(plainText), []byte(key))

	message := aeswrapper.DecryptAES(encryptionFile, []byte(key))
	fmt.Println(string(message))
}
