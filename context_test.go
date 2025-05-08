package go_i18n

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func TestMustGetLocalizer(t *testing.T) {
	// Chuẩn bị dữ liệu test
	Bundle = i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// Tạo context Gin
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Test case 1: Localizer không có trong context
	localizer := MustGetLocalizer(c)
	assert.NotNil(t, localizer)
	// Kiểm tra localizer mặc định là tiếng Anh
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: "test", Other: "test"},
	})
	assert.NoError(t, err)
	assert.Equal(t, "test", msg)

	// Test case 2: Localizer có trong context
	expectedLocalizer := i18n.NewLocalizer(Bundle, "vi")
	c.Set("localizer", expectedLocalizer)

	localizer = MustGetLocalizer(c)
	assert.Equal(t, expectedLocalizer, localizer)
}
