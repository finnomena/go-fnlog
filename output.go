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
	msg := defaultLog(level)
	if f != nil {
		msg = logAllField(msg, f)
	}
	if args != nil {
		msg += fmt.Sprintf(`"message":"%v",`, args...)
	}
	msg = msg[:len(msg)-1] + "}\n"
	io.WriteString(os.Stdout, msg)
}
