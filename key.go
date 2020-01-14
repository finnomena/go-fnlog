package fnlog

import "context"

type fields map[string]interface{}

var logctx map[context.Context]fields
var logkey map[interface{}]fields

func init() {
	logctx = make(map[context.Context]fields)
	logkey = make(map[interface{}]fields)
}
