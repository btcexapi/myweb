package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user info v0.4",
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}