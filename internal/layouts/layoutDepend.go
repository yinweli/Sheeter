package layouts

import (
	"fmt"
	"sort"

	"github.com/emirpasic/gods/sets/hashset"

	"github.com/yinweli/Sheeter/internal/layers"
)

// NewLayoutDepend 建立依賴布局器
func NewLayoutDepend() *LayoutDepend {
	return &LayoutDepend{
		depends: map[string]*hashset.Set{},
	}
}

// LayoutDepend 依賴布局器
type LayoutDepend struct {
	depends map[string]*hashset.Set // 依賴列表
	current *hashset.Set            // 當前依賴列表
}

// Begin 開始類型紀錄
func (this *LayoutDepend) Begin(name string) error {
	if this.current != nil {
		return fmt.Errorf("layoutDepend begin failed, not closed")
	} // if

	if _, ok := this.depends[name]; ok == false {
		this.depends[name] = hashset.New()
	} // if

	this.current = this.depends[name]
	return nil
}

// End 結束類型紀錄
func (this *LayoutDepend) End() error {
	this.current = nil
	return nil
}

// Add 新增依賴
func (this *LayoutDepend) Add(layer []layers.Layer) error {
	if this.current == nil {
		return fmt.Errorf("layoutDepend add failed, not begin")
	} // if

	for _, itor := range layer {
		if itor.Type != layers.LayerDivider { // layers.LayerDivider不必處理, 因為結構/陣列未結束
			this.current.Add(itor.Name)
		} // if
	} // for

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
func (this *LayoutDepend) Depends(name string) (results []string) {
	if depend, ok := this.depends[name]; ok {
		for _, itor := range depend.Values() {
			results = append(results, itor.(string))
		} // for

		sort.Slice(results, func(r, l int) bool {
			return results[r] < results[l]
		})
	} // if

	return results
}
