package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	type testCase struct {
		name          string
		cmd           []string
		env           Environment
		expectedExist bool
		expectedValue string
		expectedCode  int
	}

	testCases := []testCase{
		{
			name:          "check env created",
			cmd:           []string{"echo", ""},
			env:           Environment{"MY_ENV": EnvValue{Value: "true", NeedRemove: false}},
			expectedExist: true,
			expectedValue: "true",
			expectedCode:  0,
		},
		{
			name:          "check env fixed",
			cmd:           []string{"echo", ""},
			env:           Environment{"MY_ENV": EnvValue{Value: "second value", NeedRemove: false}},
			expectedExist: true,
			expectedValue: "second value",
			expectedCode:  0,
		},
		{
			name:          "env unset",
			cmd:           []string{"echo", ""},
			env:           Environment{"MY_ENV": EnvValue{Value: "unset value", NeedRemove: true}},
			expectedExist: false,
			expectedValue: "",
			expectedCode:  0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			resultCode := RunCmd(test.cmd, test.env)
			require.Equal(t, test.expectedCode, resultCode)
			value, exist := os.LookupEnv("MY_ENV")
			require.Equal(t, test.expectedValue, value)
			require.Equal(t, test.expectedExist, exist)
		})
	}
}
