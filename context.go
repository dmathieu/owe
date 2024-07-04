package owe

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type traceContextKeyType int

const currentSpanKey traceContextKeyType = iota

var noopSpanInstance trace.Span = noop.Span{}

// ContextWithSpan returns a copy of parent with span set as the current Span.
func ContextWithSpan(parent context.Context, span trace.Span) context.Context {
	return context.WithValue(parent, currentSpanKey, span)
}

// SpanFromContext returns the current Span from ctx.
//
// If no Span is currently set in ctx an implementation of a Span that
// performs no operations is returned.
func SpanFromContext(ctx context.Context) trace.Span {
	if ctx == nil {
		return noopSpanInstance
	}
	if span, ok := ctx.Value(currentSpanKey).(trace.Span); ok {
		return span
	}
	return noopSpanInstance
}
