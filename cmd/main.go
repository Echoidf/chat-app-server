package main

import (
	"chat-app/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.GetConfig().Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	_ = r.Run(fmt.Sprintf("127.0.0.1:%d", config.GetConfig().HttpPort))
}
