package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "ad6c420f-aa9d-4e33-8b49-c1f88e0b675a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://bitbucket.org/"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://pkg.go.dev/testing"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://translate.google.com/"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "2HWK3xxw")
	assert.Equal(t, shortLink_2, "S4LoJsD8")
	assert.Equal(t, shortLink_3, "69mXPARr")
}
