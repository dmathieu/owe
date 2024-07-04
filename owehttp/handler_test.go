package owehttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dmathieu/owe"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type mockSpan struct {
	noop.Span
}

type createSpanHandler struct {
	next http.Handler
}

func (csh createSpanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	span := mockSpan{}
	ctx := trace.ContextWithSpan(r.Context(), span)

	csh.next.ServeHTTP(w, r.WithContext(ctx))
}

func TestHandler(t *testing.T) {
	for _, tt := range []struct {
		name    string
		handler func(*testing.T) http.Handler

		wantStatusCode int
	}{
		{
			name: "with a parent span",
			handler: func(t *testing.T) http.Handler {
				return createSpanHandler{
					next: NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						span := owe.SpanFromContext(r.Context())
						assert.Equal(t, span, mockSpan{})
					})),
				}
			},

			wantStatusCode: http.StatusOK,
		},
		{
			name: "without a parent span",
			handler: func(t *testing.T) http.Handler {
				return NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					span := owe.SpanFromContext(r.Context())
					assert.NotEqual(t, span, mockSpan{})
				}))
			},

			wantStatusCode: http.StatusOK,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "http://localhost/", nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			tt.handler(t).ServeHTTP(rr, r)

			assert.Equal(t, tt.wantStatusCode, rr.Result().StatusCode)
		})
	}
}
