package go_i18n

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var Bundle *i18n.Bundle

func InitBundle() {
	Bundle = i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	localesPath := os.Getenv("LOCALES_PATH")
	if localesPath == "" {
		localesPath = "locales" // Đường dẫn mặc định
	}
	files, err := filepath.Glob(filepath.Join(localesPath, "*.yaml"))
	if err != nil {
		log.Fatalf("failed to read locale files: %v", err)
	}

	log.Printf("Found %d locale files: %v", len(files), files)

	for _, file := range files {
		if _, err := Bundle.LoadMessageFile(file); err != nil {
			log.Printf("failed to load message file %s: %v", file, err)
		} else {
			log.Printf("Successfully loaded message file: %s", file)
		}
	}

	// In ra danh sách các ngôn ngữ đã tải
	log.Printf("Available language tags: %v", Bundle.LanguageTags())
}

func MustSafeLocalize(localizer *i18n.Localizer, messageID string) string {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
	if err != nil {
		log.Printf("message not found: %s, fallback to messageID", messageID)
		return messageID
	}
	return msg
}
