package field

import "Sheeter/internal/util"

// FloatArray 32位元浮點數陣列
type FloatArray struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *FloatArray) TypeExcel() string {
	return "floatArray"
}

// TypeCpp 取得c++欄位類型
func (this *FloatArray) TypeCpp() string {
	return "std::vector<float>"
}

// TypeCs 取得c#欄位類型
func (this *FloatArray) TypeCs() string {
	return "List<float>"
}

// TypeGo 取得go欄位類型
func (this *FloatArray) TypeGo() string {
	return "[]float32"
}

// Hide 是否隱藏
func (this *FloatArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FloatArray) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *FloatArray) FillToMetas(metas Metas, data string) error {
	values, err := util.StringToFloat32Array(data)

	if err != nil {
		return err
	} // if

	metas[this.Name] = values
	return nil
}
