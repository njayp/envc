package gcm

import (
	"context"
	"testing"
)

type config struct {
	Port int `env:"GRPC_PORT" default:"9090"`
}

func TestWithEnv(t *testing.T) {
	ctx := WithEnv[config](context.Background())
	config := GetEnv[config](ctx)
	if config.Port != 9090 {
		t.Error(config.Port)
	}
}
