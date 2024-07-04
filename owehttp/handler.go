package owehttp

import (
	"net/http"

	"github.com/dmathieu/owe"
	"go.opentelemetry.io/otel/trace"
)

type handler struct {
	next http.Handler
}

func NewHandler(next http.Handler) http.Handler {
	return handler{
		next: next,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	span := trace.SpanFromContext(ctx)
	ctx = owe.ContextWithSpan(ctx, span)

	h.next.ServeHTTP(w, r.WithContext(ctx))
}
