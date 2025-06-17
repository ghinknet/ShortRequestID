# ShortRequestID

#### ShortRequestID is a simple Request ID generator

#### This project is only active on GitHub (https://github.com/ghinknet/ShortRequestID)
#### Repositories on any other platform are mirrors
#### We will not process PRs or Issues outside of GitHub

#### ShortRequestID is suitable for a low QPS environment
#### And can make sure a low repeat, while keeping the Request ID easy to read

## Usage

`import "github.com/ghinknet/ShortRequestID"`

Example:

```go
package main

import (
	"fmt"
	"github.com/ghinknet/ShortRequestID"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.New()

	r.Use(
		ShortRequestID.GinMiddleware(
			ShortRequestID.GinWithCustomHeaderStrKey("X-Custom-Header"),
			ShortRequestID.GinWithCustomParamStrKey("CustomParam"),
		),
	)

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		requestID, ok := c.Get("CustomParam")
		if !ok {
			requestID = "unknown"
		}
		c.String(http.StatusOK, fmt.Sprintf("%s pong %d", requestID.(string), time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

### Decode

Run python script "decode.py" in this project.