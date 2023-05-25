package wiki

import (
	"fmt"
	"log"

	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
)

func GetRandTitle() {
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

	query, err := response.GetObject("query")

	if err != nil {
		log.Fatalln("[ERROR]: JSON response parsing failed: ", err)
	}

	jsonRand, err := query.GetObjectArray("random")

	if err != nil {
		log.Fatalln("[ERROR]: JSON query parsing failed: ", err)
	}

	for _, el := range jsonRand {
		fmt.Println(el)
	}
}
