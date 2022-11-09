package builds

// Context 執行資料
type Context struct {
	*Global        // 全域設定
	Generate []any // 產生資料列表
	Encoding []any // 編碼資料列表
	Poststep []any // 後製資料列表
}
