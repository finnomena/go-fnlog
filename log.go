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

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := s.getFields(args[0])
	log(InfoLevel, f, args...)
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

	f := s.getFields(args[0])
	log(DebugLevel, f, args...)
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

	f := s.getFields(args[0])
	log(ErrorLevel, f, args...)
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

	f := s.getFields(args[0])
	log(PanicLevel, f, args...)

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

	f := s.getFields(args[0])
	log(TraceLevel, f, args...)
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

	f := s.getFields(args[0])
	log(WarnLevel, f, args...)
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

	f := s.getFields(args[0])
	log(FatalLevel, f, args...)

	os.Exit(1)
}

func (s *standard) IsFatalEnabled() bool {
	return s.Level <= FatalLevel
}

func (s *standard) Access(args ...interface{}) {
	if !s.IsInfoEnabled() {
		return
	}

	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}

	f := s.getFields(args[0])
	log(accessLevel, f, args...)
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
