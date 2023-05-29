package wiki

import (
	"log"
	"strings"

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

	queryRand, err := query.GetObjectArray("random")

	if err != nil {
		log.Fatalln("[ERROR]: JSON query parsing failed: ", err)
	}

	title, err := queryRand[0].GetString("title")

	if err != nil {
		log.Fatalln("[ERROR]: JSON id parsing failed: ", err)
	}

	title = strings.ReplaceAll(title, " ", "_")
	titleSplit := strings.Split(title, ":")
	title = titleSplit[len(titleSplit)-1]

	return title
}

func GetArticleContent(title string) string {
	wiki, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wikiGetContent")

	if err != nil {
		log.Fatalln("[ERROR]: Something went wrong while connecting to the API: ", err)
	}

	parameters := params.Values{
		"action":        "query",
		"prop":          "revisions",
		"rvslots":       "*",
		"rvprop":        "content",
		"formatversion": "2",
		"format":        "json",
		"titles":        title,
	}

	pageData, err := wiki.Get(parameters)

	if err != nil {
		log.Fatalln("[ERROR]: Invalid API server response: ", err)
	}

	query, err := pageData.GetObject("query")

	if err != nil {
		log.Fatalln("[ERROR]: Query parsing failed: ", err)
	}

	pages, err := query.GetObjectArray("pages")

	if err != nil {
		log.Fatalln("[ERROR]: Pages parsing failed: ", err)
	}

	if _, err := pages[0].GetBoolean("missing"); err == nil {
		panicTitle := GetRandArticle()
		return GetArticleContent(panicTitle)
	}

	revisions, err := pages[0].GetObjectArray("revisions")

	if err != nil {
		log.Fatalln("[ERROR]: Revisions parsing failed: ", err)
	}

	slots, err := revisions[0].GetObject("slots")

	if err != nil {
		log.Fatalln("[ERROR]: Slots parsing failed: ", err)
	}

	mainObj, err := slots.GetObject("main")

	if err != nil {
		log.Fatalln("[ERROR]: MainObj parsing failed: ", err)
	}

	content, err := mainObj.GetString("content")

	if err != nil {
		log.Fatalln("[ERROR]: Content parsing failed: ", err)
	}

	if strings.Contains(content, "#REDIRECT") || strings.Contains(content, "#redirect") {
		panicTitle := GetRandArticle()
		return GetArticleContent(panicTitle)
	}

	return content
}

func ExtractWords(content string) []string {
	fullExtracted := strings.FieldsFunc(content, func(r rune) bool {
		return (r == ' '  ||
			    r == '\n' ||
			    r == ','  ||
			    r == '.'  ||
			    r == '\'' ||
			    r == '"'  ||
			    r == ':'  ||
			    r == '/'  ||
			    r == '\\' ||
			    r == '{'  ||
			    r == '}'  ||
			    r == '('  ||
			    r == ')'  ||
			    r == '['  ||
			    r == ']'  ||
			    r == '|'  ||
			    r == '<'  ||
			    r == '>'  ||
			    r == '='  )
	})

	var wordList []string

	for _, str := range fullExtracted {
		if len(str) > 4 && len(str) < 15 {
			wordList = append(wordList, str)
		}
	}

	return wordList
}
