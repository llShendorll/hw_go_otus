package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	arr := make([]string, 0)
	var flag bool
	for i, str := range text {
		if unicode.IsDigit(str) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(str) && unicode.IsDigit(rune(text[i-1])) && rune(text[i-2]) != '\\' {
			return "", ErrInvalidString
		}
		if unicode.IsSpace(str) {
			return "", ErrInvalidString
		}
		if string(str) == `\` {
			if !flag {
				flag = true
			} else {
				flag = false
			}
		}

		if atoi, err := strconv.Atoi(string(str)); err == nil {
			switch {
			case flag:
				if unicode.IsLetter(rune(text[i-1])) {
					arr = append(arr, strings.Repeat("\\"+string(text[i-1]), atoi))
				} else {
					arr = append(arr, string(str))
				}
				flag = false
			default:
				arr = arr[:len(arr)-1]
				arr = append(arr, strings.Repeat(string(text[i-1]), atoi))
			}
		} else if !flag {
			arr = append(arr, string(str))
		}
	}

	var sb strings.Builder
	for _, item := range arr {
		sb.WriteString(item)
	}

	return sb.String(), nil
}
