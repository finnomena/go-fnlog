package fnlog

import "context"

func (s *standard) SetContext(ctx context.Context) {
	s.Context = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	s.logctx[ctx] = fields{}
	s.logkey[ctx.Value(requestID)] = fields{}
}

// SetContext - set log context
func SetContext(ctx context.Context) {
	standardLoger.SetContext(ctx)
}

// DeleteKey - delete key by context
func DeleteKey(ctx context.Context) {
	standardLoger.DeleteKey(ctx)
}

func (s *standard) DeleteKey(key interface{}) {
	c, ok := key.(context.Context)
	if !ok {
		delete(s.logkey, c)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.logctx, c)
	delete(s.logkey, c.Value(requestID))
}
