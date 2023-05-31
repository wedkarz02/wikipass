package pwder

import (
	"math/rand"
	"strings"
	"sync"
	"wikipass/pkg/wiki"
)

func InitWordList() []string {
	title := wiki.GetRandArticle()
	content := wiki.GetArticleContent(title)
	wordList := wiki.ExtractWords(content)

	for i, word := range wordList {
		wordList[i] = RuleTransform(word)
	}

	return wordList
}

func SliceContains(list []string, el string) bool {
	for _, str := range list {
		if el == str {
			return true
		}
	}

	return false
}

func GenPassword(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	var words []string
	wordList := InitWordList()

	for i := 0; i < 4; i++ {
		randIdx := rand.Intn(len(wordList))

		for SliceContains(words, wordList[randIdx]) {
			randIdx = rand.Intn(len(wordList))
		}

		words = append(words, wordList[randIdx])
	}

	ch <- strings.Join(words, "-")
}
