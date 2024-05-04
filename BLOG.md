Feb 17, 2024

## Mutating the context of a gRPC Server

It was a bit frustrating that gRPC server contexts are generated from `context.Background`, and not as a child of provided context, so I made a small library that utilitzes an interceptor to mutate the server context.

### Stream Interceptor

By wrapping the stream, the gRPC method `Context` can be overwritten with a method that calls the mutated context.

```go
type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

func newWrappedStream(ss grpc.ServerStream, f CtxMut) grpc.ServerStream {
	return &wrappedStream{ss, f(ss.Context())}
}

func NewGcmStreamInterceptor(f CtxMut) grpc.ServerOption {
	return grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, newWrappedStream(ss, f))
	})
}
```

### Unary Interceptor

This one is much simpler, just return the mutated context.

```go
func NewGcmUnaryInterceptor(f CtxMut) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(f(ctx), req)
	})
}
```