package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"

	"github.com/felipeversiane/donation-server/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"

	_ "github.com/felipeversiane/donation-server/docs"
)

const (
	ErrPanicRecovered      = "panic recovered"
	MsgStartingHTTPServer  = "starting HTTP server"
	ErrServerFailedToStart = "server failed to start"
	MsgInitiatingShutdown  = "initiating graceful shutdown"
	ErrShutdownFailed      = "server shutdown failed"
	MsgShutdownSuccessful  = "server shutdown completed successfully"
	MsgHTTPRequest         = "HTTP request"
)

type server struct {
	router *gin.Engine
	srv    *http.Server
	config config.HTTPServerConfig
}

type ServerInterface interface {
	Start() error
	Shutdown(ctx context.Context) error
	InitRoutes()
}

func New(
	httpConfig config.HTTPServerConfig,
) ServerInterface {
	setupGinMode(httpConfig)
	router := setupRouter(httpConfig)

	server := &server{
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

func (s *server) InitRoutes() {
	v1 := s.router.Group("/api/v1")
	{
		// @Summary Health Check
		// @Description Returns the status of the server
		// @Tags Health
		// @Produce json
		// @Success 200 {object} map[string]interface{}
		// @Router /api/v1/health [get]
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status":    "up",
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			})
		})
		swagger := v1.Group("/swagger")
		if s.config.Environment != "development" {
			swagger.Use(swaggerAuthMiddleware(s.config.SwaggerUser, s.config.SwaggerPassword))
		}
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}
}

func (s *server) Start() error {
	slog.Info(MsgStartingHTTPServer, slog.String("port", s.config.Port))

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error(ErrServerFailedToStart, "error", err)
		return err
	}

	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	slog.Info(MsgInitiatingShutdown)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error(ErrShutdownFailed, "error", err)
		return err
	}

	slog.Info(MsgShutdownSuccessful)

	return nil
}

func setupGinMode(httpConfig config.HTTPServerConfig) {
	if httpConfig.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
		return
	}
	gin.SetMode(gin.DebugMode)
}

func setupRouter(httpConfig config.HTTPServerConfig) *gin.Engine {
	router := gin.New()

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		slog.Error(ErrPanicRecovered, "error", recovered)
		c.AbortWithStatus(500)
	}))

	router.Use(logMiddleware())
	router.Use(corsMiddleware())
	router.Use(securityMiddleware(httpConfig.Environment))

	if httpConfig.Environment != "development" {
		rate, _ := limiter.NewRateFromFormatted(httpConfig.RateLimit)
		store := memoryStore.NewStore()
		router.Use(ginLimiter.NewMiddleware(limiter.New(store, rate)))
	}

	return router
}
