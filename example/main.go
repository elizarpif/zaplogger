package main

import (
	"context"

	"github.com/elizarpif/zaplogger"
)

/*
OUTPUT:

{"level":"info","ts":"2022-10-06T22:38:44+02:00","caller":"zaplogger/logger.go:65","msg":"hello [liza]"}
{"level":"info","ts":"2022-10-06T22:38:44+02:00","caller":"zaplogger/logger.go:58","msg":"[hello 2 j fff]"}
{"level":"info","ts":"2022-10-06T22:38:44+02:00","caller":"zaplogger/logger.go:58","msg":"[hello 3]","trace_id":"trace_id"}
{"level":"info","ts":"2022-10-06T22:38:44+02:00","caller":"zaplogger/logger.go:58","msg":"[info]","field":"field12345","trace_id":"trace_id"}
{"level":"info","ts":"2022-10-06T22:38:44+02:00","caller":"zaplogger/logger.go:58","msg":"[info]","field":"field12345"}
*/

func main() {
	ctx := context.Background()

	zaplogger.Infof(ctx, "hello %v", "liza")

	zaplogger.Info(ctx, "hello 2", "j", "fff")

	newCtx := zaplogger.WithTraceId(ctx, "trace_id")

	zaplogger.Info(newCtx, "hello 3")
	newCtx.Done()

	s := server{}
	zaplogger.With("field", "field12345")
	s.logWithInfo(newCtx)
	s.logWithInfo(ctx)
}

type server struct {
	variable int64
}

func (s *server) logWithInfo(ctx context.Context) {
	zaplogger.Info(ctx, "info")
}
