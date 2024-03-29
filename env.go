package gcm

import (
	"context"

	"github.com/codingconcepts/env"
)

type envKey struct{}

func WithEnv[T any](ctx context.Context) context.Context {
	var e T
	err := env.Set(&e)
	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, envKey{}, e)
}

func GetEnv[T any](ctx context.Context) T {
	e, ok := ctx.Value(envKey{}).(T)
	if !ok {
		panic("no value found in ctx for envcKey")
	}
	return e
}
