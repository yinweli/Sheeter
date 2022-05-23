package field

// Field 欄位介面
type Field interface {
	// TypeExcel 取得excel欄位類型
	TypeExcel() string

	// TypeCpp 取得c++欄位類型
	TypeCpp() string

	// TypeCs 取得c#欄位類型
	TypeCs() string

	// TypeGo 取得go欄位類型
	TypeGo() string

	// Hide 是否隱藏
	Hide() bool

	// PrimaryKey 是否是主要索引
	PrimaryKey() bool

	// FillToMetas 寫入到元資料列表
	FillToMetas(metas Metas, data string) error
}

// Metas 元資料列表型態 TODO: 應該要放到writer中!
type Metas map[string]interface{}

// Data 欄位資料
type Data struct {
	Raw  string // 原始字串
	Note string // 註解名稱
	Name string // 欄位名稱
}
