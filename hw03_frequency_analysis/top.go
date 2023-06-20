package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordinfo struct {
	word  string
	count int
}

var re = regexp.MustCompile(`[^-\pL]`)

func add(character string, sliceStr *[]wordinfo) {
	for pos := range *sliceStr {
		if (*sliceStr)[pos].word == character {
			(*sliceStr)[pos].count++
			return
		}
	}
	*sliceStr = append(*sliceStr, wordinfo{character, 1})
}

func Top10(inputStr string) []string {
	words := strings.Fields(inputStr)
	sliceStr := make([]wordinfo, 0, len(words))
	outputStr := make([]string, 0, 10)
	for _, word := range words {
		word = strings.ToLower(re.ReplaceAllString(word, ""))
		if len(word) > 0 && word != "-" {
			add(word, &sliceStr)
		}
	}
	sort.Slice(sliceStr, func(i, j int) bool {
		if sliceStr[i].count == sliceStr[j].count {
			return sliceStr[i].word < sliceStr[j].word
		}
		return sliceStr[i].count > sliceStr[j].count
	})
	for pos, wordInfo := range sliceStr {
		if pos == 10 {
			break
		}
		outputStr = append(outputStr, wordInfo.word)
	}
	return outputStr
}
