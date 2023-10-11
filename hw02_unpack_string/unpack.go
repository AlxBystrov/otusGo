package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidString       = errors.New("invalid string")
	ErrInvalidStringEnding = errors.New("invalid string ending")
)

func Unpack(t string) (string, error) {
	var result strings.Builder

	for i := 0; i < len(t); i++ {
		_, err := strconv.Atoi(string(t[i]))
		// Check if the char is digit
		if err == nil {
			return "", ErrInvalidString
		}
		// Do not check next char at the end of cycle
		if i == len(t)-1 {
			// String ending couldn't be `\`
			if string(t[i]) == `\` {
				return "", ErrInvalidStringEnding
			}
			fmt.Fprintf(&result, "%c", t[i])
			continue
		}
		// Check on escaping case
		if string(t[i]) == `\` {
			// When the next char is not a number and `\`
			if _, err := strconv.Atoi(string(t[i+1])); err != nil && string(t[i+1]) != `\` {
				return "", ErrInvalidString
			}
			// Do not check repeat number when ending of the string
			if i+2 >= len(t) {
				fmt.Fprintf(&result, "%s", string(t[i+1]))
				i++
				continue
			}
			// Check on repeat
			if nextNum, err := strconv.Atoi(string(t[i+2])); err == nil {
				fmt.Fprintf(&result, "%s", strings.Repeat(string(t[i+1]), nextNum))
				i += 2
			} else {
				fmt.Fprintf(&result, "%s", string(t[i+1]))
				i++
			}
			continue
		}

		// Use Repeat when the next char is digit and increase counter
		if num, err := strconv.Atoi(string(t[i+1])); err == nil {
			fmt.Fprintf(&result, "%s", strings.Repeat(string(t[i]), num))
			i++
		} else {
			fmt.Fprintf(&result, "%c", t[i])
		}
	}

	return result.String(), nil
}
