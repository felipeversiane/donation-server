package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"

	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/internal/adapter/out/database"
	"github.com/gin-gonic/gin"
)

var (
	ErrPanicRecovered      = "panic recovered"
	MsgStartingHTTPServer  = "starting HTTP server"
	ErrServerFailedToStart = "server failed to start"
	MsgInitiatingShutdown  = "initiating graceful shutdown"
	ErrShutdownFailed      = "server shutdown failed"
	MsgShutdownSuccessful  = "server shutdown completed successfully"
	MsgHTTPRequest         = "HTTP request"
)

type httpServer struct {
	router *gin.Engine
	srv    *http.Server
	config config.HttpServerConfig
	db     database.DatabaseInterface
	env    string
}

type HttpServerInterface interface {
	Start() error
	Shutdown(ctx context.Context) error
	InitRoutes()
}

func New(config config.HttpServerConfig, env string, db database.DatabaseInterface) HttpServerInterface {
	if env == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		slog.Error(ErrPanicRecovered, "error", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.Use(logMiddleware())
	router.Use(corsMiddleware())
	router.Use(securityMiddleware(env))

	if env != "development" {
		rate, _ := limiter.NewRateFromFormatted(config.RateLimit)
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
		env:    env,
	}

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
	slog.Info(MsgStartingHTTPServer, slog.String("port", s.config.Port))

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error(ErrServerFailedToStart, "error", err)
		return err
	}

	return nil
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	slog.Info(MsgInitiatingShutdown)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error(ErrShutdownFailed, "error", err)
		return err
	}

	slog.Info(MsgShutdownSuccessful)
	s.db.Close()

	return nil
}
