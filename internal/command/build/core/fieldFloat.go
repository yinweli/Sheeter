package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldFloat 32位元浮點數
type FieldFloat struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldFloat) TypeExcel() string {
	return "float"
}

// TypeCpp 取得c++欄位類型
func (this *FieldFloat) TypeCpp() string {
	return "float"
}

// TypeCs 取得c#欄位類型
func (this *FieldFloat) TypeCs() string {
	return "float"
}

// TypeGo 取得go欄位類型
func (this *FieldFloat) TypeGo() string {
	return "float32"
}

// IsShow 是否顯示
func (this *FieldFloat) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldFloat) IsPkey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldFloat) Transform(input string) (result interface{}, err error) {
	return util.StrToFloat(input)
}
