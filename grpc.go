package envc

import (
	"context"

	"google.golang.org/grpc"
)

type wrappedStream[T any] struct {
	grpc.ServerStream
}

func (w *wrappedStream[T]) Context() context.Context {
	return WithConfig[T](w.ServerStream.Context())
}

func newWrappedStream[T any](s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream[T]{s}
}

func NewEnvcStreamInterceptor[T any]() grpc.ServerOption {
	return grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, newWrappedStream[T](ss))
	})
}

func NewEnvcUnaryInterceptor[T any]() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(WithConfig[T](ctx), req)
	})
}
