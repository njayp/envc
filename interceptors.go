package gcm

import (
	"context"

	"google.golang.org/grpc"
)

type ctxMut func(context.Context) context.Context

type wrappedStream struct {
	grpc.ServerStream
	f ctxMut
}

func (w *wrappedStream) Context() context.Context {
	return w.f(w.ServerStream.Context())
}

func newWrappedStream(ss grpc.ServerStream, f ctxMut) grpc.ServerStream {
	return &wrappedStream{ss, f}
}

func NewGcmStreamInterceptor(f ctxMut) grpc.ServerOption {
	return grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, newWrappedStream(ss, f))
	})
}

func NewGcmUnaryInterceptor(f ctxMut) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(f(ctx), req)
	})
}

func NewGcmServerOpts(f ctxMut) []grpc.ServerOption {
	return []grpc.ServerOption{NewGcmUnaryInterceptor(f), NewGcmStreamInterceptor(f)}
}
