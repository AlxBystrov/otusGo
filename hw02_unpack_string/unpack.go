package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	// "unicode"
)

var (
	ErrInvalidStringEnding = errors.New("invalid string ending")
	ErrDigitWrongPlace     = errors.New("invalid string. Digit is placed on wrong place")
	ErrEscapingLimit       = errors.New("only digit and '\\' escaping allowed")
)

func IsDigit(t string) bool {
	if len(t) > 1 {
		return false
	}
	return t[0] >= '0' && t[0] <= '9'
}

func Unpack(t string) (string, error) {
	const escChar string = `\`
	var result strings.Builder

	gliphs := strings.Split(t, "")

	for i := 0; i < len(gliphs); i++ {
		// Check if the char is digit
		if IsDigit(gliphs[i]) {
			return "", ErrDigitWrongPlace
		}
		// Do not check next char at the end of cycle
		if i == len(gliphs)-1 {
			// String ending couldn't be `\`
			if gliphs[i] == escChar {
				return "", ErrInvalidStringEnding
			}
			fmt.Fprintf(&result, "%v", gliphs[i])
			continue
		}
		// Check on escaping case
		if gliphs[i] == escChar {
			// When the next char is not a digit and `\`
			if !IsDigit(gliphs[i+1]) && gliphs[i+1] != escChar {
				return "", ErrEscapingLimit
			}
			// Do not check repeat number when ending of the string
			if i+2 >= len(gliphs) {
				fmt.Fprintf(&result, "%v", gliphs[i+1])
				i++
				continue
			}
			// Check on repeat
			if nextNum, err := strconv.Atoi(string(gliphs[i+2])); err == nil {
				fmt.Fprintf(&result, "%s", strings.Repeat(string(gliphs[i+1]), nextNum))
				i += 2
			} else {
				fmt.Fprintf(&result, "%v", gliphs[i+1])
				i++
			}
			continue
		}

		// Use Repeat when the next char is digit and increase counter
		if num, err := strconv.Atoi(string(gliphs[i+1])); err == nil {
			fmt.Fprintf(&result, "%s", strings.Repeat(string(gliphs[i]), num))
			i++
		} else {
			fmt.Fprintf(&result, "%v", gliphs[i])
		}
	}

	return result.String(), nil
}
