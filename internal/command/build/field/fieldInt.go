package field

import "strconv"

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
	return "int"
}

// TypeCs 取得c#欄位類型
func (this *Int) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *Int) TypeGo() string {
	return "int"
}

// Hide 是否隱藏
func (this *Int) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Int) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *Int) FillToMetas(metas Metas, data string) error {
	value, err := strconv.ParseInt(data, 10, 32) // 32位元整數

	if err != nil {
		return err
	} // if

	metas[this.Name] = int(value)
	return nil
}
