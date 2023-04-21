package main

import (
	"fmt"

	"github.com/fixtheclouds/go-short/handler"
	"github.com/fixtheclouds/go-short/store"
	"github.com/gin-gonic/gin"
)

func main() {
	request := gin.Default()
	request.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go short !",
		})
	})

	request.POST("/", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	request.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := request.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
