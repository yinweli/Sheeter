package layouts

import (
	"container/list"
)

// stacker 布局堆疊器
type stacker struct {
	datas *list.List // 資料列表
}

type layoutStruct = map[string]interface{} // 結構布局類型
type layoutArray = []layoutStruct          // 陣列布局類型

// pushArray 新增陣列元素
func (this *stacker) pushArray(name string) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(layoutStruct); ok {
			object := &layoutArray{}
			layout[name] = object
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// pushStructA 新增結構元素到陣列中
func (this *stacker) pushStructA() bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(*layoutArray); ok {
			object := layoutStruct{}
			*layout = append(*layout, object)
			last.Value = layout
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// pushStructS 新增結構元素到結構中
func (this *stacker) pushStructS(name string) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(layoutStruct); ok {
			object := layoutStruct{}
			layout[name] = object
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// pushValue 新增值元素
func (this *stacker) pushValue(name string, value interface{}) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(layoutStruct); ok {
			layout[name] = value
			return true
		} // if
	} // if

	return false
}

// pop 移除元素
func (this *stacker) pop(count int, removeArray bool) {
	for i := 0; i < count; i++ {
		if last := this.datas.Back(); last != nil {
			this.datas.Remove(last)
		} // if

		if removeArray { // 如果移除元素後, 發現最後是陣列元素, 就多移除一次
			if last := this.datas.Back(); last != nil {
				if _, ok := last.Value.(*layoutArray); ok {
					this.datas.Remove(last)
				} // if
			} // if
		} // if
	} // for
}

// closure 取得是否閉合
func (this *stacker) closure() bool {
	return this.datas.Len() == 1
}

// result 取得結果
func (this *stacker) result() layoutStruct {
	if result, ok := this.datas.Front().Value.(layoutStruct); ok {
		return result
	} // if

	return nil
}

// newStacker 建立布局堆疊器
func newStacker() *stacker {
	stacker := &stacker{
		datas: list.New(),
	}
	stacker.datas.PushBack(layoutStruct{}) // 布局堆疊器從一個結構開始
	return stacker
}
