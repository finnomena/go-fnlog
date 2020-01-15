package fnlog

import "context"

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

type standard struct {
	Level   LogLevel
	Context context.Context
	logctx  map[context.Context]fields
	logkey  map[interface{}]fields
}

const requestID string = "request_id"

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
	}
}
