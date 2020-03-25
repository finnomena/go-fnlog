package fnlog

import (
	"context"
)

type fields map[string]interface{}

const requestID string = "request_id"

// AddField - add key and value
func (s *standard) AddField(ctx context.Context, key string, value interface{}) {
	if ctx == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.logctx[ctx]; ok {
		s.logctx[ctx][key] = value
		traceID := ctx.Value(requestID)
		if traceID == nil {
			return
		}
		if _, ok := s.logkey[traceID]; ok {
			s.logkey[traceID][key] = value
		}
	} else {
		traceID := ctx.Value(requestID)
		if traceID == nil {
			return
		}
		if _, ok := s.logkey[traceID]; ok {
			s.logkey[traceID][key] = value
		}
	}
}

// AddField - add key and value for global standard
func AddField(ctx context.Context, key string, value interface{}) {
	standardLoger.AddField(ctx, key, value)
}

func (s *standard) getFields(arg interface{}) fields {
	ctx, ok := arg.(context.Context)
	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, has := s.logctx[ctx]; ok && has {
		return v
	}
	if v, has := s.logkey[arg]; has {
		return v
	}
	return nil
}
