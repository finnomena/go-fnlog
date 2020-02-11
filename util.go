package fnlog

import (
	"fmt"
	"runtime"
	"strings"
)

var calldepth = 4

// ReportCaller - get function and file of caller
func ReportCaller(calldepth int) (string, string, error) {
	pc, fileDir, lineNum, ok := runtime.Caller(calldepth)
	if !ok {
		return "", "", nil
	}

	fr := runtime.CallersFrames([]uintptr{pc})
	frame, _ := fr.Next()

	s := strings.Split(frame.Function, ".")
	funcname := s[len(s)-1]

	format := fmt.Sprintf("%s:%d", fileDir, lineNum)
	return fmt.Sprintf("%s", funcname), format, nil
}

// GetCaller - get caller reference
func GetCaller() string {
	pc, _, _, ok := runtime.Caller(calldepth)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}

	return ""
}
