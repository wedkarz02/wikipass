package main

import (
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir, consts.IVFile, 32)

	masterPassword := "VeryStrongPasswordThatWillBeHashedInAMoment"
	key := aeswrapper.GenKey(masterPassword)

	plainText := []byte("Hello world!")

	aeswrapper.EncryptAES(consts.EncryptionFile, plainText, key)

	// message := aeswrapper.DecryptAES(consts.EncryptionFile, key)
	// fmt.Println(string(message))
}
