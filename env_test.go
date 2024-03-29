package gcm

import (
	"context"
	"testing"
)

type config struct {
	Port int `env:"ENVC_PORT" default:"3"`
}

func TestWithEnv(t *testing.T) {
	ctx := WithEnv[config](context.Background())
	config := GetEnv[config](ctx)
	if config.Port != 3 {
		t.Error(config.Port)
	}
}
