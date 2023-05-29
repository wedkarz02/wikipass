package main

import (
	"fmt"
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

	for _, word := range wordList {
		word = pwder.RuleTransform(word)
		fmt.Println(word)
	}
}
