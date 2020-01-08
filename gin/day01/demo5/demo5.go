package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// r = gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})
	r.Run()
}
