package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile("[^\t\v\r\n\f]+")

type wordConverted struct {
	key   string
	value int
}

func Top10(s string) []string {

	lines := re.FindAllString(s, -1)
	words := []string{}
	for _, line := range lines {
		wordsLine := strings.Fields(line)
		for _, word := range wordsLine {
			words = append(words, word)
		}
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		count := wordMap[word] + 1
		wordMap[word] = count
	}

	wordSlice := make([]wordConverted, 0, len(wordMap))
	for key, value := range wordMap {
		wordSlice = append(wordSlice, wordConverted{key, value})
	}

	sort.Slice(wordSlice, func(i, j int) bool {
		if wordSlice[i].value == wordSlice[j].value {
			return strings.Compare(wordSlice[i].key, wordSlice[j].key) == -1
		} else {
			return wordSlice[i].value > wordSlice[j].value
		}
	})

	n := int(math.Min(10, float64(len(wordSlice))))
	wordSlice = wordSlice[:n]
	var keys []string
	for _, word := range wordSlice {
		keys = append(keys, word.key)
	}
	return keys
}
