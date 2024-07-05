package owehttp_test

import (
	"io"
	"log"
	"net/http"

	"github.com/dmathieu/owe"
	"github.com/dmathieu/owe/owehttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
)

func ExampleNewHandler() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// Always retrieve the wide event span, which was defined by the otelhttp
		// middleware.
		span := owe.SpanFromContext(req.Context())
		// Wide event spans are meant to hold lots of attributes. So let's add some
		// data to it!
		span.SetAttributes(attribute.String("user.id", "42"))

		_, _ = io.WriteString(w, "Hello, world!\n")
	}

	handler := otelhttp.NewHandler(
		owehttp.NewHandler(
			http.HandlerFunc(helloHandler),
		), "Hello")

	http.Handle("/hello", handler)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}
