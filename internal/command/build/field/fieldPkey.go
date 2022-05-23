package field

import "strconv"

// Pkey 主要索引
type Pkey struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *Pkey) TypeExcel() string {
	return "pkey"
}

// TypeCpp 取得c++欄位類型
func (this *Pkey) TypeCpp() string {
	return "Sheet::pkey" // pkey型態宣告在命名空間Sheet中
}

// TypeCs 取得c#欄位類型
func (this *Pkey) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *Pkey) TypeGo() string {
	return "int"
}

// Hide 是否隱藏
func (this *Pkey) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Pkey) PrimaryKey() bool {
	return true
}

// FillToMetas 寫入到元資料列表
func (this *Pkey) FillToMetas(metas Metas, data string) error {
	value, err := strconv.Atoi(data)

	if err != nil {
		return err
	} // if

	metas[this.Name] = value
	return nil
}
