package bench

import "testing"

type Header struct {
	Key   string
	Value string
}

// 方案1：两个平行切片
var headerKeys = []string{
	"Host",
	"User-Agent",
	"Accept",
	"Accept-Language",
	"Accept-Encoding",
	"Content-Type",
	"Content-Length",
	"Authorization",
	"Referer",
	"Origin",
	"Cookie",
	"Connection",
	"Upgrade-Insecure-Requests",
	"Cache-Control",
	"Pragma",
	"DNT",
	"Sec-Fetch-Dest",
	"Sec-Fetch-Mode",
	"Sec-Fetch-Site",
	"Te",
}
var headersMap = map[string]string{
	"Host":                      "api.example.com:443",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
	"Accept-Language":           "en-US,en;q=0.5",
	"Accept-Encoding":           "gzip, deflate, br",
	"Content-Type":              "application/json; charset=utf-8",
	"Content-Length":            "1024",
	"Authorization":             "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	"Referer":                   "https://example.com/previous-page",
	"Origin":                    "https://example.com",
	"Cookie":                    "sessionId=abc123; userId=789; authToken=xyz456",
	"Connection":                "keep-alive",
	"Upgrade-Insecure-Requests": "1",
	"Cache-Control":             "no-cache",
	"Pragma":                    "no-cache",
	"DNT":                       "1",
	"Sec-Fetch-Dest":            "document",
	"Sec-Fetch-Mode":            "navigate",
	"Sec-Fetch-Site":            "same-origin",
	"Te":                        "trailers",
}

// 基准测试示例
func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range headerKeys {
			if v == "Te" { // 查找最后一个，最坏情况
				_ = v
			}
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = headersMap["Te"]
	}
}
