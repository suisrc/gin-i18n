package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	i18n "github.com/suisrc/gin-i18n"
	"golang.org/x/text/language"
)

func TestPingRoute(t *testing.T) {
	bundle := i18n.NewBundle(
		language.Chinese,
		"active.zh-CN.toml",
		"active.en-US.toml",
		"active.ja-JP.toml",
	)

	router := setupRouter(bundle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "你好,gin", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ping?lang=en-US", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello,gin", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ping?lang=ja-JP", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "こんにちは,gin", w.Body.String())
}
