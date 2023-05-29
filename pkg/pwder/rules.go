package pwder

import (
	"log"
	"math/rand"
	"unicode"
)

var ruleSet = map[rune]rune{
	'a': '4',
	'A': '4',
	'b': '8',
	'B': '8',
	'e': '3',
	'E': '3',
	'i': '7',
	'I': '7',
	'l': '1',
	'L': '1',
	'o': '0',
	'O': '0',
	's': '5',
	'S': '5',
	'z': '2',
	'Z': '2',
}

func CaseTransform(chr rune) rune {
	if unicode.IsUpper(chr) {
		return unicode.ToLower(chr)
	}

	return unicode.ToUpper(chr)
}

func ReplaceAtIndex(str string, chr rune, idx int) string {
	if len(str) <= idx {
		log.Fatalln("[ERROR]: Invalid index provided.")
	}

	result := []rune(str)
	result[idx] = chr
	return string(result)
}

func RuleTransform(str string, n int) string {
	if len(str) < n {
		log.Fatalln("[ERROR]: Word is shorter than number of transforms.")
	}

	for i := 0; i < n; i++ {
		idx := rand.Intn(len(str) - 2) + 2

		if n & 2 == 0 {
			chr := CaseTransform(rune(str[idx]))
			str = ReplaceAtIndex(str, chr, idx)
		} else {
			chr, inRules := ruleSet[rune(str[idx])]

			if !inRules {
				continue
			}

			str = ReplaceAtIndex(str, chr, idx)
		}
	}

	return str
}
