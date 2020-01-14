package fnlog_test

import (
	"fmt"
	"testing"

	"github.com/finnomena/go-fnlog"
)

func assertStatus(t *testing.T, result, expect bool) {
	message := fmt.Sprintf("esult: %v, expect: %v", result, expect)
	if expect == result {
		t.Log(message)
	} else {
		t.Error(message)
	}
}

func TestIsEnableShouldBeCorrect(t *testing.T) {
	trace := fnlog.LogWithoutContext(fnlog.TraceLevel)
	assertStatus(t, trace.IsTraceEnabled(), true)
	assertStatus(t, trace.IsDebugEnabled(), true)
	assertStatus(t, trace.IsInfoEnabled(), true)
	assertStatus(t, trace.IsWarnEnabled(), true)
	assertStatus(t, trace.IsErrorEnabled(), true)
	assertStatus(t, trace.IsFatalEnabled(), true)
	assertStatus(t, trace.IsPanicEnabled(), true)

	debug := fnlog.LogWithoutContext(fnlog.DebugLevel)
	assertStatus(t, debug.IsTraceEnabled(), false)
	assertStatus(t, debug.IsDebugEnabled(), true)
	assertStatus(t, debug.IsInfoEnabled(), true)
	assertStatus(t, debug.IsWarnEnabled(), true)
	assertStatus(t, debug.IsErrorEnabled(), true)
	assertStatus(t, debug.IsFatalEnabled(), true)
	assertStatus(t, debug.IsPanicEnabled(), true)

	info := fnlog.LogWithoutContext(fnlog.InfoLevel)
	assertStatus(t, info.IsTraceEnabled(), false)
	assertStatus(t, info.IsDebugEnabled(), false)
	assertStatus(t, info.IsInfoEnabled(), true)
	assertStatus(t, info.IsWarnEnabled(), true)
	assertStatus(t, info.IsErrorEnabled(), true)
	assertStatus(t, info.IsFatalEnabled(), true)
	assertStatus(t, info.IsPanicEnabled(), true)

	warn := fnlog.LogWithoutContext(fnlog.WarnLevel)
	assertStatus(t, warn.IsTraceEnabled(), false)
	assertStatus(t, warn.IsDebugEnabled(), false)
	assertStatus(t, warn.IsInfoEnabled(), false)
	assertStatus(t, warn.IsWarnEnabled(), true)
	assertStatus(t, warn.IsErrorEnabled(), true)
	assertStatus(t, warn.IsFatalEnabled(), true)
	assertStatus(t, warn.IsPanicEnabled(), true)

	err := fnlog.LogWithoutContext(fnlog.ErrorLevel)
	assertStatus(t, err.IsTraceEnabled(), false)
	assertStatus(t, err.IsDebugEnabled(), false)
	assertStatus(t, err.IsInfoEnabled(), false)
	assertStatus(t, err.IsWarnEnabled(), false)
	assertStatus(t, err.IsErrorEnabled(), true)
	assertStatus(t, err.IsFatalEnabled(), true)
	assertStatus(t, err.IsPanicEnabled(), true)

	fatal := fnlog.LogWithoutContext(fnlog.FatalLevel)
	assertStatus(t, fatal.IsTraceEnabled(), false)
	assertStatus(t, fatal.IsDebugEnabled(), false)
	assertStatus(t, fatal.IsInfoEnabled(), false)
	assertStatus(t, fatal.IsWarnEnabled(), false)
	assertStatus(t, fatal.IsErrorEnabled(), false)
	assertStatus(t, fatal.IsFatalEnabled(), true)
	assertStatus(t, fatal.IsPanicEnabled(), true)

	panic := fnlog.LogWithoutContext(fnlog.PanicLevel)
	assertStatus(t, panic.IsTraceEnabled(), false)
	assertStatus(t, panic.IsDebugEnabled(), false)
	assertStatus(t, panic.IsInfoEnabled(), false)
	assertStatus(t, panic.IsWarnEnabled(), false)
	assertStatus(t, panic.IsErrorEnabled(), false)
	assertStatus(t, panic.IsFatalEnabled(), false)
	assertStatus(t, panic.IsPanicEnabled(), true)

	off := fnlog.LogWithoutContext(fnlog.OffLevel)
	assertStatus(t, off.IsTraceEnabled(), false)
	assertStatus(t, off.IsDebugEnabled(), false)
	assertStatus(t, off.IsInfoEnabled(), false)
	assertStatus(t, off.IsWarnEnabled(), false)
	assertStatus(t, off.IsErrorEnabled(), false)
	assertStatus(t, off.IsFatalEnabled(), false)
	assertStatus(t, off.IsPanicEnabled(), false)
}
