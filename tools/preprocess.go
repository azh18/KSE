package tools

import (
	"sort"
	"strings"
	"unicode"
)

const ValidLength = 4

// pre-process a text into formulated words
// including removing all words that is not letters, removing 's' at the end, and transform to lowercase
func TextPreProcess(text string) []string{
	ret := make([]string, 0)
	checkLetterFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	keywords := strings.FieldsFunc(text, checkLetterFunc) // trim all punctuations, spaces...
	for _, s := range keywords {
		s = strings.TrimRight(s, "s") // trim 's' at the tail of words
		s = strings.ToLower(s)
		if len(s) < ValidLength{ // omit words with length smaller than 4
			continue
		}
		ret = append(ret, s)
	}
	return ret
}

func GetSortedKeysOfMap(m map[string]float64) []string{
	s := make([]string, 0)
	for k := range m{
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}