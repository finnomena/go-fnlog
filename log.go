package fnlog

import (
	"fmt"
	"io"
	"os"
)

func (s *standard) Info(args ...interface{}) {
	if !s.IsInfoEnabled() {
		return
	}

	s.print(InfoLevel, args...)
}

func (s *standard) IsInfoEnabled() bool {
	return s.Level <= InfoLevel
}

func (s *standard) Debug(args ...interface{}) {
	if !s.IsDebugEnabled() {
		return
	}

	s.print(DebugLevel, args...)
}

func (s *standard) IsDebugEnabled() bool {
	return s.Level <= DebugLevel
}

func (s *standard) Error(args ...interface{}) {
	if !s.IsErrorEnabled() {
		return
	}

	s.print(ErrorLevel, args...)
}

func (s *standard) IsErrorEnabled() bool {
	return s.Level <= ErrorLevel
}

func (s *standard) Panic(args ...interface{}) {
	if !s.IsPanicEnabled() {
		return
	}

	s.print(PanicLevel, args...)

	panic(nil)
}

func (s *standard) IsPanicEnabled() bool {
	return s.Level <= PanicLevel
}

func (s *standard) Trace(args ...interface{}) {
	if !s.IsTraceEnabled() {
		return
	}

	s.print(TraceLevel, args...)
}

func (s *standard) IsTraceEnabled() bool {
	return s.Level <= TraceLevel
}

func (s *standard) Warn(args ...interface{}) {
	if !s.IsWarnEnabled() {
		return
	}

	s.print(WarnLevel, args...)
}

func (s *standard) IsWarnEnabled() bool {
	return s.Level <= WarnLevel
}

func (s *standard) Fatal(args ...interface{}) {
	if !s.IsFatalEnabled() {
		return
	}

	s.print(FatalLevel, args...)

	os.Exit(1)
}

func (s *standard) IsFatalEnabled() bool {
	return s.Level <= FatalLevel
}

func (s *standard) Access(args ...interface{}) {
	if !s.IsInfoEnabled() {
		return
	}

	s.print(accessLevel, args...)
}

// Info - global info log
func Info(args ...interface{}) {
	standardLoger.Info(args...)
}

// Debug - global debug log
func Debug(args ...interface{}) {
	standardLoger.Debug(args...)
}

// Error - global error log
func Error(args ...interface{}) {
	standardLoger.Error(args...)
}

// Panic - global panic log
func Panic(args ...interface{}) {
	standardLoger.Panic(args...)
}

// Trace - global trace log
func Trace(args ...interface{}) {
	standardLoger.Trace(args...)
}

// Warn - global warn log
func Warn(args ...interface{}) {
	standardLoger.Warn(args...)
}

// Fatal - global fatal log
func Fatal(args ...interface{}) {
	standardLoger.Fatal(args...)
}

// Access - global access log
func Access(args ...interface{}) {
	standardLoger.Access(args...)
}

func (s *standard) print(level LogLevel, args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(s.writer, fmt.Sprint(args...)+"\n")
	}

	s.log(level, s.getFields(args[0]), args...)
}

func (s *standard) SetFormatter(formatter Formatter) {
	s.formatter = formatter
}

// SetFormatter - set global formatter
func SetFormatter(formatter Formatter) {
	standardLoger.SetFormatter(formatter)
}
