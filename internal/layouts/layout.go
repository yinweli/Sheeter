package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/utils"
)

// NewBuilder 建立布局建造器
func NewBuilder() *Builder {
	return &Builder{
		types:   map[string]int{},
		layouts: []Layout{},
	}
}

// Layout 布局資料
type Layout struct {
	Name   string         // 欄位名稱
	Note   string         // 欄位註解
	Field  fields.Field   // 欄位類型
	Layers []layers.Layer // 階層列表
	Back   int            // 倒退數量
}

// Builder 布局建造器
type Builder struct {
	types   map[string]int // 類型列表
	layouts []Layout       // 布局列表
}

// Add 新增布局
func (this *Builder) Add(name, note string, field fields.Field, layer []layers.Layer, back int) error {
	if name == "" {
		return fmt.Errorf("add builder failed, name empty")
	} // if

	if field == nil {
		return fmt.Errorf("%s: add builder failed, field nil", name)
	} // if

	for _, itor := range layer {
		if type_, ok := this.types[itor.Name]; ok {
			if type_ != itor.Type {
				return fmt.Errorf("%s: add builder failed, layer duplicate", name)
			} // if
		} else {
			this.types[itor.Name] = itor.Type
		} // if
	} // for

	if back < 0 {
		return fmt.Errorf("%s: add builder failed, back < 0", name)
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
func (this *Builder) Pack(datas []string, preset bool) (packs map[string]interface{}, pkey string, err error) {
	stacker := newStacker()

	for i, itor := range this.layouts {
		if itor.Field.IsShow() == false {
			continue
		} // if

		data := utils.GetItem(datas, i)
		value, err := itor.Field.ToJsonValue(data, preset)

		if err != nil {
			return nil, "", fmt.Errorf("%s: pack builder failed, value error: %w", itor.Name, err)
		} // if

		if itor.Field.IsPkey() {
			pkey = fmt.Sprintf("%v", value)
		} // if

		for _, layer := range itor.Layers {
			if layer.Type == layers.LayerArray {
				if stacker.pushArray(layer.Name) == false || stacker.pushStructA() == false {
					return nil, "", fmt.Errorf("%s: pack builder failed, invalid format", itor.Name)
				} // if

				continue
			} // if

			if layer.Type == layers.LayerStruct {
				if stacker.pushStructS(layer.Name) == false {
					return nil, "", fmt.Errorf("%s: pack builder failed, invalid format", itor.Name)
				} // if

				continue
			} // if

			if layer.Type == layers.LayerDivider {
				stacker.pop(1, false)

				if stacker.pushStructA() == false {
					return nil, "", fmt.Errorf("%s: pack builder failed, invalid format", itor.Name)
				} // if

				continue
			} // if

			return nil, "", fmt.Errorf("%s: pack builder failed, unknown layer", itor.Name)
		} // for

		if stacker.pushValue(itor.Name, value) == false {
			return nil, "", fmt.Errorf("%s: pack builder failed, push value", itor.Name)
		} // if

		stacker.pop(itor.Back, true)
	} // for

	if stacker.closure() == false {
		return nil, "", fmt.Errorf("pack builder failed, not closure")
	} // if

	return stacker.result(), pkey, nil
}

// Layouts 取得布局列表
func (this *Builder) Layouts() []Layout {
	return this.layouts
}

// PkeyCount 主要索引數量
func (this *Builder) PkeyCount() int {
	count := 0

	for _, itor := range this.layouts {
		if itor.Field.IsPkey() {
			count++
		} // if
	} // for

	return count
}
