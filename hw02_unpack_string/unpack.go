package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	sb := &strings.Builder{}
	var lastSymbol string
	var shielded bool
	for _, r := range input {
		if val, err := strconv.Atoi(string(r)); !shielded && err == nil {
			if lastSymbol == "" {
				return "", ErrInvalidString
			}
			if _, err := sb.WriteString(strings.Repeat(lastSymbol, val)); err != nil {
				return "", err
			}
			lastSymbol = ""
			continue
		}
		if shielded {
			if !unicode.IsDigit(r) && r != '\\' {
				return "", ErrInvalidString
			}
			shielded = false
			lastSymbol = string(r)
			continue
		}
		if lastSymbol != "" {
			if _, err := sb.WriteString(lastSymbol); err != nil {
				return "", err
			}
		}
		if r == '\\' {
			shielded = true
			lastSymbol = ""
			continue
		}
		lastSymbol = string(r)
	}
	if shielded {
		return "", ErrInvalidString
	}
	if lastSymbol != "" {
		if _, err := sb.WriteString(lastSymbol); err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}
