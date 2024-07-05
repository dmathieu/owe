package owegrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/stats"

	"github.com/dmathieu/owe"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type mockSpan struct {
	noop.Span
}

func TestTagRPC(t *testing.T) {
	ctx := trace.ContextWithSpan(context.Background(), mockSpan{})
	h := NewHandler()

	ctx = h.TagRPC(ctx, &stats.RPCTagInfo{})
	assert.Equal(t, mockSpan{}, owe.SpanFromContext(ctx))
}
