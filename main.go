package main

import (
	"github.com/gin-gonic/gin"
)

type Recipe struct {
  Name string `json:"name"`
  Tags [] string `json:"tags"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
  r.Run()
}
