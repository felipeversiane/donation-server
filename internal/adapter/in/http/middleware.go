package http

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unrolled/secure"

	"github.com/felipeversiane/donation-server/pkg/logger"
)

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

func logMiddleware(log logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		log.WithContext(c.Request.Context()).Info(MsgHTTPRequest,
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", c.Writer.Status()),
			slog.String("client_ip", c.ClientIP()),
			slog.Duration("latency", latency),
		)
	}
}

func contextMiddleware(log logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), "request_id", reqID)

		if userID := c.GetHeader("X-User-ID"); userID != "" {
			ctx = context.WithValue(ctx, "user_id", userID)
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func securityMiddleware(env string) gin.HandlerFunc {
	opts := secure.Options{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		SSLRedirect:          env != "development",
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
		ReferrerPolicy:       "strict-origin-when-cross-origin",
		IsDevelopment:        env == "development",
	}
	secureMiddleware := secure.New(opts)
	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.Next()
	}
}

func swaggerAuthMiddleware(user, password string) gin.HandlerFunc {
	accounts := gin.Accounts{
		user: password,
	}
	return gin.BasicAuth(accounts)
}
