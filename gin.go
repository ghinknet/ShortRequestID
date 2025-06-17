package ShortRequestID

import (
	"github.com/gin-gonic/gin"
)

// ginConfig is the config options for Gin middleware
type ginConfig struct {
	headerKey string
	paramKey  string
}

// headerXRequestID is the global var to store header str key
var headerXRequestID string

// Option for queue system
type Option func(*ginConfig)

// GinMiddleware is a middleware for Gin framework
func GinMiddleware(opts ...Option) gin.HandlerFunc {
	cfg := &ginConfig{
		headerKey: "X-Request-ID",
		paramKey:  "RequestID",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	headerXRequestID = cfg.headerKey

	return func(c *gin.Context) {
		headers := c.Writer.Header()

		requestID := GenerateRequestID()

		headers.Add(cfg.headerKey, requestID)
		c.Set(cfg.paramKey, requestID)

		c.Next()
	}
}

// GinWithCustomHeaderStrKey set custom header key for request id
func GinWithCustomHeaderStrKey(s string) Option {
	return func(cfg *ginConfig) {
		cfg.headerKey = s
	}
}

// GinWithCustomParamStrKey set custom param key for request id
func GinWithCustomParamStrKey(s string) Option {
	return func(cfg *ginConfig) {
		cfg.paramKey = s
	}
}
