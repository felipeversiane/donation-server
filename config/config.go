package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	Log         LogConfig         `envPrefix:"LOG_"`
	Database    DatabaseConfig    `envPrefix:"DB_"`
	HTTPServer  HTTPServerConfig  `envPrefix:"HTTP_SERVER_"`
	JwtToken    JwtTokenConfig    `envPrefix:"JWT_"`
	FileStorage FileStorageConfig `envPrefix:"FILE_STORAGE_"`
	Sentry      SentryConfig      `envPrefix:"SENTRY_"`
}

type Interface interface {
	DatabaseConfig() DatabaseConfig
	HTTPServerConfig() HTTPServerConfig
	JwtTokenConfig() JwtTokenConfig
	FileStorageConfig() FileStorageConfig
	SentryConfig() SentryConfig
	LogConfig() LogConfig
}

type LogConfig struct {
	Level  string `env:"LOG_LEVEL" default:"info"`
	Path   string `env:"LOG_PATH" default:"logs/app.log"`
	Stdout bool   `env:"LOG_STDOUT" default:"true"`
}

type DatabaseConfig struct {
	Host            string `env:"HOST" envDefault:"localhost"`
	Port            string `env:"PORT" envDefault:"5432"`
	User            string `env:"USER" envDefault:"user"`
	Password        string `env:"PASSWORD" envDefault:""`
	Name            string `env:"NAME" envDefault:"db"`
	SslMode         string `env:"SSL" envDefault:"disable"`
	MaxConnections  int    `env:"MAX_CONNECTIONS" envDefault:"20"`
	MinConnections  int    `env:"MIN_CONNECTIONS" envDefault:"1"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME" envDefault:"300"`
}

type HTTPServerConfig struct {
	Port            string `env:"PORT" envDefault:"8000"`
	ReadTimeout     int    `env:"READ_TIMEOUT" envDefault:"15"`
	WriteTimeout    int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	IdleTimeout     int    `env:"IDLE_TIMEOUT" envDefault:"60"`
	RateLimit       string `env:"RATE_LIMIT" envDefault:"100-S"`
	Environment     string `env:"ENVIRONMENT" envDefault:"development"`
	SwaggerUser     string `env:"SWAGGER_USER", envDefault:"admin"`
	SwaggerPassword string `env:"SWAGGER_PASSWORD", envDefault:"admin123"`
}

type JwtTokenConfig struct {
	SecretKey        string `env:"SECRET_KEY" envDefault:"3891aSDk23aSDa3j#@sd"`
	SecretRefreshKey string `env:"REFRESH_SECRET_KEY" envDefault:"h3i12iaSD32u98da@#%aisd"`
}

type FileStorageConfig struct {
	AccessKey string `env:"FILE_STORAGE_ACCESS_KEY" envDefault:"admin"`
	SecretKey string `env:"FILE_STORAGE_SECRET_KEY" envDefault:"admin"`
	Endpoint  string `env:"FILE_STORAGE_ENDPOINT" envDefault:"http://localhost:9000"`
	Region    string `env:"FILE_STORAGE_REGION" envDefault:"us-east-1"`
	Bucket    string `env:"FILE_STORAGE_BUCKET" envDefault:"donation-storage"`
	ACL       string `env:"FILE_STORAGE_ACL" envDefault:"public-read"`
	URL       string `env:"FILE_STORAGE_URL" envDefault:"http://localhost:9000/donation-storage"`
}

type SentryConfig struct {
	DSN              string  `env:"DSN"`
	TracesSampleRate float64 `env:"TRACES_SAMPLE_RATE" envDefault:"1.0"`
}

func New() (Interface, error) {
	_ = godotenv.Load()

	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *config) LogConfig() LogConfig {
	return c.Log
}

func (c *config) DatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *config) HTTPServerConfig() HTTPServerConfig {
	return c.HTTPServer
}

func (c *config) JwtTokenConfig() JwtTokenConfig {
	return c.JwtToken
}

func (c *config) FileStorageConfig() FileStorageConfig {
	return c.FileStorage
}

func (c *config) SentryConfig() SentryConfig {
	return c.Sentry
}
