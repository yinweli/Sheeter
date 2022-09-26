package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/utils"
)

// NewLayoutJson 建立json布局器
func NewLayoutJson() *LayoutJson {
	return &LayoutJson{
		types:   map[string]int{},
		layouts: []layoutJson{},
	}
}

// LayoutJson json布局器
type LayoutJson struct {
	types   map[string]int // 類型列表
	layouts []layoutJson   // 布局列表
}

// layoutJson 布局資料
type layoutJson struct {
	name   string         // 欄位名稱
	field  fields.Field   // 欄位類型
	layers []layers.Layer // 階層列表
	back   int            // 倒退數量
}

// Add 新增布局
func (this *LayoutJson) Add(name string, field fields.Field, layer []layers.Layer, back int) error {
	if name == "" {
		return fmt.Errorf("layoutJson add failed: name empty")
	} // if

	if field == nil {
		return fmt.Errorf("layoutJson add failed: field nil")
	} // if

	for _, itor := range layer {
		if type_, ok := this.types[itor.Name]; ok {
			if type_ != itor.Type {
				return fmt.Errorf("layoutJson add failed: layer duplicate")
			} // if
		} else {
			this.types[itor.Name] = itor.Type
		} // if
	} // for

	if back < 0 {
		return fmt.Errorf("layoutJson add failed: back < 0")
	} // if

	this.layouts = append(this.layouts, layoutJson{
		name:   name,
		field:  field,
		layers: layer,
		back:   back,
	})
	return nil
}

// Pack 打包資料
func (this *LayoutJson) Pack(datas []string, preset bool) (results map[string]interface{}, pkey int64, err error) {
	structor := newStructor()

	for i, itor := range this.layouts {
		if itor.field.IsShow() == false {
			continue
		} // if

		data := utils.GetItem(datas, i)
		value, err := itor.field.ToJsonValue(data, preset)

		if err != nil {
			return nil, 0, fmt.Errorf("layoutJson pack failed: %w", err)
		} // if

		if itor.field.IsPkey() {
			pkey = value.(int64)
		} // if

		for _, layer := range itor.layers {
			if layer.Type == layers.LayerArray {
				if structor.pushArray(layer.Name) == false || structor.pushStructA() == false {
					return nil, 0, fmt.Errorf("layoutJson pack failed: invalid format")
				} // if

				continue
			} // if

			if layer.Type == layers.LayerStruct {
				if structor.pushStructS(layer.Name) == false {
					return nil, 0, fmt.Errorf("layoutJson pack failed: invalid format")
				} // if

				continue
			} // if

			if layer.Type == layers.LayerDivider {
				structor.pop(1, false)

				if structor.pushStructA() == false {
					return nil, 0, fmt.Errorf("layoutJson pack failed: invalid format")
				} // if

				continue
			} // if

			return nil, 0, fmt.Errorf("layoutJson pack failed: unknown layer")
		} // for

		if structor.pushValue(itor.name, value) == false {
			return nil, 0, fmt.Errorf("layoutJson pack failed: push value")
		} // if

		structor.pop(itor.back, true)
	} // for

	if structor.closure() == false {
		return nil, 0, fmt.Errorf("layoutJson pack failed: not closure")
	} // if

	return structor.result(), pkey, nil
}

// PkeyCount 主要索引數量
func (this *LayoutJson) PkeyCount() int {
	count := 0

	for _, itor := range this.layouts {
		if itor.field.IsPkey() {
			count++
		} // if
	} // for

	return count
}
