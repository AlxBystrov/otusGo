package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	type testCase struct {
		name         string
		cmd          []string
		env          Environment
		expectedCode int
	}

	testCases := []testCase{
		{
			name:         "echo env",
			cmd:          []string{"env"},
			env:          Environment{"MY_ENV": EnvValue{Value: "true", NeedRemove: false}},
			expectedCode: 0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			resultCode := RunCmd(test.cmd, test.env)
			require.Equal(t, test.expectedCode, resultCode)
		})
	}
}
