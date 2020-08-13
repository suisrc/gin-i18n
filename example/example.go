package main

import (
	"github.com/gin-gonic/gin"
	i18n "github.com/suisrc/gin-i18n"
	"golang.org/x/text/language"
)

func setupRouter(bundle *i18n.Bundle) *gin.Engine {
	r := gin.Default()

	//r.Use(i18n.Serve(bundle))
	r.GET("/ping", func(c *gin.Context) {

		text := i18n.FormatMessage(c,
			&i18n.Message{
				ID:    "ping-text",
				Other: "你好,{{.who}}",
			},
			map[string]interface{}{
				"who": "gin",
			})

		c.String(200, text)
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.String(200, i18n.FormatMessage(c, &i18n.Message{ID: "ping2-text", Other: "我是{{.who}}"}, i18n.Data{"who": "gin"}))
	})
	r.GET("/ping3", func(c *gin.Context) {
		c.String(200, i18n.FormatText(c, &i18n.Message{ID: "ping3-text", Other: "测试"}))
	})

	return r
}

func main() {

	bundle := i18n.NewBundle(
		language.Chinese,
		"example/active.zh-CN.toml",
		"example/active.en-US.toml",
		"example/active.ja-JP.toml",
	)

	r := setupRouter(bundle)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
