package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		log.Println("输出到控制台的日志")
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}