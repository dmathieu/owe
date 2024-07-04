package owe

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type testSpan struct {
	noop.Span

	ID     byte
	Remote bool
}

func (s testSpan) SpanContext() trace.SpanContext {
	return trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: [16]byte{1},
		SpanID:  [8]byte{s.ID},
		Remote:  s.Remote,
	})
}

var (
	emptySpan = noop.Span{}
	localSpan = testSpan{ID: 1, Remote: false}
)

func TestSpanContext(t *testing.T) {
	testCases := []struct {
		name         string
		context      context.Context
		expectedSpan trace.Span
	}{
		{
			name:         "empty context",
			context:      nil,
			expectedSpan: emptySpan,
		},
		{
			name:         "background context",
			context:      context.Background(),
			expectedSpan: emptySpan,
		},
		{
			name:         "local span",
			context:      ContextWithSpan(context.Background(), localSpan),
			expectedSpan: localSpan,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSpan, SpanFromContext(tc.context))

			// Check that SpanFromContext does not produce any heap allocation.
			assert.Equal(t, 0.0, testing.AllocsPerRun(5, func() {
				SpanFromContext(tc.context)
			}), "SpanFromContext allocs")
		})
	}
}
