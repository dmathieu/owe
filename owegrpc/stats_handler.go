package owegrpc

import (
	"context"

	"google.golang.org/grpc/stats"

	"github.com/dmathieu/owe"
	"go.opentelemetry.io/otel/trace"
)

type handler struct {
}

// NewHandler creates a stats.Handler for a gRPC server.
func NewHandler() stats.Handler {
	return &handler{}
}

// TagConn can attach some information to the given context.
func (h *handler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

// HandleConn processes the Conn stats.
func (h *handler) HandleConn(ctx context.Context, info stats.ConnStats) {
}

// TagRPC can attach some information to the given context.
func (h *handler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	span := trace.SpanFromContext(ctx)
	ctx = owe.ContextWithSpan(ctx, span)

	return ctx
}

// HandleRPC processes the RPC stats.
func (h *handler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
}
