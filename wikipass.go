package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
	"wikipass/pkg/wiki"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir, consts.IVFile, 32)

	title := wiki.GetRandArticle()
	content := wiki.GetArticleContent(title)

	arr := wiki.ExtractWords(content)

	for _, el := range arr {
		fmt.Println(el)
	}
}
