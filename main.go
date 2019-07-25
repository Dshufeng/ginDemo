package main

import (
	"ginDemo/routers"
	"github.com/gin-gonic/gin"
	"ginDemo/http/middleware"
	"ginDemo/config"
)

//import "github.com/gin-gonic/gin"

func main() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	// log middleware
	engine.Use(middleware.LoggerToFile())

	routers.InitRouter(engine)
	engine.Run(config.PORT) // 监听并在 0.0.0.0:4242 上启动服务
}
