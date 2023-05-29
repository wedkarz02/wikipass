package main

import (
	"fmt"
	"math/rand"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
	"wikipass/pkg/pwder"
	"wikipass/pkg/wiki"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir, consts.IVFile, 32)

	title := wiki.GetRandArticle()
	content := wiki.GetArticleContent(title)
	wordList := wiki.ExtractWords(content)

	fmt.Println(len(wordList))

	for _, word := range wordList {
		word = pwder.RuleTransform(word, rand.Intn(len(word)))
		fmt.Println(word)
	}
}
