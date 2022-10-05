package zaplogger

import (
	"context"

	"go.uber.org/zap"
)

type inCtx int

const (
	requestIdKey inCtx = iota
	traceIdKey
)

var logger *zap.SugaredLogger

func init() {
	baseLogger, _ := zap.NewProduction()
	defer baseLogger.Sync() // flushes buffer, if any

	logger = baseLogger.Sugar()
}

// WithReqId returns a context which knows its request ID
func WithReqId(ctx context.Context, rqId string) context.Context {
	return context.WithValue(ctx, requestIdKey, rqId)
}

// WithTraceId returns a context which knows its trace ID
func WithTraceId(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIdKey, traceID)
}

func addValuesFromCtx(ctx context.Context) *zap.SugaredLogger {
	newLogger := logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("req_id", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(traceIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("trace_id", ctxSessionId))
		}
	}

	return newLogger
}

func Info(ctx context.Context, args ...interface{}) {
	log := addValuesFromCtx(ctx)

	log.Info(args)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	log := addValuesFromCtx(ctx)

	log.Infof(template, args)
}

func Debug(ctx context.Context, args ...interface{}) {
	log := addValuesFromCtx(ctx)

	log.Debug(args)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	log := addValuesFromCtx(ctx)

	log.Debugf(template, args)
}

func Error(ctx context.Context, args ...interface{}) {
	log := addValuesFromCtx(ctx)

	log.Error(args)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	log := addValuesFromCtx(ctx)

	log.Errorf(template, args)
}
