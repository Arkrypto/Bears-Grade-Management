package routes

import (
	"StudentManagement/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		var m = service.GetSortedStudents(0, 0)
		fmt.Println(m)

		var t = service.GetTopStudents()
		fmt.Println(t)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
