package main

import (
	"dejian/dj_task/config"
	"dejian/dj_task/log"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	configPath := flag.String("config", "", "配置信息路径")
	port := flag.Int("port", -1, "服务监听端口")
	flag.Parse()
	config.InitConfig(*configPath, *port)
	log.Init()

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.Run()
}
