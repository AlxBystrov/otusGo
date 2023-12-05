package main

import (
	"errors"
	"flag"
	"log/slog"
)

var (
	from, to       string
	limit, offset  int64
	ErrMissingArgs = errors.New("missing required arguments: [from, to]")
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	if from == "" || to == "" {
		slog.Error("empty argument", "error", ErrMissingArgs)
		return
	}
	if err := Copy(from, to, offset, limit); err != nil {
		slog.Error("failed to copy file", "error", err)
	}
}
