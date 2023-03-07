package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter/fields"
	"github.com/yinweli/Sheeter/sheeter/layers"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// NewLayoutData 建立資料布局器
func NewLayoutData() *LayoutData {
	return &LayoutData{
		types:   map[string]int{},
		layouts: []Data{},
	}
}

// LayoutData 資料布局器
type LayoutData struct {
	layouts []Data         // 布局列表
	types   map[string]int // 類型列表
	pkey    fields.Field   // 主要索引欄位
}

// Data 布局資料
type Data struct {
	Name   string         // 欄位名稱
	Field  fields.Field   // 欄位類型
	Layers []layers.Layer // 階層列表
	Back   int            // 倒退數量
	Tag    string         // 標籤字串
}

// Add 新增布局
func (this *LayoutData) Add(name string, field fields.Field, layer []layers.Layer, back int, tag string) error {
	if name == "" {
		return fmt.Errorf("layoutData add failed: name empty")
	} // if

	if field == nil {
		return fmt.Errorf("layoutData add failed: field nil")
	} // if

	for _, itor := range layer {
		if type_, ok := this.types[itor.Name]; ok {
			if type_ != itor.Type {
				return fmt.Errorf("layoutData add failed: layer duplicate")
			} // if
		} else {
			this.types[itor.Name] = itor.Type
		} // if
	} // for

	if back < 0 {
		return fmt.Errorf("layoutData add failed: back < 0")
	} // if

	if field.IsPkey() && this.pkey != nil {
		return fmt.Errorf("layoutData add failed: too many pkey")
	} // if

	if field.IsPkey() {
		this.pkey = field
	} // if

	this.layouts = append(this.layouts, Data{
		Name:   name,
		Field:  field,
		Layers: layer,
		Back:   back,
		Tag:    tag,
	})
	return nil
}

// Pack 打包資料
func (this *LayoutData) Pack(datas []string, tags string) (result map[string]interface{}, pkey any, err error) {
	structor := newStructor()

	for i, itor := range this.layouts {
		if itor.Field.IsShow() == false {
			continue
		} // if

		if utils.TagMatch(tags, itor.Tag) == false {
			continue
		} // if

		data := utils.GetItem(datas, i)
		value, err := itor.Field.ToJsonValue(data)

		if err != nil {
			return nil, 0, fmt.Errorf("layoutData pack failed: %w", err)
		} // if

		if itor.Field.IsPkey() {
			pkey = value
		} // if

		for _, layer := range itor.Layers {
			if layer.Type == layers.LayerArray {
				if structor.pushArray(layer.Name) == false || structor.pushStructA() == false {
					return nil, 0, fmt.Errorf("layoutData pack failed: invalid format")
				} // if

				continue
			} // if

			if layer.Type == layers.LayerStruct {
				if structor.pushStructS(layer.Name) == false {
					return nil, 0, fmt.Errorf("layoutData pack failed: invalid format")
				} // if

				continue
			} // if

			if layer.Type == layers.LayerDivider {
				structor.pop(1, false)

				if structor.pushStructA() == false {
					return nil, 0, fmt.Errorf("layoutData pack failed: invalid format")
				} // if

				continue
			} // if

			return nil, 0, fmt.Errorf("layoutData pack failed: unknown layer")
		} // for

		if structor.pushValue(itor.Name, value) == false {
			return nil, 0, fmt.Errorf("layoutData pack failed: push value")
		} // if

		structor.pop(itor.Back, true)
	} // for

	if structor.closure() == false {
		return nil, 0, fmt.Errorf("layoutData pack failed: not closure")
	} // if

	result = structor.result()

	if result == nil {
		return nil, 0, fmt.Errorf("layoutData pack failed: result nil")
	} // if

	return result, pkey, nil
}

// Pkey 取得主要索引欄位
func (this *LayoutData) Pkey() fields.Field {
	return this.pkey
}
