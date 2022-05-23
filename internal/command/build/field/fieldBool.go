package field

import "strconv"

// Bool 布林值
type Bool struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *Bool) TypeExcel() string {
	return "bool"
}

// TypeCpp 取得c++欄位類型
func (this *Bool) TypeCpp() string {
	return "bool"
}

// TypeCs 取得c#欄位類型
func (this *Bool) TypeCs() string {
	return "bool"
}

// TypeGo 取得go欄位類型
func (this *Bool) TypeGo() string {
	return "bool"
}

// Hide 是否隱藏
func (this *Bool) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Bool) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *Bool) FillToMetas(metas Metas, data string) error {
	value, err := strconv.ParseBool(data)

	if err != nil {
		return err
	} // if

	metas[this.Name] = value
	return nil
}
