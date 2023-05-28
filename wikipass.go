package main

import (
	"wikipass/pkg/aeswrapper"
	"wikipass/pkg/consts"
	"wikipass/pkg/wiki"
)

func main() {
	aeswrapper.InitSecretDir(consts.SecretDir, consts.IVFile, 32)

	title := wiki.GetRandArticle()
	wiki.GetArticleContent(title)
}
