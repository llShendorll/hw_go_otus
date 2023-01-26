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
	runeText := []rune(text)
	var sb strings.Builder

	for i, str := range runeText {
		if unicode.IsDigit(str) && i == 0 {
			return "", ErrInvalidString
		} else if unicode.IsDigit(str) && unicode.IsDigit(runeText[i-1]) && runeText[i-2] != '\\' {
			return "", ErrInvalidString
		}

		atoi, _ := strconv.Atoi(string(str))
		switch {
		case string(str) == `\` && !flag:
			flag = true
		case unicode.IsLetter(str):
			if i+1 < len(runeText) && !unicode.IsDigit(runeText[i+1]) {
				sb.WriteString(string(str))
			} else if i+1 == len(runeText) && !unicode.IsDigit(runeText[i]) {
				sb.WriteString(string(str))
			}
		case !flag && unicode.IsDigit(runeText[i]):
			if unicode.IsDigit(runeText[i-1]) || string(runeText[i-1]) == "\\" {
				sb.WriteString(strings.Repeat(string(runeText[i-1]), atoi-1))
			} else {
				sb.WriteString(strings.Repeat(string(runeText[i-1]), atoi))
			}
		case flag && unicode.IsLetter(runeText[i-1]):
			sb.WriteString(strings.Repeat("\\"+string(runeText[i-1]), atoi))
		default:
			sb.WriteString(string(str))
			flag = false
		}
	}

	return sb.String(), nil
}
