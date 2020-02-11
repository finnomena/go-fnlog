package fnlog_test

import (
	"fmt"
	"testing"

	"github.com/finnomena/go-fnlog"
)

func assertStatus(t *testing.T, result, expect bool) {
	message := fmt.Sprintf("result: %v, expect: %v", result, expect)
	if expect == result {
		t.Log(message)
	} else {
		t.Error(message)
	}
}

type test struct {
	logger fnlog.Logger
}

func (t *test) print() {
	t.logger.Error("log with method")
}

func TestLog(t *testing.T) {
	fnlog.Info("global")

	logger := fnlog.NewLogger()
	logger.SetLevel(fnlog.TraceLevel)

	logger.Debug("logging with struct")

	var obj = &test{
		logger: logger,
	}

	obj.logger.Warn("log attribute")
	obj.print()
}

func TestIsEnableShouldBeCorrect(t *testing.T) {
	trace := fnlog.NewLogger()
	trace.SetLevel(fnlog.TraceLevel)
	assertStatus(t, trace.IsTraceEnabled(), true)
	assertStatus(t, trace.IsDebugEnabled(), true)
	assertStatus(t, trace.IsInfoEnabled(), true)
	assertStatus(t, trace.IsWarnEnabled(), true)
	assertStatus(t, trace.IsErrorEnabled(), true)
	assertStatus(t, trace.IsFatalEnabled(), true)
	assertStatus(t, trace.IsPanicEnabled(), true)

	debug := fnlog.NewLogger()
	debug.SetLevel(fnlog.DebugLevel)
	assertStatus(t, debug.IsTraceEnabled(), false)
	assertStatus(t, debug.IsDebugEnabled(), true)
	assertStatus(t, debug.IsInfoEnabled(), true)
	assertStatus(t, debug.IsWarnEnabled(), true)
	assertStatus(t, debug.IsErrorEnabled(), true)
	assertStatus(t, debug.IsFatalEnabled(), true)
	assertStatus(t, debug.IsPanicEnabled(), true)

	info := fnlog.NewLogger()
	info.SetLevel(fnlog.InfoLevel)
	assertStatus(t, info.IsTraceEnabled(), false)
	assertStatus(t, info.IsDebugEnabled(), false)
	assertStatus(t, info.IsInfoEnabled(), true)
	assertStatus(t, info.IsWarnEnabled(), true)
	assertStatus(t, info.IsErrorEnabled(), true)
	assertStatus(t, info.IsFatalEnabled(), true)
	assertStatus(t, info.IsPanicEnabled(), true)

	warn := fnlog.NewLogger()
	warn.SetLevel(fnlog.WarnLevel)
	assertStatus(t, warn.IsTraceEnabled(), false)
	assertStatus(t, warn.IsDebugEnabled(), false)
	assertStatus(t, warn.IsInfoEnabled(), false)
	assertStatus(t, warn.IsWarnEnabled(), true)
	assertStatus(t, warn.IsErrorEnabled(), true)
	assertStatus(t, warn.IsFatalEnabled(), true)
	assertStatus(t, warn.IsPanicEnabled(), true)

	err := fnlog.NewLogger()
	err.SetLevel(fnlog.ErrorLevel)
	assertStatus(t, err.IsTraceEnabled(), false)
	assertStatus(t, err.IsDebugEnabled(), false)
	assertStatus(t, err.IsInfoEnabled(), false)
	assertStatus(t, err.IsWarnEnabled(), false)
	assertStatus(t, err.IsErrorEnabled(), true)
	assertStatus(t, err.IsFatalEnabled(), true)
	assertStatus(t, err.IsPanicEnabled(), true)

	fatal := fnlog.NewLogger()
	fatal.SetLevel(fnlog.FatalLevel)
	assertStatus(t, fatal.IsTraceEnabled(), false)
	assertStatus(t, fatal.IsDebugEnabled(), false)
	assertStatus(t, fatal.IsInfoEnabled(), false)
	assertStatus(t, fatal.IsWarnEnabled(), false)
	assertStatus(t, fatal.IsErrorEnabled(), false)
	assertStatus(t, fatal.IsFatalEnabled(), true)
	assertStatus(t, fatal.IsPanicEnabled(), true)

	panic := fnlog.NewLogger()
	panic.SetLevel(fnlog.PanicLevel)
	assertStatus(t, panic.IsTraceEnabled(), false)
	assertStatus(t, panic.IsDebugEnabled(), false)
	assertStatus(t, panic.IsInfoEnabled(), false)
	assertStatus(t, panic.IsWarnEnabled(), false)
	assertStatus(t, panic.IsErrorEnabled(), false)
	assertStatus(t, panic.IsFatalEnabled(), false)
	assertStatus(t, panic.IsPanicEnabled(), true)

	off := fnlog.NewLogger()
	off.SetLevel(fnlog.OffLevel)
	assertStatus(t, off.IsTraceEnabled(), false)
	assertStatus(t, off.IsDebugEnabled(), false)
	assertStatus(t, off.IsInfoEnabled(), false)
	assertStatus(t, off.IsWarnEnabled(), false)
	assertStatus(t, off.IsErrorEnabled(), false)
	assertStatus(t, off.IsFatalEnabled(), false)
	assertStatus(t, off.IsPanicEnabled(), false)
}

func BenchmarkCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fnlog.GetCaller()
	}
}

func BenchmarkReportCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fnlog.ReportCaller(4)
	}
}
