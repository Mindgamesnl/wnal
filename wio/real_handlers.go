package wio

import (
	"io"
	"os"
)

var (
	RealHandlerSet = handlerSet{
		Out:   os.Stdout,
		Error: os.Stderr,
	}
)

type handlerSet struct {
	Out io.Writer
	Error io.Writer
}
