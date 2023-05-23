package wiki

import (

	// https://www.mediawiki.org/wiki/API:Main_page

	"fmt"

	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
)

// TODO: big refactor of this module. (just testing for now)

// func saveJSON(fileName string, data string) {
// 	file, err := os.Create(fileName)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()

// 	file.WriteString(data)
// }

func ApiTest() {
	// fmt.Println("hello from wikiapi")

	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wikiPasswordSearch")

	if err != nil {
		panic(err)
	}

	parameters := params.Values{
		"action":  "query",
		"format":  "json",
		"list":    "random",
		"rnlimit": "5",
	}

	response, err := w.Get(parameters)

	if err != nil {
		panic(err)
	}

	// data := map[string]string{}

	fmt.Println(response)
	// var data map[string]interface{}
	// json.Unmarshal([]byte(response), &data)

	// for _, el := range response.Map() {
	// 	fmt.Println(*el)
	// }

	// saveJSON("./pkg/wiki/data.json", data)
}
