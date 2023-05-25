package wiki

import (
	"fmt"
	"log"

	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
)

func GetRandArticle() string {
	wiki, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wikiRandArticle")

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

	query, err := response.GetObject("query")

	if err != nil {
		log.Fatalln("[ERROR]: JSON response parsing failed: ", err)
	}

	jsonRand, err := query.GetObjectArray("random")

	if err != nil {
		log.Fatalln("[ERROR]: JSON query parsing failed: ", err)
	}

	// pageId, err := jsonRand[0].GetNumber("id")
	title, err := jsonRand[0].GetString("title")

	if err != nil {
		log.Fatalln("[ERROR]: JSON id parsing failed: ", err)
	}

	fmt.Println(title)

	return title
}

func GetArticleContent(title string) {
	wiki, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wikiGetContent")

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while connecting to the API: ", err)
	}

	parameters := params.Values{
		"action":        "query",
		"prop":          "revisions",
		"titles":        title,
		"rvslots":       "",
		"rvprop":        "content",
		"formatversion": "2",
	}

	content, err := wiki.Get(parameters)

	if err != nil {
		log.Fatalln("[ERROR]: Invalid API server response: ", err)
	}

	fmt.Println(content)
}
