package util

import "github.com/emirpasic/gods/sets/hashset"

// Duplicate 重複檢查器
type Duplicate struct {
	datas *hashset.Set // 資料列表
}

// Check 重複檢查, true表示正常, false則否
func (this *Duplicate) Check(item interface{}) bool {
	if this.datas.Contains(item) {
		return false
	} // if

	this.datas.Add(item)
	return true
}

// NewDuplicate 建立重複檢查器
func NewDuplicate() *Duplicate {
	return &Duplicate{
		datas: hashset.New(),
	}
}
