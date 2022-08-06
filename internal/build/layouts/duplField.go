package layouts

// duplField 欄位重複檢查器
type duplField struct {
	datas map[string]bool // 資料列表
}

// Check 重複檢查
func (this *duplField) Check(item string) bool {
	if _, ok := this.datas[item]; ok {
		return false
	} // if

	this.datas[item] = true
	return true
}

// NewDuplField 建立欄位重複檢查器
func NewDuplField() *duplField {
	return &duplField{
		datas: map[string]bool{},
	}
}
