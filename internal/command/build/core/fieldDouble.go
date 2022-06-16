package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldDouble 64位元浮點數
type FieldDouble struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldDouble) TypeExcel() string {
	return "double"
}

// TypeCs 取得c#欄位類型
func (this *FieldDouble) TypeCs() string {
	return "double"
}

// TypeGo 取得go欄位類型
func (this *FieldDouble) TypeGo() string {
	return "float64"
}

// IsShow 是否顯示
func (this *FieldDouble) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldDouble) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldDouble) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToFloat(input)
}

// ToLuaValue 轉換為lua值
func (this *FieldDouble) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToFloat(input); err != nil {
		return "", err
	} // if

	return input, nil
}
