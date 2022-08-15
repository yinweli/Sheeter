package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/internal/build/layers"
)

// Builder 布局建造器
type Builder struct {
	checker checker  // 階層檢查器
	layouts []Layout // 布局列表
}

// Add 新增布局
func (this *Builder) Add(name, note string, field fields.Field, layer []layers.Layer, back int) error {
	if name == "" {
		return fmt.Errorf("name empty")
	} // if

	if field == nil {
		return fmt.Errorf("field nil: %s", name)
	} // if

	if this.checker.check(layer...) == false {
		return fmt.Errorf("layer duplicate: %s", name)
	} // if

	if back < 0 {
		return fmt.Errorf("back < 0: %s", name)
	} // if

	this.layouts = append(this.layouts, Layout{
		Name:   name,
		Note:   note,
		Field:  field,
		Layers: layer,
		Back:   back,
	})
	return nil
}

// Pack 打包資料
func (this *Builder) Pack(datas []string) (packs map[string]interface{}, pkey string, err error) {
	stacker := newStacker()

	for i, itor := range this.layouts {
		if itor.Field.IsShow() == false {
			continue
		} // if

		data := ""

		if i >= 0 && i < len(datas) { // 資料的數量可能因為空白格的關係會短缺, 所以要檢查一下
			data = datas[i]
		} // if

		if itor.Field.IsPkey() {
			pkey = data
		} // if

		value, err := itor.Field.ToJsonValue(data)

		if err != nil {
			return nil, "", fmt.Errorf("value error: %s\n%w", itor.Name, err)
		} // if

		for _, layer := range itor.Layers {
			if layer.Type == layers.LayerArray {
				if stacker.pushArray(layer.Name) == false || stacker.pushStructA() == false {
					return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
				} // if

				continue
			} // if

			if layer.Type == layers.LayerStruct {
				if stacker.pushStructS(layer.Name) == false {
					return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
				} // if

				continue
			} // if

			if layer.Type == layers.LayerDivider {
				stacker.pop(1, false)

				if stacker.pushStructA() == false {
					return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
				} // if

				continue
			} // if

			return nil, "", fmt.Errorf("layer unknown: %s", itor.Name)
		} // for

		if stacker.pushValue(itor.Name, value) == false {
			return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
		} // if

		stacker.pop(itor.Back, true)
	} // for

	if stacker.closure() == false {
		return nil, "", fmt.Errorf("not closure")
	} // if

	return stacker.result(), pkey, nil
}

// Layouts 取得布局列表
func (this *Builder) Layouts() []Layout {
	return this.layouts
}

// PkeyCount 主要索引數量 // TODO: 單元測試還沒做
func (this *Builder) PkeyCount() int {
	count := 0

	for _, itor := range this.layouts {
		if itor.Field.IsPkey() {
			count++
		} // if
	} // for

	return count
}

// Layout 布局資料
type Layout struct {
	Name   string         // 欄位名稱
	Note   string         // 欄位註解
	Field  fields.Field   // 欄位類型
	Layers []layers.Layer // 階層列表
	Back   int            // 倒退數量
}

// NewBuilder 建立布局建造器
func NewBuilder() *Builder {
	return &Builder{
		checker: checker{},
		layouts: []Layout{},
	}
}
