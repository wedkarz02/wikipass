package main

import (
	"log"
	"os"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
)

func logFile(fileName string, logData string) {
	file, err := os.OpenFile(consts.LogDir+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while creating log directory: ", err)
	}

	defer file.Close()

	if _, err = file.WriteString(logData + "\n"); err != nil {
		log.Fatalln("[ERROR]: Something went wrong while writing to logs: ", err)
	}
}

func main() {
	aeswrapper.MakeSecretDir(consts.SecretDir)

	plainText := "Hello, world!"
	key := "#secret-that-has-to-be-32-bytes!"

	logFile(consts.PrintLog, plainText)
	logFile(consts.PrintLog, key)

	aeswrapper.EncryptAES(consts.EncryptionFile, []byte(plainText), []byte(key))

	message := aeswrapper.DecryptAES(consts.EncryptionFile, []byte(key))
	logFile(consts.PrintLog, string(message))

	hashMessage := aeswrapper.HashBytes([]byte(plainText))
	stringHash := aeswrapper.ByteToString(hashMessage[:])

	logFile(consts.PrintLog, stringHash)
}
