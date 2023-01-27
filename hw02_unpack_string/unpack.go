package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var flag bool
	var repeatLetter rune
	runeText := []rune(text)
	var sb strings.Builder
	for i, str := range runeText {
		switch {
		case string(str) == `\` && !flag:
			flag = true
		case flag:
			sb.WriteString(string(str))
			repeatLetter = str
			flag = false
		case unicode.IsDigit(str) && unicode.IsDigit(runeText[i-1]) && runeText[i-2] != '\\':
			return "", ErrInvalidString
		case unicode.IsDigit(str):
			if atoi, err := strconv.Atoi(string(str)); err == nil {
				if atoi > 0 {
					sb.WriteString(strings.Repeat(string(repeatLetter), atoi-1))
				}
			}
		default:
			if (i+1 < len(runeText) && runeText[i+1] != '0') || len(runeText) == i+1 {
				sb.WriteString(string(str))
				repeatLetter = str
			}
		}

		if unicode.IsDigit(rune(text[0])) || runeText[len(runeText)-1] == '\\' {
			return "", ErrInvalidString
		} else if unicode.IsLetter(str) && flag {
			return "", ErrInvalidString
		}
	}

	return sb.String(), nil
}
