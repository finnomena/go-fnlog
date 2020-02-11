package fnlog

import (
	"context"
	"io"
	"os"
	"time"
)

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

	SetLevel(level LogLevel)
	SetContext(ctx context.Context)
	DeleteKey(key interface{})
}

// Formatter - log formatter
type Formatter interface {
	Message(level LogLevel, fieldMap fields, args ...interface{}) string
}

type standard struct {
	Level     LogLevel
	Context   context.Context
	logctx    map[context.Context]fields
	logkey    map[interface{}]fields
	formatter Formatter
	writer    io.Writer
}

// Options - logger options
type Options struct {
	Formatter Formatter
	Writer    io.Writer
}

var standardLoger *standard

func init() {
	standardLoger = new()
}

// NewLogger - get log instance
func NewLogger() Logger {
	return new()
}

func new() *standard {
	return &standard{
		Level:  TraceLevel,
		logctx: make(map[context.Context]fields),
		logkey: make(map[interface{}]fields),
		formatter: &JSONFormatter{
			timeformat: time.RFC3339Nano,
		},
		writer: os.Stdout,
	}
}

// NewLoggerWithOptions - create custom logger
func NewLoggerWithOptions(opts Options) Logger {
	return &standard{
		Level:     TraceLevel,
		logctx:    make(map[context.Context]fields),
		logkey:    make(map[interface{}]fields),
		formatter: opts.Formatter,
		writer:    opts.Writer,
	}
}
