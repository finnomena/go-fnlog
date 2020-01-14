package fnlog_test

import (
	"fmt"
	"testing"

	"github.com/finnomena/go-fnlog"
)

func assertLogLevel(t *testing.T, name string, level fnlog.LogLevel) {
	result, err := fnlog.GetLogLevel(name)
	message := fmt.Sprintf("[%s] result: %d, expect: %d", name, result, level)

	if err == nil && level == result {
		t.Log(message)
	} else {
		t.Error(message)
	}
}

func TestGetLogLevelShouldBeCorrect(t *testing.T) {
	assertLogLevel(t, "trace", fnlog.TraceLevel)
	assertLogLevel(t, "debug", fnlog.DebugLevel)
	assertLogLevel(t, "info", fnlog.InfoLevel)
	assertLogLevel(t, "warn", fnlog.WarnLevel)
	assertLogLevel(t, "error", fnlog.ErrorLevel)
	assertLogLevel(t, "fatal", fnlog.FatalLevel)
	assertLogLevel(t, "panic", fnlog.PanicLevel)
}

func TestGetLogLevelShouldBeError(t *testing.T) {
	_, err := fnlog.GetLogLevel("foo")

	if err != nil {
		t.Log("error as expect")
	} else {
		t.Error("no error")
	}
}
