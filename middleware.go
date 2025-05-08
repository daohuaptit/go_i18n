package go_i18n

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LocalizerMiddleware(bundle *i18n.Bundle) gin.HandlerFunc {
	matcher := language.NewMatcher(bundle.LanguageTags())

	return func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		log.Printf("Request Accept-Language: %s", acceptLang)

		tags, _, err := language.ParseAcceptLanguage(acceptLang)
		if err != nil || len(tags) == 0 {
			log.Printf("Using default language (English) due to: %v", err)
			tags = []language.Tag{language.English}
		}

		matchedTag, _, _ := matcher.Match(tags...)
		log.Printf("Matched language tag: %s", matchedTag.String())

		localizer := i18n.NewLocalizer(bundle, matchedTag.String())
		c.Set("localizer", localizer)
		c.Next()
	}
}
