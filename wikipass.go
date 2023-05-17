package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
)

func main() {
	cipher := aeswrapper.EncryptAES([]byte("Hello, world!"), []byte("#secret-that-has-to-be-32-bytes!"))
	fmt.Println(string(cipher))

	plain := aeswrapper.DecryptAES(cipher, []byte("#secret-that-has-to-be-32-bytes!"))
	fmt.Println(string(plain))
}
