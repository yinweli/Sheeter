package build

import (
	"container/list"
)

// layoutStacker 布局堆疊器
type layoutStacker struct {
	datas *list.List // 資料列表
}

type layoutStruct = map[string]interface{} // 結構布局類型
type layoutArray = []layoutStruct          // 陣列布局類型

// PushArray 新增陣列元素
func (this *layoutStacker) PushArray(name string) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(layoutStruct); ok {
			object := layoutArray{}
			layout[name] = object
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// PushStructA 新增結構元素到陣列中
func (this *layoutStacker) PushStructA() bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(layoutArray); ok {
			object := layoutStruct{}
			layout = append(layout, object)
			last.Value = layout
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// PushStructS 新增結構元素到結構中
func (this *layoutStacker) PushStructS(name string) bool {
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

// PushValue 新增值元素
func (this *layoutStacker) PushValue(name string, value interface{}) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(layoutStruct); ok {
			layout[name] = value
			return true
		} // if
	} // if

	return false
}

// Pop 移除元素
func (this *layoutStacker) Pop(count int, removeArray bool) {
	for i := 0; i < count; i++ {
		if last := this.datas.Back(); last != nil {
			this.datas.Remove(last)
		} // if

		if removeArray { // 如果移除元素後, 發現最後是陣列元素, 就多移除一次
			if last := this.datas.Back(); last != nil {
				if _, ok := last.Value.(layoutArray); ok {
					this.datas.Remove(last)
				} // if
			} // if
		} // if
	} // for
}

// Closure 取得是否閉合
func (this *layoutStacker) Closure() bool {
	return this.datas.Len() == 1
}

// Result 取得結果
func (this *layoutStacker) Result() layoutStruct {
	if result, ok := this.datas.Front().Value.(layoutStruct); ok {
		return result
	} // if

	return nil
}

// NewLayoutStacker 建立布局堆疊器
func NewLayoutStacker() *layoutStacker {
	stacker := &layoutStacker{
		datas: list.New(),
	}
	stacker.datas.PushBack(layoutStruct{}) // 布局堆疊器從一個結構開始
	return stacker
}
