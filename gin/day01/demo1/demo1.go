package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog": "https://mowuya.cn",
			"weat": "123456",
		})
	})
	r.Run(":8080")
}
