package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)
const backslash rune = '\\'

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var shielding bool
	var next bool
	var letter rune
	var textResult strings.Builder
	runeText := []rune(text)

	for i, r := range runeText {
		switch {
		case r == backslash && !shielding:
			if i == len(text)-1 {
				return "", ErrInvalidString
			}
			textResult.WriteRune(letter)
			shielding = true
		case unicode.IsDigit(runeText[0]) || (unicode.IsLetter(r) && shielding):
			return "", ErrInvalidString
		case unicode.IsDigit(r) && unicode.IsDigit(runeText[i-1]) && runeText[i-2] != backslash:
			return "", ErrInvalidString
		case shielding:
			if i+1 == len(runeText) {
				textResult.WriteRune(r)
			}
			shielding, letter, next = false, r, false
		case unicode.IsDigit(r):
			if r != '0' && !next {
				textResult.WriteString(strings.Repeat(string(letter), int(r-'0')))
			}
			next = true
		default:
			if !next && letter != 0 {
				textResult.WriteRune(letter)
			}
			if i+1 == len(runeText) {
				textResult.WriteRune(r)
			}
			letter, next = r, false
		}
	}

	return textResult.String(), nil
}
