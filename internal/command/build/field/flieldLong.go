package field

import (
	"strconv"

	"Sheeter/internal"
)

// Long 64位元整數
type Long struct {
	Data
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

// GetNote 取得註解名稱
func (this *Long) GetNote() string {
	return this.Note
}

// GetName 取得欄位名稱
func (this *Long) GetName() string {
	return this.Name
}

// GetField 取得欄位類型
func (this *Long) GetField() string {
	return this.Field
}

// FillToMetas 寫入到元資料列表
func (this *Long) FillToMetas(metas Metas, data string) error {
	value, err := strconv.ParseInt(data, internal.Decimal, 64)

	if err != nil {
		return err
	} // if

	metas[this.Name] = value
	return nil
}
