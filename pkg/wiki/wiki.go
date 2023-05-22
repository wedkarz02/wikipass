package wiki

import (
	"fmt"
	"os"

	// https://www.mediawiki.org/wiki/API:Main_page
	"cgt.name/pkg/go-mwclient"
)

// TODO: big refactor of this module. (just testing for now)

func saveJSON(fileName string, data string) {
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString(data)
}

func ApiTest() {
	fmt.Println("hello from wikiapi")

	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wikiPasswordSearch")

	if err != nil {
		panic(err)
	}

	data, _, err := w.GetPageByName("pizza")

	if err != nil {
		panic(err)
	}

	saveJSON("./pkg/wiki/data.json", data)
}
