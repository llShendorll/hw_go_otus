package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var shielding bool
	var repeatLetter rune
	runeText := []rune(text)
	var sb strings.Builder
	for i, str := range runeText {
		switch {
		case string(str) == `\` && !shielding:
			if i == len(runeText)-1 {
				return "", ErrInvalidString
			}
			shielding = true
		case unicode.IsDigit(rune(text[0])) || (unicode.IsLetter(str) && shielding):
			return "", ErrInvalidString
		case shielding:
			sb.WriteString(string(str))
			repeatLetter = str
			shielding = false
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
	}

	return sb.String(), nil
}
