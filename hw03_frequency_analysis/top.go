package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type List struct {
	word  string
	count int
}

var reg = regexp.MustCompile(`[\s]+-|[\s,."!]+`)

func Top10(text string) []string {
	text = strings.ToLower(text)
	replaceText := reg.ReplaceAllString(text, string(' '))
	words := strings.Fields(replaceText)
	mapWords := make(map[string]int)

	for _, word := range words {
		mapWords[word]++
	}

	list := make([]List, 0, len(mapWords))
	for w, c := range mapWords {
		list = append(list, List{w, c})
	}

	sort.Slice(list, func(i int, j int) bool {
		if list[i].count == list[j].count {
			return list[i].word < list[j].word
		}
		return list[i].count > list[j].count
	})

	lengthList := 10
	if len(list) < 10 {
		lengthList = len(list)
	}

	listSlice := make([]string, 0, lengthList)
	for i, k := range list {
		listSlice = append(listSlice, k.word)
		if i == 9 || i > lengthList {
			break
		}
	}

	return listSlice
}
