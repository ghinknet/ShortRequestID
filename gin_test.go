package ShortRequestID

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testParamID = "TestParam"
const testXRequestID = "X-TestRequest-ID"

func TestGinWithCustomHeaderStrKey(t *testing.T) {
	cfg := &ginConfig{
		headerKey: "X-Request-ID",
		paramKey:  "RequestID",
	}

	GinWithCustomHeaderStrKey("X-Test-Header")(cfg)

	if cfg.headerKey != "X-Test-Header" {
		t.Errorf("X-Test-Header not set")
	}
}

func TestGinWithCustomParamStrKey(t *testing.T) {
	cfg := &ginConfig{
		headerKey: "X-Request-ID",
		paramKey:  "RequestID",
	}

	GinWithCustomParamStrKey("TestParam")(cfg)

	if cfg.paramKey != "TestParam" {
		t.Errorf("TestParam not set")
	}
}

func emptySuccessResponse(c *gin.Context) {
	requestID, _ := c.Get(testParamID)
	c.String(http.StatusOK, requestID.(string))
}

func TestGinMiddleware(t *testing.T) {
	r := gin.New()
	r.Use(GinMiddleware(
		GinWithCustomParamStrKey(testParamID),
		GinWithCustomHeaderStrKey(testXRequestID),
	))
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Header().Get(testXRequestID))
	assert.Equal(t, w.Header().Get(testXRequestID), w.Body.String())
}
