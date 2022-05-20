package field

// Field 欄位介面
type Field interface {
	// StyleJson 取得json欄位類型
	StyleJson() string

	// StyleCpp 取得c++欄位類型
	StyleCpp() string

	// StyleCs 取得c#欄位類型
	StyleCs() string

	// StyleGo 取得go欄位類型
	StyleGo() string

	// Hide 是否隱藏
	Hide() bool

	// PrimaryKey 是否是主要索引
	PrimaryKey() bool

	// FillToMetas 寫入到元資料列表
	FillToMetas(metas Metas, value string) error
}

// Metas 元資料列表型態
type Metas map[string]interface{}

// Data 欄位資料
type Data struct {
	Raw  string // 原始字串
	Note string // 註解名稱
	Name string // 欄位名稱
}
