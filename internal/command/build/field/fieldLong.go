package field

import (
	"strconv"

	"Sheeter/internal"
)

// Long 64位元整數
type Long struct {
}

// TypeExcel 取得excel欄位類型
func (this *Long) TypeExcel() string {
	return "long"
}

// TypeCpp 取得c++欄位類型
func (this *Long) TypeCpp() string {
	return "int64_t"
}

// TypeCs 取得c#欄位類型
func (this *Long) TypeCs() string {
	return "long"
}

// TypeGo 取得go欄位類型
func (this *Long) TypeGo() string {
	return "int64"
}

// Hide 是否隱藏
func (this *Long) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Long) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *Long) Transform(input string) (result interface{}, err error) {
	return strconv.ParseInt(input, internal.Decimal, 64)
}
