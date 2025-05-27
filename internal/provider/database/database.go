package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/pkg/logger"
	"github.com/felipeversiane/donation-server/pkg/number"
)

type database struct {
	pool   *pgxpool.Pool
	config config.DatabaseConfig
	logger logger.Interface
}

type Interface interface {
	Pool() *pgxpool.Pool
	Ping(ctx context.Context) error
	Close()
}

func New(config config.DatabaseConfig, logger logger.Interface) (Interface, error) {
	logger.Logger().Info("initializing database connection...")
	dsn := getConnectionString(config)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Logger().Error("error parsing pool config", "error", err)
		return nil, fmt.Errorf("error parsing pool config: %w", err)
	}

	maxConns, err := number.SafeIntToInt32(config.MaxConnections)
	if err != nil {
		logger.Logger().Error("invalid max connections", "error", err)
		return nil, fmt.Errorf("invalid max connections: %w", err)
	}
	poolConfig.MaxConns = maxConns

	minConns, err := number.SafeIntToInt32(config.MinConnections)
	if err != nil {
		logger.Logger().Error("invalid min connections", "error", err)
		return nil, fmt.Errorf("invalid min connections: %w", err)
	}
	poolConfig.MinConns = minConns

	poolConfig.MaxConnLifetime = time.Duration(config.ConnMaxLifetime) * time.Second

	logger.Logger().Info("creating database connection pool...")

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		logger.Logger().Error("error creating connection pool", "error", err)
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}

	db := &database{
		pool:   pool,
		config: config,
		logger: logger,
	}

	logger.Logger().Info("attempting to ping database")

	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		logger.Logger().Error("error pinging database", "error", err)
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	logger.Logger().Info("database connection established successfully")

	return db, nil
}

func (d *database) Ping(ctx context.Context) error {
	err := d.pool.Ping(ctx)
	if err != nil {
		d.logger.WithContext(ctx).Warn("database ping failed", "error", err)
	}
	return err
}

func (d *database) Close() {
	if d.pool != nil {
		d.pool.Close()
		d.logger.Logger().Info("database connection closed")
	}
}

func (d *database) Pool() *pgxpool.Pool {
	return d.pool
}

func getConnectionString(config config.DatabaseConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=%s",
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.Host,
		config.SslMode,
	)
}
