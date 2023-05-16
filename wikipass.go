package main

import (
	"wikipass/src"
)

func main() {
	src.EncryptAES([]byte("Hello, world!"), []byte("#secret-that-has-to-be-32-bytes!"))
}
