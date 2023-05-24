package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir)

	// Ask the user for the master password here.
	masterPassword := "VeryStrongPasswordThatWillBeHashedInAMoment"
	key := aeswrapper.GenKey(masterPassword)

	plainText := "Hello world!"

	fmt.Println(plainText)
	fmt.Println(aeswrapper.ByteToString(key))

	aeswrapper.EncryptAES(consts.EncryptionFile, []byte(plainText), []byte(key))

	message := aeswrapper.DecryptAES(consts.EncryptionFile, []byte(key))
	fmt.Println(string(message))
}
