package util

import "io"

// SilentClose 關閉物件, 不會處理錯誤
func SilentClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	} // if
}
