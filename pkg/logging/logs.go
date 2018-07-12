package logging

import (
	"io"
	"os"
	"strings"
)

type writerFunc func(p []byte) (n int, err error)

func (w writerFunc) Write(p []byte) (n int, err error) {
	return w(p)
}

// NewLogWriter returns a io.writer func that conditionally outputs debug flags
func NewLogWriter(debug bool) io.Writer {
	return writerFunc(func(p []byte) (n int, err error) {
		if !debug && strings.Contains(string(p), "[DEBUG]") {
			return 0, nil
		}

		return os.Stdout.Write(p)
	})
}
