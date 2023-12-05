package main

import (
	"crypto/md5"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	type testCase struct {
		name   string
		from   string
		to     string
		etalon string
		offset int64
		limit  int64
	}

	testCases := []testCase{
		{name: "offset 0 limit 0", from: "./testdata/input.txt", to: "/tmp/tempfile", etalon: "./testdata/out_offset0_limit0.txt", offset: 0, limit: 0},
		{name: "offset 0 limit 10", from: "./testdata/input.txt", to: "/tmp/tempfile", etalon: "./testdata/out_offset0_limit10.txt", offset: 0, limit: 10},
		{name: "0ffset 0 limit 1000", from: "./testdata/input.txt", to: "/tmp/tempfile", etalon: "./testdata/out_offset0_limit1000.txt", offset: 0, limit: 1000},
		{name: "0ffset 0 limit 10000", from: "./testdata/input.txt", to: "/tmp/tempfile", etalon: "./testdata/out_offset0_limit10000.txt", offset: 0, limit: 10000},
		{name: "0ffset 100 limit 1000", from: "./testdata/input.txt", to: "/tmp/tempfile", etalon: "./testdata/out_offset100_limit1000.txt", offset: 100, limit: 1000},
		{name: "0ffset 6000 limit 1000", from: "./testdata/input.txt", to: "/tmp/tempfile", etalon: "./testdata/out_offset6000_limit1000.txt", offset: 6000, limit: 1000},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, nil, Copy(test.from, test.to, test.offset, test.limit), "simple copy return an error")

			dest, err := os.Open(test.to)
			require.Equal(t, nil, err, "destination file couldn't be open")

			etalon, _ := os.Open(test.etalon)
			destHash := md5.New()
			if _, err := io.Copy(destHash, dest); err != nil {
				t.Logf("Failed to get hash for destination %s\n", test.to)
				t.Fail()
			}
			etalonHash := md5.New()
			if _, err := io.Copy(etalonHash, etalon); err != nil {
				t.Logf("Failed to get hash for etalon %s\n", test.etalon)
				t.Fail()
			}

			require.Equal(t, etalonHash.Sum(nil), destHash.Sum(nil))
			etalon.Close()
			dest.Close()

		})
	}
}
