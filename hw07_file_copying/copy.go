package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const chunkSize = 1

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type ProgressBar struct {
	total    int
	position int
	percent  int
}

func (progressBar *ProgressBar) Init(total int) {
	progressBar.total = total
}

func (progressBar *ProgressBar) Add(val int) {
	progressBar.position += val
	progressBar.percent = progressBar.position * 100 / progressBar.total
	fmt.Printf("\rProgress: [%s%s] Bytes: %v, Percents: %v",
		strings.Repeat("+", progressBar.percent),
		strings.Repeat("-", 100-progressBar.percent),
		progressBar.position,
		progressBar.position*100/progressBar.total,
	)
}

func (progressBar *ProgressBar) Close() {
	fmt.Printf("\n")
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	var pb ProgressBar
	source, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer source.Close()

	sourceStat, err := source.Stat()
	if err != nil {
		return err
	}

	if sourceStat.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > sourceStat.Size() {
		return ErrOffsetExceedsFileSize
	}
	if limit > sourceStat.Size() || limit == 0 {
		limit = sourceStat.Size()
	}
	if limit > sourceStat.Size()-offset {
		limit = sourceStat.Size() - offset
	}

	dest, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	pb.Init(int(limit))
	source.Seek(offset, 0)
	for i := offset; i < limit+offset; i += chunkSize {
		bytesCopied, err := io.CopyN(dest, source, chunkSize)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		pb.Add(int(bytesCopied))
	}
	pb.Close()

	return nil
}
