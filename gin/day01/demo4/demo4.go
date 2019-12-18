package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// 默认参数
		c.String(200, c.DefaultQuery("id", "12"))
	})
	r.Run(":8080")
}
