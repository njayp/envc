package gcm

import (
	"context"

	"google.golang.org/grpc"
)

type CtxMut func(context.Context) context.Context

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

func NewGcmUnaryInterceptor(f CtxMut) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(f(ctx), req)
	})
}

func NewGcmServerOpts(f CtxMut) []grpc.ServerOption {
	return []grpc.ServerOption{NewGcmUnaryInterceptor(f), NewGcmStreamInterceptor(f)}
}
