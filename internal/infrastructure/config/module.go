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
		func(cfg ConfigInterface) JwtTokenConfig {
			return cfg.GetJwtTokenConfig()
		},
		func(cfg ConfigInterface) FileStorageConfig {
			return cfg.GetFileStorageConfig()
		},
	),
)
