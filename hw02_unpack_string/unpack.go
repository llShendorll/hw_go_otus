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
		if unicode.IsDigit(str) && i == 0 {
			return "", ErrInvalidString
		} else if unicode.IsDigit(str) && unicode.IsDigit(runeText[i-1]) && runeText[i-2] != '\\' {
			return "", ErrInvalidString
		}

		switch {
		case string(str) == `\` && !flag:
			flag = true
		case flag:
			sb.WriteString(string(str))
			repeatLetter = str
			flag = false
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
	}

	return sb.String(), nil
}
