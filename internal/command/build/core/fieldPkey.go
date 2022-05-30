package core

import "strconv"

// FieldPkey 主要索引
type FieldPkey struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldPkey) TypeExcel() string {
	return "pkey"
}

// TypeCpp 取得c++欄位類型
func (this *FieldPkey) TypeCpp() string {
	return CppNamespace + "::pkey" // pkey型態宣告在命名空間中
}

// TypeCs 取得c#欄位類型
func (this *FieldPkey) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *FieldPkey) TypeGo() string {
	return "int"
}

// Hide 是否隱藏
func (this *FieldPkey) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldPkey) PrimaryKey() bool {
	return true
}

// Transform 字串轉換
func (this *FieldPkey) Transform(input string) (result interface{}, err error) {
	return strconv.Atoi(input)
}
