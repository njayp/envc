package envc

import (
	"context"

	"github.com/codingconcepts/env"
)

type envcKey struct{}

func WithConfig[T any](ctx context.Context) context.Context {
	var config T
	if err := env.Set(&config); err != nil {
		panic(err)
	}

	return context.WithValue(ctx, envcKey{}, config)
}

func GetConfig[T any](ctx context.Context) T {
	config, ok := ctx.Value(envcKey{}).(T)
	if !ok {
		panic("no value found in ctx for envcKey")
	}
	return config
}
