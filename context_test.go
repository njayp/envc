package envc

import (
	"context"
	"testing"
)

type config struct {
	Port int `env:"ENVC_PORT" default:"3"`
}

func TestWithConfig(t *testing.T) {
	ctx := WithConfig[config](context.Background())
	config := GetConfig[config](ctx)
	if config.Port != 3 {
		t.Error(config.Port)
	}
}
