package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldPkey 主要索引
type FieldPkey struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldPkey) TypeExcel() string {
	return "pkey"
}

// TypeCs 取得c#欄位類型
func (this *FieldPkey) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *FieldPkey) TypeGo() string {
	return "int32"
}

// IsShow 是否顯示
func (this *FieldPkey) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldPkey) IsPkey() bool {
	return true
}

// ToJsonValue 轉換為json值
func (this *FieldPkey) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToInt(input)
}

// ToLuaValue 轉換為lua值
func (this *FieldPkey) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToInt(input); err != nil {
		return "", err
	} // if

	return input, nil
}
