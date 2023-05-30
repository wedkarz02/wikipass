package pwder

import (
	"math/rand"
	"unicode"
)

var extraSet string = "!@#$%^&*?"

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
	return str[:idx] + string(chr) + str[idx+1:]
}

func RuleTransform(str string) string {
	for i, el := range str {
		transformChance := rand.Float32()

		if transformChance < 0.3 {
			continue
		}

		transformType := rand.Intn(2)

		if transformType == 0 {
			chr := CaseTransform(el)
			str = ReplaceAtIndex(str, chr, i)
		} else {
			chr, inRules := ruleSet[el]

			if !inRules {
				continue
			}

			str = ReplaceAtIndex(str, chr, i)
		}
	}

	extraChance := rand.Float32()

	if extraChance >= 0.5 {
		randChr := extraSet[rand.Intn(len(extraSet))]
		str = str + string(randChr)
	}

	return str
}
