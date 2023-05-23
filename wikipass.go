package main

import (
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/wiki"
)

const secretDir = "./secret"
const encryptionFile = "./secret/passwd.aes"
const hashFile = "./secret/mpasswd.hash"

func main() {
	aeswrapper.MakeSecretDir(secretDir)

	// plainText := "Hello, world!"
	// key := "#secret-that-has-to-be-32-bytes!"

	// fmt.Println(plainText, key)
	// aeswrapper.EncryptAES(encryptionFile, []byte(plainText), []byte(key))

	// message := aeswrapper.DecryptAES(encryptionFile, []byte(key))
	// fmt.Println(string(message))

	// hashMessage := aeswrapper.HashBytes([]byte(plainText))
	// stringHash := aeswrapper.ByteToString(hashMessage[:])

	// fmt.Println(stringHash)

	wiki.ApiTest()
}
