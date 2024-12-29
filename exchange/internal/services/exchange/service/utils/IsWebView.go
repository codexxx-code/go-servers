package utils

import (
	"github.com/mileusna/useragent"
)

// IsWebView isWebView проверяет является ли userAgent типа WebView
func IsWebView(userAgent string) bool {

	ua := useragent.Parse(userAgent)

	// Пока считаем, что если IOS и браузер Safari,
	// то тут не определить что это браузер на пк
	return ua.IsSafari() && ua.IsIOS()
}
