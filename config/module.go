package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		New,
		func(cfg Interface) Log {
			return cfg.Log()
		},
		func(cfg Interface) Database {
			return cfg.Database()
		},
		func(cfg Interface) HTTPServer {
			return cfg.HTTPServer()
		},
		func(cfg Interface) JwtToken {
			return cfg.JwtToken()
		},
		func(cfg Interface) FileStorage {
			return cfg.FileStorage()
		},
	),
)
