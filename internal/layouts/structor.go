package layouts

import (
	"container/list"
)

// newStructor 建立結構器
func newStructor() *structor {
	structor := &structor{
		datas: list.New(),
	}
	structor.datas.PushBack(struct_{}) // 結構器從一個結構開始
	return structor
}

// structor 結構器
type structor struct {
	datas *list.List // 資料列表
}

type struct_ = map[string]interface{} // 結構元素類型
type array_ = []struct_               // 陣列元素類型

// pushArray 新增陣列元素
func (this *structor) pushArray(name string) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(struct_); ok {
			object := &array_{}
			layout[name] = object
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// pushStructA 新增結構元素到陣列中
func (this *structor) pushStructA() bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(*array_); ok {
			object := struct_{}
			*layout = append(*layout, object)
			last.Value = layout
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// pushStructS 新增結構元素到結構中
func (this *structor) pushStructS(name string) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(struct_); ok {
			object := struct_{}
			layout[name] = object
			this.datas.PushBack(object)
			return true
		} // if
	} // if

	return false
}

// pushValue 新增值元素
func (this *structor) pushValue(name string, value interface{}) bool {
	if last := this.datas.Back(); last != nil {
		if layout, ok := last.Value.(struct_); ok {
			layout[name] = value
			return true
		} // if
	} // if

	return false
}

// pop 移除元素
func (this *structor) pop(count int, removeArray bool) {
	for i := 0; i < count; i++ {
		if last := this.datas.Back(); last != nil {
			this.datas.Remove(last)
		} // if

		if removeArray { // 如果移除元素後, 發現最後是陣列元素, 就多移除一次
			if last := this.datas.Back(); last != nil {
				if _, ok := last.Value.(*array_); ok {
					this.datas.Remove(last)
				} // if
			} // if
		} // if
	} // for
}

// closure 取得是否閉合
func (this *structor) closure() bool {
	return this.datas.Len() == 1
}

// result 取得結果
func (this *structor) result() struct_ {
	if result, ok := this.datas.Front().Value.(struct_); ok {
		return result
	} // if

	return nil
}
