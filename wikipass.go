package main

import (
	"fmt"
	"wikipass/src"
)

func main() {
	cipher := src.EncryptAES([]byte("Hello, world!"), []byte("#secret-that-has-to-be-32-bytes!"))
	fmt.Println(string(cipher))

	plain := src.DecryptAES(cipher, []byte("#secret-that-has-to-be-32-bytes!"))
	fmt.Println(string(plain))
}
