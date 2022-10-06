package zaplogger

import (
	"context"
	"testing"
)

func TestWithTraceId(t *testing.T) {
	type args struct {
		ctx     context.Context
		traceID string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "check adding value to ctx",
			args: args{
				context.Background(),
				"some_trace",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := WithTraceId(tt.args.ctx, tt.args.traceID)

			log := addValuesFromCtxToLogger(ctx)
			log.Info("check")
		})
	}
}
