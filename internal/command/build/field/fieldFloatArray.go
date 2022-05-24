package field

import "Sheeter/internal/util"

// FloatArray 32位元浮點數陣列
type FloatArray struct {
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

// Transform 字串轉換
func (this *FloatArray) Transform(input string) (result interface{}, err error) {
	return util.StringToFloat32Array(input)
}
