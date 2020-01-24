package fnlog

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// LoggingMiddleware - logging middleware
func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			traceID := fmt.Sprintf("%v%v", now.UnixNano(), rand.Int())
			ctx := context.WithValue(r.Context(), requestID, traceID) // nolint
			r = r.WithContext(ctx)

			SetContext(r.Context())

			AddField(ctx, "ip", r.RemoteAddr)
			AddField(ctx, "method", r.Method)
			AddField(ctx, "uri", r.RequestURI)
			AddField(ctx, "request_id", traceID)
			AddField(ctx, "user_agent", r.UserAgent())
			AddField(ctx, "latency", time.Since(now).Nanoseconds())

			fnlogWriter := responseWriter{
				w: w,
			}

			next.ServeHTTP(&fnlogWriter, r)
			AddField(ctx, "status", fnlogWriter.code)
			Access(ctx)
		})
	}
}
