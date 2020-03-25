package fnlog

import (
	"fmt"
	"strings"
	"time"
)

// TextFormatter - log with json format
type TextFormatter struct {
	Timeformat string
	Delimiter  string
	CallDepth  *int
}

// Message - json message
func (p *TextFormatter) Message(level LogLevel, fieldMap fields, args ...interface{}) string {
	msg := "\033[0;90m" + time.Now().Format(p.Timeformat) + "\033[0m"
	msg += logLevelWithColor(level)
	fun, _, _ := ReportCaller(p.CallDepth)
	msg += "\033[0;96m" + fun + "\033[0m \033[0;90m-\033[0m"

	if fieldMap != nil {
		for k, v := range fieldMap {
			msg += fmt.Sprintf(` %v=%v`, k, v)
		}

		msg = strings.TrimSuffix(msg, ",")
	}

	if args != nil {
		delimiter := p.Delimiter
		if delimiter == "" {
			delimiter = " "
		}
		prefix := "%+v" + delimiter
		prefix = strings.Repeat(prefix, len(args))
		msg += fmt.Sprintf(fmt.Sprintf(` %s`, prefix[:len(prefix)-len(delimiter)]), args...)
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
