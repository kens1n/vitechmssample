package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func StartSpanForContext(ctx context.Context) context.Context {
	tracer := opentracing.GlobalTracer()
	serverSpan := tracer.StartSpan("hashcode-query")
	return opentracing.ContextWithSpan(ctx, serverSpan)
}

func StartSpanWithRootSpanInContext(ctx context.Context, operationName string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		tracer := opentracing.GlobalTracer()
		childSpan := tracer.StartSpan(
			operationName,
			opentracing.ChildOf(parentSpan.Context()),
		)
		return childSpan
	}

	return nil
}
