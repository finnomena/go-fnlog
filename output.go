package fnlog

import (
	"context"
	"io"
	"os"
)

func (s *standard) log(level LogLevel, f fields, args ...interface{}) {
	_, ctx := args[0].(context.Context)
	if ctx {
		if len(args) != 1 {
			args = args[1:]
		} else {
			args = nil
		}
	}

	io.WriteString(os.Stdout, s.formatter.Message(level, f, args...))
}
