package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/felipeversiane/donation-server/config"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"
)

const (
	ErrPanicRecovered      = "panic recovered"
	MsgStartingHTTPServer  = "starting HTTP server"
	ErrServerFailedToStart = "server failed to start"
	MsgInitiatingShutdown  = "initiating graceful shutdown"
	ErrShutdownFailed      = "server shutdown failed"
	MsgShutdownSuccessful  = "server shutdown completed successfully"
	ErrSentryInit          = "error initializing sentry"
	MsgHTTPRequest         = "HTTP request"
)

type httpServer struct {
	router *gin.Engine
	srv    *http.Server
	config config.HttpServerConfig
}

type HttpServerInterface interface {
	Start() error
	Shutdown(ctx context.Context) error
	InitRoutes()
}

func New(
	httpConfig config.HttpServerConfig,
	sentryConfig config.SentryConfig,
) HttpServerInterface {
	setupGinMode(httpConfig)
	setupSentry(sentryConfig, httpConfig)
	router := setupRouter(httpConfig)

	server := &httpServer{
		router: router,
		srv: &http.Server{
			Addr:         ":" + httpConfig.Port,
			Handler:      router,
			ReadTimeout:  time.Duration(httpConfig.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(httpConfig.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(httpConfig.IdleTimeout) * time.Second,
		},
		config: httpConfig,
	}

	return server
}

func (s *httpServer) InitRoutes() {
	v1 := s.router.Group("/api/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
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
	sentry.Flush(2 * time.Second)

	return nil
}

func setupGinMode(httpConfig config.HttpServerConfig) {
	if httpConfig.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
		return
	}
	gin.SetMode(gin.DebugMode)
}

func setupSentry(sentryConfig config.SentryConfig, httpConfig config.HttpServerConfig) {
	if httpConfig.Environment == "development" {
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryConfig.DSN,
		EnableTracing:    true,
		TracesSampleRate: sentryConfig.TracesSampleRate,
	}); err != nil {
		slog.Error(ErrSentryInit, "error", err)
	}
}

func setupRouter(httpConfig config.HttpServerConfig) *gin.Engine {
	router := gin.New()

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		slog.Error(ErrPanicRecovered, "error", recovered)
		c.AbortWithStatus(500)
	}))

	router.Use(logMiddleware())
	router.Use(corsMiddleware())
	router.Use(securityMiddleware(httpConfig.Environment))
	router.Use(sentrygin.New(sentrygin.Options{}))

	if httpConfig.Environment != "development" {
		rate, _ := limiter.NewRateFromFormatted(httpConfig.RateLimit)
		store := memoryStore.NewStore()
		router.Use(ginLimiter.NewMiddleware(limiter.New(store, rate)))
	}

	return router
}
