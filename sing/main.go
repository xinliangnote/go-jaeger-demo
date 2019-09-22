package main

import (
	"github.com/gin-gonic/gin"
	"sing/app/config"
	"sing/app/route"
	"log"
	"os"
)

func main() {
	gin.SetMode(config.AppMode)

	engine := gin.New()

	// 设置路由
	route.SetupRouter(engine)

	log.Println("server listen port" + config.AppPort)

	// 启动服务
	if err := engine.Run(config.AppPort); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
