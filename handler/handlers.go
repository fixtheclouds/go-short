package handler

import (
	"net/http"

	"github.com/fixtheclouds/go-short/shortener"
	"github.com/fixtheclouds/go-short/store"
	"github.com/gin-gonic/gin"
)

type UrlInputRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	UserId      string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var request UrlInputRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(request.OriginalUrl, request.UserId)
	store.SaveUrlMapping(shortUrl, request.OriginalUrl, request.UserId)

	// TODO: configurable port
	host := "http://localhost:9808/"
	ctx.JSON(200, gin.H{
		"message":   "short URL created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}
