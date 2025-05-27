package logger

import (
	"github.com/felipeversiane/donation-server/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func(cfg config.LogConfig) Interface {
		return New(cfg)
	}),
)
