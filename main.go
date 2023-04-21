package main

import (
	"fmt"
	"os"

	"github.com/fixtheclouds/go-short/handler"
	"github.com/fixtheclouds/go-short/store"
	"github.com/gin-gonic/gin"
)

func main() {
	var port string = "9800"
	request := gin.Default()
	request.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Go short !",
		})
	})

	request.POST("/", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	request.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	store.InitializeStore()

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	err := request.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
