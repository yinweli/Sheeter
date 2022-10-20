package builds

// Context 執行資料
type Context struct {
	*Global        // 全域設定
	Element  []any // 項目列表
	Generate []any // 產生資料列表
	Encoding []any // 編碼資料列表
	Poststep []any // 後製資料列表
}

// Close 關閉執行資料
func (this *Context) Close() {
	for _, itor := range this.Element {
		if data, ok := itor.(*initializeElement); ok {
			data.Close()
		} // if
	} // for
}
