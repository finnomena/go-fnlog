package fnlog

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"
)

func AddField(ctx context.Context, key string, value interface{}) {
	if ctx == nil {
		return
	}
	if _, ok := logctx[ctx]; ok {
		logctx[ctx][key] = value
		traceID := ctx.Value("request_id")
		if traceID == nil {
			return
		}
		if _, ok := logkey[traceID]; ok {
			logkey[traceID][key] = value
		}
	} else {
		traceID := ctx.Value("request_id")
		if traceID == nil {
			return
		}
		if _, ok := logkey[traceID]; ok {
			logkey[traceID][key] = value
		}
	}
}

func getFields(arg interface{}) fields {
	ctx, ok := arg.(context.Context)
	if v, has := logctx[ctx]; ok && has {
		return v
	}
	if v, has := logkey[arg]; has {
		return v
	}
	return nil
}

var calldepth = 4

func defaultLog(level logLevel) string {
	funcName, fileName, _ := reportCaller(calldepth)
	s := "{"
	s += `"level":"` + levelTxt[level] + `",`
	s += fmt.Sprintf(`"time":"%v",`, time.Now().Format(time.RFC3339Nano))
	s += fmt.Sprintf(`"func":"%v",`, funcName)
	s += fmt.Sprintf(`"file":"%v",`, fileName)
	return s
}

func logAllField(s string, f fields) string {
	for k, v := range f {
		switch v.(type) {
		case int, int8, int16, int32, int64:
			s += fmt.Sprintf(`"%v":%v,`, k, v)
		case uint, uint8, uint16, uint32, uint64:
			s += fmt.Sprintf(`"%v":%v,`, k, v)
		case float32, float64:
			s += fmt.Sprintf(`"%v":%v,`, k, v)
		case error:
			s += fmt.Sprintf(`"%v":"%v",`, k, v.(error).Error())
		case string:
			s += fmt.Sprintf(`"%v":"%v",`, k, v)
		default:
			mar, err := json.Marshal(v)
			if err != nil {
				break
			}
			s += fmt.Sprintf(`"%v":%v,`, k, string(mar))
		}
	}
	return s
}
func reportCaller(calldepth int) (string, string, error) {
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
