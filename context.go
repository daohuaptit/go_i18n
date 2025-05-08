package go_i18n

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func MustGetLocalizer(c *gin.Context) *i18n.Localizer {
	localizer, ok := c.Get("localizer")
	if !ok {
		log.Println("localizer not found in context, fallback to English")
		return i18n.NewLocalizer(Bundle, language.English.String())
	}
	return localizer.(*i18n.Localizer)
}

