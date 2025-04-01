package main

import (
	"StudentManagement/config"
	"StudentManagement/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {

	config.LoadConfig()
	fmt.Println(config.AppConfig)

	r := routes.SetupRouter()

	// 设置模板目录
	r.LoadHTMLGlob("templates/*")

	// 处理静态资源
	r.Static("/static", "./static")

	// 访问首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "首页",
			"username": "北船",
		})
	})

	// 启动服务，监听 12345 端口
	r.Run(":" + strconv.Itoa(config.AppConfig.Server.Port))
}
