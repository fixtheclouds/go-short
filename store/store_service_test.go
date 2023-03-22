package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://market.yandex.ru/catalog--stroitelstvo-i-remont/54503/list?srnum=1266&was_redir=1&rt=9&rs=eJwBTACz_zJKUAF6BQj_muMHmgEx0YPQv9C70L7RgtC90LjRgtC10LvRjCDRgdCw0LzQvtC60LvQtdGO0YnQuNC50YHRj6oBCgi2uJ8QENrMjQqjEy0R&suggest=1&suggest_type=search&text=%D1%83%D0%BF%D0%BB%D0%BE%D1%82%D0%BD%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%20%D1%81%D0%B0%D0%BC%D0%BE%D0%BA%D0%BB%D0%B5%D1%8E%D1%89%D0%B8%D0%B9%D1%81%D1%8F&hid=91597&allowCollapsing=1&local-offers-first=0"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(initialLink, shortURL, userUUId)

	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
