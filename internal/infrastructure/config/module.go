package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		New,
		func(cfg ConfigInterface) DatabaseConfig {
			return cfg.GetDatabaseConfig()
		},
		func(cfg ConfigInterface) HttpServerConfig {
			return cfg.GetHttpServerConfig()
		},
	),
)
