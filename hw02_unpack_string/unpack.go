package hw02unpackstring

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	builder := strings.Builder{}
	re := regexp.MustCompile("[^0-9][0-9]+|[^0-9]|^[0-9]")
	d := re.FindAllString(s, -1)
	for _, latter := range d {
		newWord := []rune(latter)
		switch {
		case len(newWord) > 2 && unicode.IsDigit(newWord[1]) || unicode.IsDigit(newWord[0]):
			return "", ErrInvalidString
		case len(newWord) > 1:
			n, err := strconv.Atoi(string(newWord[len(newWord)-1]))
			if err != nil {
				return "", fmt.Errorf("error atoi : %w", err)
			}
			newLetters := strings.Repeat(string(newWord[0:len(newWord)-1]), n)
			builder.WriteString(newLetters)
		default:
			builder.WriteString(string(newWord[0]))
		}
	}
	result := builder.String()
	return result, nil
}
