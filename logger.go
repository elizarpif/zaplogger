package zaplogger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func newLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = zapcore.RFC3339TimeEncoder
	consoleEncoder := zapcore.NewJSONEncoder(config)
	defaultLogLevel := zapcore.DebugLevel

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	baseLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return baseLogger
}

func init() {
	baseLogger := newLogger()
	defer baseLogger.Sync() // flushes buffer, if any

	logger = baseLogger.Sugar()
}

func addValuesFromCtxToLogger(ctx context.Context) *zap.SugaredLogger {
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

// With ...
func With(args ...interface{}) {
	logger = logger.With(args...)
}

// Info ...
func Info(ctx context.Context, args ...interface{}) {
	log := addValuesFromCtxToLogger(ctx)

	log.Info(args)
}

// Infof ...
func Infof(ctx context.Context, template string, args ...interface{}) {
	log := addValuesFromCtxToLogger(ctx)

	log.Infof(template, args)
}

// Debug ...
func Debug(ctx context.Context, args ...interface{}) {
	log := addValuesFromCtxToLogger(ctx)

	log.Debug(args)
}

// Debugf ...
func Debugf(ctx context.Context, template string, args ...interface{}) {
	log := addValuesFromCtxToLogger(ctx)

	log.Debugf(template, args)
}

// Error ...
func Error(ctx context.Context, args ...interface{}) {
	log := addValuesFromCtxToLogger(ctx)

	log.Error(args)
}

// Errorf ...
func Errorf(ctx context.Context, template string, args ...interface{}) {
	log := addValuesFromCtxToLogger(ctx)

	log.Errorf(template, args)
}
