package routes

import (
	"StudentManagement/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/getRank", controller.GetRank)
	r.GET("/getTop", controller.GetTop)
	return r
}
