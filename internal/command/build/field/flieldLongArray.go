package field

import "Sheeter/internal/util"

// LongArray 64位元整數陣列
type LongArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *LongArray) TypeExcel() string {
	return "longArray"
}

// TypeCpp 取得c++欄位類型
func (this *LongArray) TypeCpp() string {
	return "std::vector<int64_t>"
}

// TypeCs 取得c#欄位類型
func (this *LongArray) TypeCs() string {
	return "List<long>"
}

// TypeGo 取得go欄位類型
func (this *LongArray) TypeGo() string {
	return "[]int64"
}

// Hide 是否隱藏
func (this *LongArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *LongArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *LongArray) Transform(input string) (result interface{}, err error) {
	return util.StringToInt64Array(input)
}
