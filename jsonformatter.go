package fnlog

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// JSONFormatter - log with json format
type JSONFormatter struct {
	timeformat string
}

// Message - json message
func (p *JSONFormatter) Message(level LogLevel, fieldMap fields, args ...interface{}) string {
	_, ctx := args[0].(context.Context)
	if ctx {
		if len(args) != 1 {
			args = args[1:]
		} else {
			args = nil
		}
	}

	msg := p.defaultLog(level)
	if fieldMap != nil {
		msg = p.logWithField(msg, fieldMap)
	}

	if args != nil {
		msg += fmt.Sprintf(`"message":"%v",`, args...)
	}

	msg = msg[:len(msg)-1] + "}\n"

	return msg

}

func (p *JSONFormatter) defaultLog(level LogLevel) string {
	needCaller := false
	msg := "{"

	if level != accessLevel {
		needCaller = true
	}

	msg += `"serverity":"` + levelText[level] + `",`
	msg += fmt.Sprintf(`"timestamp":"%v",`, time.Now().Format(p.timeformat))

	if needCaller {
		if needCaller {
			msg += fmt.Sprintf(`"caller":"%v",`, GetCaller())
		}
	}

	return msg
}

func (p *JSONFormatter) logWithField(s string, f fields) string {
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
