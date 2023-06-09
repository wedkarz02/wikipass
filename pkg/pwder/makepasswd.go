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

	return wordList
}

func SliceContains(list []string, el string) bool {
	for _, str := range list {
		if strings.EqualFold(str, el) {
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

		transformedWord := RuleTransform(wordList[randIdx])
		words = append(words, transformedWord)
	}

	ch <- strings.Join(words, "-")
}

func GetPasswords(n int) []string {
	var wg sync.WaitGroup
	var passwords []string

	ch := make(chan string)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go GenPassword(ch, &wg)
	}

	for i := 0; i < n; i++ {
		passwords = append(passwords, <-ch)
	}

	wg.Wait()
	return passwords
}
