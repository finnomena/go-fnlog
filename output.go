package fnlog

import (
	"context"
	"fmt"
	"io"
	"os"
)

func log(level LogLevel, f fields, args ...interface{}) {
	_, ctx := args[0].(context.Context)
	if ctx {
		if len(args) != 1 {
			args = args[1:]
		} else {
			args = nil
		}
	}
	s := defaultLog(level)
	if f != nil {
		s = logAllField(s, f)
	}
	if args != nil {
		s += fmt.Sprintf(`"msg":"%v",`, args...)
	}
	s = s[:len(s)-1] + "}\n"
	io.WriteString(os.Stdout, s)
}
