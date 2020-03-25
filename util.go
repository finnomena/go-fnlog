package fnlog

import (
	"fmt"
	"runtime"
	"strings"
)

var calldepthDefault = 7

// ReportCaller - get function and file of caller
func ReportCaller(calldepth *int) (string, string, error) {
	if calldepth == nil {
		calldepth = &calldepthDefault
	}
	pc, fileDir, lineNum, ok := runtime.Caller(*calldepth)
	if !ok {
		return "", "", nil
	}

	fr := runtime.CallersFrames([]uintptr{pc})
	frame, _ := fr.Next()

	ss := strings.Split(frame.Function, ".")
	funcname := ss[len(ss)-1]

	format := fmt.Sprintf("%s:%d", fileDir, lineNum)
	return fmt.Sprintf("%s", funcname), format, nil
}

// GetCaller - get caller reference
func GetCaller(calldepth *int) string {
	if calldepth == nil {
		calldepth = &calldepthDefault
	}
	pc, _, _, ok := runtime.Caller(*calldepth)
	details := runtime.FuncForPC(pc)
	_, line := details.FileLine(pc)
	if ok && details != nil {
		return details.Name() + ":" + fmt.Sprint(line)
	}

	return ""
}
