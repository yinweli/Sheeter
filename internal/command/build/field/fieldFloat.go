package field

import "strconv"

// Float 布林值
type Float struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *Float) TypeExcel() string {
	return "float"
}

// TypeCpp 取得c++欄位類型
func (this *Float) TypeCpp() string {
	return "float"
}

// TypeCs 取得c#欄位類型
func (this *Float) TypeCs() string {
	return "float"
}

// TypeGo 取得go欄位類型
func (this *Float) TypeGo() string {
	return "float32"
}

// Hide 是否隱藏
func (this *Float) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Float) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *Float) FillToMetas(metas Metas, data string) error {
	value, err := strconv.ParseFloat(data, 32) // 32位元浮點數, 如果要更多的小數點, 就得用floatDouble了

	if err != nil {
		return err
	} // if

	metas[this.Name] = value
	return nil
}
