package buildoo

// Context 執行資料
type Context struct {
	*Global        // 全域設定
	element  []any // 項目列表
	generate []any // 產生資料列表
	encoding []any // 編碼資料列表
	poststep []any // 後製資料列表
}
