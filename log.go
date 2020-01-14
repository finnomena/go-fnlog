package fnlog

import (
	"context"
	"fmt"
	"io"
	"os"
)

type standard struct {
	Level   LogLevel
	Context context.Context
}

// LogWithoutContext - get log instance
func LogWithoutContext(level LogLevel) Logger {
	return &standard{Level: level}
}

func (s *standard) Info(args ...interface{}) {
	if !s.IsInfoEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
}

func (s *standard) IsInfoEnabled() bool {
	return s.Level <= InfoLevel
}

func (s *standard) Debug(args ...interface{}) {
	if !s.IsDebugEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
}

func (s *standard) IsDebugEnabled() bool {
	return s.Level <= DebugLevel
}

func (s *standard) Error(args ...interface{}) {
	if !s.IsErrorEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
}

func (s *standard) IsErrorEnabled() bool {
	return s.Level <= ErrorLevel
}

func (s *standard) Panic(args ...interface{}) {
	if !s.IsPanicEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
	panic(nil)
}

func (s *standard) IsPanicEnabled() bool {
	return s.Level <= PanicLevel
}

func (s *standard) Trace(args ...interface{}) {
	if !s.IsTraceEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
}

func (s *standard) IsTraceEnabled() bool {
	return s.Level <= TraceLevel
}

func (s *standard) Warn(args ...interface{}) {
	if !s.IsWarnEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
}

func (s *standard) IsWarnEnabled() bool {
	return s.Level <= WarnLevel
}

func (s *standard) Fatal(args ...interface{}) {
	if !s.IsFatalEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := getFields(args[0])
	log(s.Level, f, args...)
	os.Exit(1)
}

func (s *standard) IsFatalEnabled() bool {
	return s.Level <= FatalLevel
}
