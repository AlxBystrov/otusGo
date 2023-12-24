package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	type testCase struct {
		name     string
		dir      string
		expected Environment
	}

	testCases := []testCase{
		{
			name: "basic case",
			dir:  "env",
			expected: Environment{
				"BAR":   EnvValue{Value: "bar", NeedRemove: false},
				"EMPTY": EnvValue{Value: "", NeedRemove: false},
				"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
				"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
				"UNSET": EnvValue{Value: "", NeedRemove: true},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			testEnv, err := ReadDir("testdata/env")
			require.Equal(t, nil, err)
			require.Equal(t, test.expected, testEnv)
		})
	}
}
