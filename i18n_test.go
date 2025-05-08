package go_i18n

import (
	"os"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func TestInitBundle(t *testing.T) {
	// Tạo thư mục tạm thời cho test
	testDir := "test_locales"
	err := os.Mkdir(testDir, 0755)
	assert.NoError(t, err)
	defer os.RemoveAll(testDir)

	// Tạo file ngôn ngữ test
	enContent := []byte(`
TEST_KEY: "Test message in English"
`)
	viContent := []byte(`
TEST_KEY: "Thông điệp kiểm tra bằng tiếng Việt"
`)

	err = os.WriteFile(testDir+"/en.yaml", enContent, 0644)
	assert.NoError(t, err)
	err = os.WriteFile(testDir+"/vi.yaml", viContent, 0644)
	assert.NoError(t, err)

	// Đặt biến môi trường cho test
	os.Setenv("LOCALES_PATH", testDir)
	defer os.Unsetenv("LOCALES_PATH")

	// Chạy hàm cần test
	InitBundle()

	// Kiểm tra kết quả
	assert.NotNil(t, Bundle)

	// Kiểm tra các ngôn ngữ đã được tải
	tags := Bundle.LanguageTags()
	assert.Contains(t, tags, language.English)
	assert.Contains(t, tags, language.Vietnamese)
}

func TestMustSafeLocalize(t *testing.T) {
	// Chuẩn bị dữ liệu test
	bundle := i18n.NewBundle(language.English)
	// Đăng ký hàm giải mã YAML
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	bundle.MustParseMessageFileBytes([]byte(`
TEST_KEY: "Test message"
`), "en.yaml")

	localizer := i18n.NewLocalizer(bundle, "en")

	// Test case 1: Khóa tồn tại
	result := MustSafeLocalize(localizer, "TEST_KEY")
	assert.Equal(t, "Test message", result)

	// Test case 2: Khóa không tồn tại
	result = MustSafeLocalize(localizer, "NONEXISTENT_KEY")
	assert.Equal(t, "NONEXISTENT_KEY", result)
}
