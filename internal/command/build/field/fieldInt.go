package field

import (
	"strconv"

	"Sheeter/internal"
)

// Int 32位元整數
type Int struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *Int) TypeExcel() string {
	return "int"
}

// TypeCpp 取得c++欄位類型
func (this *Int) TypeCpp() string {
	return "int32_t"
}

// TypeCs 取得c#欄位類型
func (this *Int) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *Int) TypeGo() string {
	return "int32"
}

// Hide 是否隱藏
func (this *Int) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Int) PrimaryKey() bool {
	return false
}

// GetNote 取得註解名稱
func (this *Int) GetNote() string {
	return this.Note
}

// GetName 取得欄位名稱
func (this *Int) GetName() string {
	return this.Name
}

// GetField 取得欄位類型
func (this *Int) GetField() string {
	return this.Field
}

// FillToMetas 寫入到元資料列表
func (this *Int) FillToMetas(metas Metas, data string) error {
	value, err := strconv.ParseInt(data, internal.Decimal, 32)

	if err != nil {
		return err
	} // if

	metas[this.Name] = int32(value)
	return nil
}
