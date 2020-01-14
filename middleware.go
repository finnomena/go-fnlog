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
			ctx := context.WithValue(r.Context(), "request_id", traceID) // nolint
			r = r.WithContext(ctx)

			logctx[ctx] = make(map[string]interface{})
			logkey[traceID] = make(map[string]interface{})

			AddField(ctx, "ip", r.RemoteAddr)
			AddField(ctx, "method", r.Method)
			AddField(ctx, "uri", r.RequestURI)
			AddField(ctx, "request_id", traceID)
			AddField(ctx, "user_agent", r.UserAgent())
			AddField(ctx, "latency", time.Since(now).Nanoseconds())
			next.ServeHTTP(w, r)

			LogWithoutContext(InfoLevel).Info(ctx)

			delete(logctx, ctx)
			delete(logkey, traceID)
		})
	}
}
