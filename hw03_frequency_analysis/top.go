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
	res := reg.ReplaceAllString(text, string(' '))
	words := strings.Fields(res)
	mapWords := make(map[string]int)

	for _, word := range words {
		if mapWords[word] != 0 {
			mapWords[word]++
		} else {
			mapWords[word] = 1
		}
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

	listSlice := make([]string, 0, 10)
	for i, k := range list {
		listSlice = append(listSlice, k.word)
		if i == 9 || i > len(list) {
			break
		}
	}

	return listSlice
}
