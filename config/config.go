package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	LogConfig         Log         `envPrefix:"LOG_"`
	DatabaseConfig    Database    `envPrefix:"DB_"`
	HTTPServerConfig  HTTPServer  `envPrefix:"HTTP_SERVER_"`
	JwtTokenConfig    JwtToken    `envPrefix:"JWT_"`
	FileStorageConfig FileStorage `envPrefix:"FILE_STORAGE_"`
}

type Interface interface {
	Database() Database
	HTTPServer() HTTPServer
	JwtToken() JwtToken
	FileStorage() FileStorage
	Log() Log
}

type Log struct {
	Level      string `env:"LEVEL" envDefault:"info"`
	Path       string `env:"PATH" envDefault:"logs/app.log"`
	Stdout     bool   `env:"STDOUT" envDefault:"true"`
	MaxSize    int    `env:"MAX_SIZE" envDefault:"10"`
	MaxBackups int    `env:"MAX_BACKUPS" envDefault:"5"`
	MaxAge     int    `env:"MAX_AGE" envDefault:"28"`
	Compress   bool   `env:"COMPRESS" envDefault:"true"`
}

type Database struct {
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

type HTTPServer struct {
	Port            string `env:"PORT" envDefault:"8000"`
	ReadTimeout     int    `env:"READ_TIMEOUT" envDefault:"15"`
	WriteTimeout    int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	IdleTimeout     int    `env:"IDLE_TIMEOUT" envDefault:"60"`
	RateLimit       string `env:"RATE_LIMIT" envDefault:"100-S"`
	Environment     string `env:"ENVIRONMENT" envDefault:"development"`
	SwaggerUser     string `env:"SWAGGER_USER" envDefault:"admin"`
	SwaggerPassword string `env:"SWAGGER_PASSWORD" envDefault:"admin123"`
}

type JwtToken struct {
	SecretKey        string `env:"SECRET_KEY" envDefault:"3891aSDk23aSDa3j#@sd"`
	SecretRefreshKey string `env:"REFRESH_SECRET_KEY" envDefault:"h3i12iaSD32u98da@#%aisd"`
}

type FileStorage struct {
	AccessKey string `env:"FILE_STORAGE_ACCESS_KEY" envDefault:"admin"`
	SecretKey string `env:"FILE_STORAGE_SECRET_KEY" envDefault:"admin"`
	Endpoint  string `env:"FILE_STORAGE_ENDPOINT" envDefault:"http://localhost:9000"`
	Region    string `env:"FILE_STORAGE_REGION" envDefault:"us-east-1"`
	Bucket    string `env:"FILE_STORAGE_BUCKET" envDefault:"donation-storage"`
	ACL       string `env:"FILE_STORAGE_ACL" envDefault:"public-read"`
	URL       string `env:"FILE_STORAGE_URL" envDefault:"http://localhost:9000/donation-storage"`
}

func New() (Interface, error) {
	_ = godotenv.Load()

	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *config) Log() Log {
	return c.LogConfig
}

func (c *config) Database() Database {
	return c.DatabaseConfig
}

func (c *config) HTTPServer() HTTPServer {
	return c.HTTPServerConfig
}

func (c *config) JwtToken() JwtToken {
	return c.JwtTokenConfig
}

func (c *config) FileStorage() FileStorage {
	return c.FileStorageConfig
}
