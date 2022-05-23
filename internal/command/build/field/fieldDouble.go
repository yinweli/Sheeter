package field

import "strconv"

// Double 64位元浮點數
type Double struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *Double) TypeExcel() string {
	return "double"
}

// TypeCpp 取得c++欄位類型
func (this *Double) TypeCpp() string {
	return "double"
}

// TypeCs 取得c#欄位類型
func (this *Double) TypeCs() string {
	return "double"
}

// TypeGo 取得go欄位類型
func (this *Double) TypeGo() string {
	return "float64"
}

// Hide 是否隱藏
func (this *Double) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Double) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *Double) FillToMetas(metas Metas, data string) error {
	value, err := strconv.ParseFloat(data, 64)

	if err != nil {
		return err
	} // if

	metas[this.Name] = value
	return nil
}
