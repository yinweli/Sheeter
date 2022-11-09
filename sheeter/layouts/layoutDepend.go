package layouts

import (
	"fmt"
	"sort"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"

	"github.com/yinweli/Sheeter/sheeter/layers"
)

// NewLayoutDepend 建立依賴布局器
func NewLayoutDepend() *LayoutDepend {
	return &LayoutDepend{
		depends: map[string]*hashset.Set{},
		stack:   arraystack.New(),
	}
}

// LayoutDepend 依賴布局器
type LayoutDepend struct {
	depends map[string]*hashset.Set // 依賴列表
	stack   *arraystack.Stack       // 類型堆疊
}

// Begin 開始類型紀錄
func (this *LayoutDepend) Begin(name string) error {
	if this.Closure() == false {
		return fmt.Errorf("layoutDepend begin failed: not closed")
	} // if

	this.stack.Push(name)
	return nil
}

// End 結束類型紀錄
func (this *LayoutDepend) End() error {
	if this.pop(1) == false {
		return fmt.Errorf("layoutDepend end failed: pop failed")
	} // if

	if this.Closure() == false {
		return fmt.Errorf("layoutDepend end failed: not closed")
	} // if

	return nil
}

// Add 新增依賴
func (this *LayoutDepend) Add(layer []layers.Layer, back int) error {
	for _, itor := range layer {
		if itor.Type != layers.LayerDivider { // layers.LayerDivider不必處理, 因為結構/陣列未結束
			if this.push(itor.Name) == false {
				return fmt.Errorf("layoutDepend add failed: push failed")
			} // if
		} // if
	} // for

	if this.pop(back) == false {
		return fmt.Errorf("layoutDepend add failed: pop failed")
	} // if

	return nil
}

// Merge 合併依賴布局器
func (this *LayoutDepend) Merge(merge *LayoutDepend) error {
	for name, source := range merge.depends {
		if _, ok := this.depends[name]; ok == false {
			this.depends[name] = hashset.New()
		} // if

		target := this.depends[name]
		target.Add(source.Values()...)
	} // for

	return nil
}

// Depends 取得依賴列表
func (this *LayoutDepend) Depends(name string) (result []string) {
	if depend, ok := this.depends[name]; ok {
		for _, itor := range depend.Values() {
			result = append(result, itor.(string))
		} // for

		sort.Slice(result, func(r, l int) bool {
			return result[r] < result[l]
		})
	} // if

	return result
}

// Closure 取得是否閉合
func (this *LayoutDepend) Closure() bool {
	return this.stack.Empty()
}

// push 推入依賴
func (this *LayoutDepend) push(name string) bool {
	level, ok := this.stack.Peek()

	if ok == false {
		return false
	} // if

	depend := level.(string)

	if _, ok := this.depends[depend]; ok == false {
		this.depends[depend] = hashset.New()
	} // if

	this.depends[depend].Add(name)
	this.stack.Push(name)
	return true
}

// pop 彈出依賴
func (this *LayoutDepend) pop(count int) bool {
	for i := 0; i < count; i++ {
		if _, ok := this.stack.Pop(); ok == false {
			return false
		} // if
	} // for

	return true
}
