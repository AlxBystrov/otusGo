package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `a\\b3ac3\\`, expected: `a\bbbaccc\`},
		{input: "вг3ф0он", expected: "вгггон"},
		{input: "১২৩", expected: "১২৩"},
		{input: `১২৩3🍁\6`, expected: "১২৩৩৩🍁6"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidStringEnding(t *testing.T) {
	invalidStrings := []string{`\3abc\`, `la4k0\`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidStringEnding), "actual error %q", err)
		})
	}
}

func TestUnpackDigitWrongPlace(t *testing.T) {
	invalidStrings := []string{"a33abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrDigitWrongPlace), "actual error %q", err)
		})
	}
}

func TestUnpackEscapingLimit(t *testing.T) {
	invalidStrings := []string{`\a33abc`, `a4d5\д`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrEscapingLimit), "actual error %q", err)
		})
	}
}
