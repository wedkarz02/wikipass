package wiki

import (
	"fmt"
	"log"

	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
)

func GetRandArticle() {
	wiki, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wikiPasswordSearch")

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while connecting to the API: ", err)
	}

	parameters := params.Values{
		"action":  "query",
		"format":  "json",
		"list":    "random",
		"rnlimit": "1",
	}

	response, err := wiki.Get(parameters)

	if err != nil {
		log.Fatalln("[ERROR]: Invalid API server response: ", err)
	}

	fmt.Println(response)
}
