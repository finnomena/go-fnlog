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

	f := s.getFields(args[0])
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

	f := s.getFields(args[0])
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

	f := s.getFields(args[0])
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

	f := s.getFields(args[0])
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

	f := s.getFields(args[0])
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

	f := s.getFields(args[0])
	log(s.Level, f, args...)
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

func Info(args ...interface{}) {
	standardLoger.Info(args...)
}

func Debug(args ...interface{}) {
	standardLoger.Debug(args...)
}

func Error(args ...interface{}) {
	standardLoger.Error(args...)
}

func Panic(args ...interface{}) {
	standardLoger.Panic(args...)
}

func Trace(args ...interface{}) {
	standardLoger.Trace(args...)
}

func Warn(args ...interface{}) {
	standardLoger.Warn(args...)
}

func Fatal(args ...interface{}) {
	standardLoger.Fatal(args...)
}

func Access(args ...interface{}) {
	standardLoger.Access(args...)
}
