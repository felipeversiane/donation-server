package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"

	"github.com/felipeversiane/donation-server/internal/infrastructure/config"
	"github.com/felipeversiane/donation-server/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	router *gin.Engine
	srv    *http.Server
	config config.HttpServerConfig
	db     database.DatabaseInterface
}

type HttpServerInterface interface {
	Start() error
	Shutdown(ctx context.Context) error
	InitRoutes()
}

func New(config config.HttpServerConfig, db database.DatabaseInterface) HttpServerInterface {
	if config.Environment == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		slog.Error("panic recovered", "error", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.Use(logMiddleware())

	router.Use(corsMiddleware())

	if config.Environment != "development" {
		rate, _ := limiter.NewRateFromFormatted("100-S")
		store := memoryStore.NewStore()
		rateMiddleware := ginLimiter.NewMiddleware(limiter.New(store, rate))
		router.Use(rateMiddleware)
	}

	server := &httpServer{
		router: router,
		srv: &http.Server{
			Addr:         ":" + config.Port,
			Handler:      router,
			ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
		},
		config: config,
		db:     db,
	}

	server.InitRoutes()

	return server
}

func (s *httpServer) InitRoutes() {
	v1 := s.router.Group("/api/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status":    "up",
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			})
		})
	}
}

func (s *httpServer) Start() error {
	slog.Info("starting HTTP server", slog.String("port", s.config.Port))

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server failed to start", "error", err)
		return err
	}

	return nil
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	slog.Info("initiating graceful shutdown")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
		return err
	}

	slog.Info("server shutdown completed successfully")
	s.db.Close()

	return nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func logMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		slog.Info("HTTP request",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", c.Writer.Status()),
			slog.String("client_ip", c.ClientIP()),
			slog.Duration("latency", latency),
		)
	}
}
