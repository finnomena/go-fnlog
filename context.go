package fnlog

import "context"

func (s *standard) SetContext(ctx context.Context) {
	s.Context = ctx
	s.logctx[ctx] = fields{}
	s.logkey[ctx.Value(requestID)] = fields{}
}

func SetContext(ctx context.Context) {
	standardLoger.SetContext(ctx)
}

func DeleteKey(ctx context.Context) {
	standardLoger.DeleteKey(ctx)
}

func (s *standard) DeleteKey(key interface{}) {
	c, ok := key.(context.Context)
	if ok {
		delete(s.logctx, c)
	}
	delete(s.logkey, requestID)
}
