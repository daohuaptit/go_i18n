# go-i18n: Thư viện đa ngôn ngữ cho Go

[![Go Reference](https://pkg.go.dev/badge/github.com/daohuaptit/go_i18n.svg)](https://pkg.go.dev/github.com/daohuaptit/go_i18n)

`go-i18n` là một thư viện đơn giản và mạnh mẽ để hỗ trợ đa ngôn ngữ (i18n) trong ứng dụng Go, đặc biệt là với Gin framework.

## Tính năng

- Tích hợp dễ dàng với Gin framework
- Tự động phát hiện ngôn ngữ từ header `Accept-Language`
- Hỗ trợ tệp ngôn ngữ định dạng YAML
- Fallback an toàn khi không tìm thấy bản dịch
- API đơn giản và dễ sử dụng

## Cài đặt 

```bash
go get github.com/daohuaptit/go_i18n
```


## Cấu trúc thư mục
```
your-app/
├── locales/
│ ├── en.yaml
│ ├── vi.yaml
│ └── ...
└── main.go
```

## Sử dụng
## Tệp ngôn ngữ

Tệp ngôn ngữ được định dạng YAML với cấu trúc đơn giản:

**locales/en.yaml**:

```yaml
GREETING: "Hello, {{.Name}}!"
ERROR_MESSAGE: "An error occurred"
USER_CREATED: "User created successfully"
INVALID_EMAIL: "Invalid email address"
```

**locales/vi.yaml**:

```yaml
GREETING: "Xin chào, {{.Name}}!"
ERROR_MESSAGE: "Đã xảy ra lỗi"
USER_CREATED: "Tạo người dùng thành công"
INVALID_EMAIL: "Email không hợp lệ"
```


## Sử dụng cơ bản

```go
    package main
    import (
    "net/http"
    "github.com/daohuaptit/go_i18n"
    "github.com/gin-gonic/gin"
    )
    func main() {
    // Khởi tạo bundle i18n
    go_i18n.InitBundle()
    // Tạo router Gin
    r := gin.Default()
    // Thêm middleware i18n
    r.Use(go_i18n.LocalizerMiddleware(go_i18n.Bundle))
    // Định nghĩa route
    r.GET("/greeting", func(c gin.Context) {
    // Lấy localizer từ context
    localizer := go_i18n.MustGetLocalizer(c)
    // Dịch thông điệp
    msg := go_i18n.MustSafeLocalize(localizer, "GREETING")
    c.JSON(http.StatusOK, gin.H{
    "message": msg,
    })
    })
    r.Run(":8080")
    }
```

## Sử dụng nâng cao

```go
    r.GET("/hello", func(c gin.Context) {
    localizer := go_i18n.MustGetLocalizer(c)
    // Dịch với tham số
    msg, err := localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "GREETING",
    TemplateData: map[string]interface{}{
    "Name": "John Doe",
    },
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
        c.JSON(http.StatusOK, gin.H{"message": msg})
    })
```

## Các phương thức API

### InitBundle()

Khởi tạo bundle i18n và tải tất cả tệp ngôn ngữ từ thư mục `locales`.

### LocalizerMiddleware(bundle *i18n.Bundle) gin.HandlerFunc

Middleware Gin để phát hiện ngôn ngữ từ header `Accept-Language` và lưu localizer vào context.

### MustGetLocalizer(c *gin.Context) *i18n.Localizer

Lấy localizer từ context Gin. Nếu không tìm thấy, trả về localizer mặc định (tiếng Anh).

### MustSafeLocalize(localizer *i18n.Localizer, messageID string) string

Dịch một thông điệp theo ID. Nếu không tìm thấy, trả về messageID.

## Ví dụ đầy đủ

```go
    package main
    import (
    "net/http"
    "github.com/daohuaptit/go_i18n"
    "github.com/gin-gonic/gin"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    )
    func main() {
        go_i18n.InitBundle()
        r := gin.Default()
        r.Use(go_i18n.LocalizerMiddleware(go_i18n.Bundle))
        r.GET("/user", func(c gin.Context) {
            localizer := go_i18n.MustGetLocalizer(c)
            msg := go_i18n.MustSafeLocalize(localizer, "USER_CREATED")
            c.JSON(http.StatusOK, gin.H{
            "code": "USER_CREATED",
            "message": msg,
            })
            })
            r.GET("/error", func(c gin.Context) {
            localizer := go_i18n.MustGetLocalizer(c)
            msg := go_i18n.MustSafeLocalize(localizer, "INVALID_EMAIL")
            c.JSON(http.StatusBadRequest, gin.H{
            "code": "INVALID_EMAIL",
            "message": msg,
            })
        })
        r.GET("/greeting", func(c gin.Context) {
            localizer := go_i18n.MustGetLocalizer(c)
            name := c.Query("name")
            if name == "" {
            name = "Guest"
            }
            msg, err := localizer.Localize(&i18n.LocalizeConfig{
            MessageID: "GREETING",
            TemplateData: map[string]interface{}{
            "Name": name,
            },
            })
            if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
            }
            c.JSON(http.StatusOK, gin.H{"message": msg})
        })
        r.Run(":8080")
    }
```

## Giấy phép

MIT