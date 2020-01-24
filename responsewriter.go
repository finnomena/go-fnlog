package fnlog

import (
	"net/http"
)

type responseWriter struct {
	code int
	w    http.ResponseWriter
}

func (w *responseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.w.Write(b)
}
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.w.WriteHeader(statusCode)
}
