package pwder

import "unicode"

func CaseTransform(chr rune) rune {
	if unicode.IsUpper(chr) {
		return unicode.ToLower(chr)
	}

	return unicode.ToUpper(chr)
}

func ReplaceAtIndex(str string, chr rune, idx int) string {
	result := []rune(str)
	result[idx] = chr
	return string(result)
}
