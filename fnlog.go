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
	standardLoger = &standard{
		Level:  InfoLevel,
		logctx: make(map[context.Context]fields),
		logkey: make(map[interface{}]fields),
	}
}
