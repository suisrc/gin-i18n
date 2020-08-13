package i18n

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// 定义上下文中的键
const (
	GinI18nKey = "suisrc/gin-i18n"
)

type (
	// Bundle i18n.Bundle
	Bundle = i18n.Bundle

	// Message i18n.Message
	Message = i18n.Message

	// LocalizeConfig i18n.LocalizeConfig
	LocalizeConfig = i18n.LocalizeConfig

	// Data TemplateData
	Data = map[string]interface{}
)

// NewBundle new bundle
func NewBundle(tag language.Tag, tomls ...string) *Bundle {
	bundle := i18n.NewBundle(tag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, file := range tomls {
		bundle.LoadMessageFile(file)
	}
	return bundle
}

// Serve 服务
func Serve(bundle *Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.FormValue("lang")
		accept := c.Request.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)
		SetLocalizer(c, localizer)

		c.Next()

		// 基于i18n句柄可能会跨域生命周期的存在,所以不主动清空
		// helper.SetLocalizer(c, nil) // 清除
	}
}

// MustFormat must
func MustFormat(c *gin.Context, lc *i18n.LocalizeConfig) string {
	return MustLocalizer(c).MustLocalize(lc)
}

// FormatText ft
func FormatText(c *gin.Context, message *Message) string {
	return FormatMessage(c, message, nil)
}

// FormatMessage fm
func FormatMessage(c *gin.Context, message *Message, args map[string]interface{}) string {
	if localizer, ok := GetLocalizer(c); ok {
		return localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: message,
			TemplateData:   args,
		})
	}
	// 加载i18n中间件后,不会进入该分支,简单处理未加载i18n中间件时候的处理内容
	// After loading the i18n middleware, it will not enter this branch and simply handle the unloaded
	return formatInternalMessage(message, args)
}

// MustLocalizer i18n
func MustLocalizer(c *gin.Context) *i18n.Localizer {
	localizer, ok := GetLocalizer(c)
	if !ok {
		panic(errors.New("context no has i18n localizer"))
	}
	return localizer
}

// GetLocalizer i18n
func GetLocalizer(c *gin.Context) (*i18n.Localizer, bool) {
	if v, ok := c.Get(GinI18nKey); ok {
		if l, b := v.(*i18n.Localizer); b {
			return l, true
		}
	}
	return nil, false
}

// SetLocalizer i18n
func SetLocalizer(c *gin.Context, l *i18n.Localizer) {
	c.Set(GinI18nKey, l)
}

func formatInternalMessage(message *i18n.Message, args map[string]interface{}) string {
	if args == nil {
		return message.Other
	}
	tpl := i18n.NewMessageTemplate(message)
	msg, err := tpl.Execute("other", args, nil)
	if err != nil {
		panic(err)
	}
	return msg
}
