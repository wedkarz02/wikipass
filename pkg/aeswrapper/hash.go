package aeswrapper

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashBytes(plainText []byte) [64]byte {
	return sha512.Sum512(plainText)
}

func ByteToString(hexSum []byte) string {
	return hex.EncodeToString(hexSum)
}

// TODO: XOR the master password with some IV for a better key generation
func GenKey(masterPassword string) []byte {
	return []byte(masterPassword)[:32]
}
