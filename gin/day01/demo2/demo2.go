package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.POST("/users", func(c *gin.Context) {
		// 创建一个用户

	})
	r.DELETE("/users/123", func(c *gin.Context) {
		// 删除用户
	})
	r.PUT("/users/123", func(c *gin.Context) {
		// 更新用户123
	})
	r.PATCH("/users/123", func(c *gin.Context) {
		// 更新用户123的部分信息
	})
	r.Run(":8080")
}
