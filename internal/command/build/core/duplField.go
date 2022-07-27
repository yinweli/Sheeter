package core

import "github.com/emirpasic/gods/sets/hashset"

// DuplField 欄位重複檢查器
type DuplField struct {
	dupls *hashset.Set // 重複列表
}

// Check 重複檢查, true表示正常, false則否
func (this *DuplField) Check(item interface{}) bool {
	if this.dupls.Contains(item) {
		return false
	} // if

	this.dupls.Add(item)
	return true
}

// NewDuplField 建立欄位重複檢查器
func NewDuplField() *DuplField {
	return &DuplField{
		dupls: hashset.New(),
	}
}
