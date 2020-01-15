package fnlog

import (
	"errors"
	"strings"
)

// LogLevel - type of log level
type LogLevel int32

// Log Level constant
const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	OffLevel
)

var levelText map[LogLevel]string = map[LogLevel]string{
	TraceLevel: "trace",
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warn",
	ErrorLevel: "error",
	FatalLevel: "fatal",
	PanicLevel: "panic",
	OffLevel:   "off",
}

// Logger - interface
type Logger interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	IsPanicEnabled() bool
	IsFatalEnabled() bool
	IsErrorEnabled() bool
	IsWarnEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsTraceEnabled() bool
}

// GetLogLevel - return int of log level, default is off
func GetLogLevel(level string) (LogLevel, error) {
	switch strings.ToLower(level) {
	case "trace":
		return TraceLevel, nil
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "off":
		return OffLevel, nil
	default:
		return OffLevel, errors.New("log is not invalid or not set")
	}
}
