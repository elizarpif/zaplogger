package zaplogger

import "context"

type inCtx int

const (
	requestIdKey inCtx = iota
	traceIdKey
)

// WithReqId returns a context which knows its request ID
func WithReqId(ctx context.Context, rqId string) context.Context {
	return context.WithValue(ctx, requestIdKey, rqId)
}

// WithTraceId returns a context which knows its trace ID
func WithTraceId(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIdKey, traceID)
}
