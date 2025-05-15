package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		New,
		func(cfg Interface) LogConfig {
			return cfg.LogConfig()
		},
		func(cfg Interface) DatabaseConfig {
			return cfg.DatabaseConfig()
		},
		func(cfg Interface) HTTPServerConfig {
			return cfg.HTTPServerConfig()
		},
		func(cfg Interface) JwtTokenConfig {
			return cfg.JwtTokenConfig()
		},
		func(cfg Interface) FileStorageConfig {
			return cfg.FileStorageConfig()
		},
		func(cfg Interface) SentryConfig {
			return cfg.SentryConfig()
		},
	),
)
