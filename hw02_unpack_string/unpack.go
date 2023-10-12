package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidStringEnding = errors.New("invalid string ending")
	ErrDigitWrongPlace     = errors.New("invalid string. Digit is placed on wrong place")
	ErrEscapingLimit       = errors.New("only digit and '\\' escaping allowed")
)

func Unpack(t string) (string, error) {
	const escChar rune = 92
	var result strings.Builder

	runes := []rune(t)

	for i := 0; i < len(runes); i++ {
		// Check if the char is digit
		if unicode.IsDigit(runes[i]) {
			return "", ErrDigitWrongPlace
		}
		// Do not check next char at the end of cycle
		if i == len(runes)-1 {
			// String ending couldn't be `\`
			if runes[i] == escChar {
				return "", ErrInvalidStringEnding
			}
			fmt.Fprintf(&result, "%c", runes[i])
			continue
		}
		// Check on escaping case
		if runes[i] == escChar {
			// When the next char is not a digit and `\`
			if !unicode.IsDigit(runes[i+1]) && runes[i+1] != escChar {
				return "", ErrEscapingLimit
			}
			// Do not check repeat number when ending of the string
			if i+2 >= len(runes) {
				fmt.Fprintf(&result, "%c", runes[i+1])
				i++
				continue
			}
			// Check on repeat
			if nextNum, err := strconv.Atoi(string(runes[i+2])); err == nil {
				fmt.Fprintf(&result, "%s", strings.Repeat(string(runes[i+1]), nextNum))
				i += 2
			} else {
				fmt.Fprintf(&result, "%c", runes[i+1])
				i++
			}
			continue
		}

		// Use Repeat when the next char is digit and increase counter
		if num, err := strconv.Atoi(string(runes[i+1])); err == nil {
			fmt.Fprintf(&result, "%s", strings.Repeat(string(runes[i]), num))
			i++
		} else {
			fmt.Fprintf(&result, "%c", runes[i])
		}
	}

	return result.String(), nil
}
