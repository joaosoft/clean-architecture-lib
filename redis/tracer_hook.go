package redis

import (
	"context"
	"fmt"
	"net"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/redis/go-redis/v9"
)

// tracerHook
type tracerHook struct {
	spanName string
}

// newTracerHook created a new tracer hook
func newTracerHook(spanName string) redis.Hook {
	return &tracerHook{
		spanName: spanName,
	}
}

// DialHook dial hook
func (th *tracerHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

// ProcessHook process hook
func (th *tracerHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		var span opentracing.Span
		span, ctx = opentracing.StartSpanFromContext(ctx, th.spanName)
		span.SetTag("cmd", cmd.Args())
		defer span.Finish()

		if err := cmd.Err(); err != nil {
			ext.Error.Set(span, true)
			span.SetTag("error", err.Error())
		}

		return next(ctx, cmd)
	}
}

// ProcessPipelineHook process pipeline hook
func (th *tracerHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		var span opentracing.Span
		span, ctx = opentracing.StartSpanFromContext(ctx, th.spanName)
		defer span.Finish()

		for i, cmd := range cmds {
			span.SetTag(fmt.Sprintf("cmd.%d", i+1), cmd.Args())
			if err := cmd.Err(); err != nil {
				ext.Error.Set(span, true)
				span.SetTag(fmt.Sprintf("error.%d", i+1), err.Error())
			}
		}

		return next(ctx, cmds)
	}
}
