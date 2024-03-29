package gcm

import (
	"context"

	"github.com/codingconcepts/env"
)

type envKey struct{}

func Env[T any]() T {
	var e T
	err := env.Set(&e)
	if err != nil {
		panic(err)
	}
	return e
}

func WithEnv[T any](ctx context.Context) context.Context {
	return context.WithValue(ctx, envKey{}, Env[T]())
}

func GetEnv[T any](ctx context.Context) T {
	e, ok := ctx.Value(envKey{}).(T)
	if !ok {
		panic("GetEnv error: no value found in ctx")
	}
	return e
}
