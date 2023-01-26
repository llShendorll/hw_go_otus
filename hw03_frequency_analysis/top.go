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

//var reg = regexp.MustCompile(`[\s]+-|[\s,."!]+`) // Раскомментировать

var reg = regexp.MustCompile(`[\s]+`) // taskWithAsteriskIsCompleted = false (true - Закомментировать)

func Top10(text string) []string {
	//text = strings.ToLower(text) // Раскомментировать taskWithAsteriskIsCompleted = true
	res := reg.ReplaceAllString(text, string(' '))
	words := strings.Fields(res)
	sort.Strings(words)
	cnt := 1
	list := make([]List, 0, len(words))
	for i, word := range words {
		if i+1 < len(words) {
			if word == words[i+1] {
				cnt++
			} else {
				list = append(list, List{word, cnt})
				cnt = 1
			}
		}
	}

	sort.SliceStable(list, func(i int, j int) bool {
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
