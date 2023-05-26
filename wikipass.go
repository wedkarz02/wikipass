package main

import (
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
	"wikipass/pkg/wiki"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir)

	title := wiki.GetRandArticle()
	wiki.GetArticleContent(title)
}
