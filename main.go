package main

import (
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initRouter(r)
	_ = r.Run(config.Port)
}
