package go_i18n

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func TestLocalizerMiddleware(t *testing.T) {
	// Chuẩn bị dữ liệu test
	bundle := i18n.NewBundle(language.English)
	// Đăng ký hàm giải mã YAML
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	bundle.MustParseMessageFileBytes([]byte(`
TEST_KEY: "Test message in English"
`), "en.yaml")
	bundle.MustParseMessageFileBytes([]byte(`
TEST_KEY: "Thông điệp kiểm tra bằng tiếng Việt"
`), "vi.yaml")

	// Tạo router Gin cho test
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(LocalizerMiddleware(bundle))

	// Tạo handler test
	var localizerFromContext *i18n.Localizer
	var matchedLanguage string

	r.GET("/test", func(c *gin.Context) {
		localizerFromContext = MustGetLocalizer(c)
		msg := MustSafeLocalize(localizerFromContext, "TEST_KEY")
		matchedLanguage = msg
		c.String(http.StatusOK, msg)
	})

	// Test case 1: Accept-Language: vi
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Language", "vi")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Thông điệp kiểm tra bằng tiếng Việt", w.Body.String())
	assert.Equal(t, "Thông điệp kiểm tra bằng tiếng Việt", matchedLanguage)

	// Test case 2: Accept-Language: en
	req = httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Language", "en")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test message in English", w.Body.String())
	assert.Equal(t, "Test message in English", matchedLanguage)

	// Test case 3: Accept-Language không hợp lệ (fallback về English)
	req = httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Language", "invalid-language")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test message in English", w.Body.String())
	assert.Equal(t, "Test message in English", matchedLanguage)

	// Test case 4: Không có Accept-Language (fallback về English)
	req = httptest.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test message in English", w.Body.String())
	assert.Equal(t, "Test message in English", matchedLanguage)
}
