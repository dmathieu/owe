# OpenTelemetry Wide Events for Go

OpenTelemetry makes it easy to retrieve the current span stored in context.

```golang
span := trace.SpanFromContext(ctx)
```

With the notion of [Wide
Events](https://isburmistrov.substack.com/p/all-you-need-is-wide-events-not-metrics)
originally introduces by Charity Majors, you will also want to retrieve that
main span all the time.

And yet, if a child span has been created, you will retrieve that new span and
not the wide event you expected.

This package allows you to define a wide event for a context, which will not
change unless you manually do so. You can then retrieve the proper main span
all the time, even if child spans were created.

```
span := owe.SpanFromContext(ctx)
```

## Usage

In order to have a span to retrieve, one needs to be added to the context.

You can do that manually with the `ContextWithSpan` method:

```
ctx := owe.ContextWithSpan(ctx, span)
```

Or you can use one of the provided handlers that will automatically create a
wide event span.

### net/http

The `owehttp` package allows setting wide events for HTTP servers.

```golang
handler := func(http.ResponseWriter, *http.Request) {
	// the HTTP handler sets my app's behavior
	event := owe.SpanFromContext(r.Context())
}

endpoint := otelhttp.NewHandler(owe.NewHandler(handler), "otelhttp")
server := &http.Server{
	Handler: endpoint,
}
log.Fatal(server.ListenAndServe())
```

### gRPC

The `owegrpc` package allows setting wide events for gRPC servers.

```golang
s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.StatsHandler(owegrpc.NewHandler()),
)

// Register your gRPC services

if err := s.Serve(lis); err != nil {
	log.Fatalf("failed to serve: %v", err)
}
```
