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
