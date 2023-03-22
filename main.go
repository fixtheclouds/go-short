package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func main() {
	request := gin.Default()
	request.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go short !",
		})
	})

	err := request.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
