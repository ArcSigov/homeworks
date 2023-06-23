package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[^-\pL]`)

func Top10(inputStr string) []string {
	inputWords := strings.Fields(inputStr)
	freqWords := make(map[string]int)
	for _, word := range inputWords {
		if word = strings.ToLower(re.ReplaceAllString(word, "")); len(word) > 0 && word != "-" {
			freqWords[word]++
		}
	}

	words := make([]string, 0, len(freqWords))
	for k := range freqWords {
		words = append(words, k)
	}

	sort.Slice(words, func(i, j int) bool {
		a := freqWords[words[i]]
		b := freqWords[words[j]]
		if a == b {
			return words[i] < words[j]
		}
		return a > b
	})

	maxresults := len(words)
	if maxresults > 10 {
		maxresults = 10
	}
	return words[:maxresults]
}
