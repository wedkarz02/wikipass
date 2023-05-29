package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
	"wikipass/pkg/pwder"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir, consts.IVFile, 32)

	word := "aAbggSeRllkK"
	fmt.Println(word)
	
	word = pwder.RuleTransform(word, 7)
	fmt.Println(word)

	
	// for i, chr := range word {
	// 	repl := pwder.CaseTransform(chr)
	// 	word = pwder.ReplaceAtIndex(word, repl, i)
	// }

	// fmt.Println(word)
}
