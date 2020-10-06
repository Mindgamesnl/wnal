package wio

import "io"

type WrappedWriter struct {
	OnWrite func(written []byte)
	Replaces io.Writer
}

func (w WrappedWriter) Write(p []byte) (n int, err error) {
	w.OnWrite(p)
	return w.Replaces.Write(p)
}