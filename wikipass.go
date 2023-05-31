package main

import (
	"fmt"
	"sync"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
	"wikipass/pkg/pwder"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir, consts.IVFile, 32)
	var wg sync.WaitGroup
	var passwords []string
	passwdChan := make(chan string)

	n := 5

	for i := 0; i < n; i++ {
		wg.Add(1)
		go pwder.GenPassword(passwdChan, &wg)
	}

	for i := 0; i < n; i++ {
		passwords = append(passwords, <-passwdChan)
	}

	wg.Wait()

	for _, passwd := range passwords {
		fmt.Println(passwd)
	}
}
