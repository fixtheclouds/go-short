package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fixtheclouds/go-short/shortener"
	"github.com/fixtheclouds/go-short/store"
	"github.com/gin-gonic/gin"
)

type UrlInputRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	UserId      string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var port string = "9800"
	var host string = "localhost"
	var request UrlInputRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(request.OriginalUrl, request.UserId)
	store.SaveUrlMapping(request.OriginalUrl, shortUrl, request.UserId)

	if os.Getenv("PORT") != "" {
		port = os.Getenv(("PORT"))
	}

	if os.Getenv("BASE_URL") != "" {
		host = os.Getenv("BASE_URL")
	}
	base := fmt.Sprintf("http://%s:%s/", host, port)
	ctx.JSON(200, gin.H{
		"message":   "short URL created successfully",
		"short_url": base + shortUrl,
	})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}
