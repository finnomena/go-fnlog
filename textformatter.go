package fnlog

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// TextFormatter - log with json format
type TextFormatter struct {
	Timeformat string
}

// Message - json message
func (p *TextFormatter) Message(level LogLevel, fieldMap fields, args ...interface{}) string {
	_, ctx := args[0].(context.Context)
	if ctx {
		if len(args) != 1 {
			args = args[1:]
		} else {
			args = nil
		}
	}

	msg := time.Now().Format(p.Timeformat)
	msg += logLevelWithColor(level)
	fun, _, _ := ReportCaller(5)
	msg += fun + " -"

	if fieldMap != nil {
		for k, v := range fieldMap {
			msg += fmt.Sprintf(` %v=%v`, k, v)
		}

		msg = strings.TrimSuffix(msg, ",")
	}

	if args != nil {
		msg += fmt.Sprintf(` %+v`, args...)
	}

	msg += "\n"

	return msg
}

func logLevelWithColor(level LogLevel) string {
	color, ok := levelColor[level]
	if !ok {
		color = "1;30"
	}

	return " [\033[" + color + "m" + strings.ToUpper(levelText[level]) + "\033[0m] "
}

var levelColor map[LogLevel]string = map[LogLevel]string{
	TraceLevel: "1;35",
	DebugLevel: "1;34",
	InfoLevel:  "1;32",
	WarnLevel:  "1;33",
	ErrorLevel: "1;91",
	FatalLevel: "1;31",
	PanicLevel: "1;30",
}
