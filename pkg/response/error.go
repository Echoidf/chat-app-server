package response

import (
	"chat-app/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, status int, err error, message ...string) {
	cfg := config.GetConfig()
	errorRes := map[string]interface{}{}

	if cfg.Environment != config.ProductionEnv {
		errorRes["debug"] = err.Error()
	}

	if len(message) > 0 {
		for i, msg := range message {
			errorRes[fmt.Sprintf("message%d", i+1)] = msg
		}
	}
	c.JSON(status, Response{Error: errorRes})
}
