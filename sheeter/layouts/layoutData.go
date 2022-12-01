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
}

// Data 布局資料
type Data struct {
	Name   string         // 欄位名稱
	Field  fields.Field   // 欄位類型
	Tag    string         // 欄位標籤
	Layers []layers.Layer // 階層列表
	Back   int            // 倒退數量
}

// Add 新增布局
func (this *LayoutData) Add(name string, field fields.Field, tag string, layer []layers.Layer, back int) error {
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

	this.layouts = append(this.layouts, Data{
		Name:   name,
		Field:  field,
		Tag:    tag,
		Layers: layer,
		Back:   back,
	})
	return nil
}

// Pack 打包資料
func (this *LayoutData) Pack(datas, excludes []string) (result map[string]interface{}, pkey int64, err error) {
	structor := newStructor()

	for i, itor := range this.layouts {
		if itor.Field.IsShow() == false {
			continue
		} // if

		if isExclude(itor.Tag, excludes) {
			continue
		} // if

		data := utils.GetItem(datas, i)
		value, err := itor.Field.ToJsonValue(data)

		if err != nil {
			return nil, 0, fmt.Errorf("layoutData pack failed: %w", err)
		} // if

		if itor.Field.IsPkey() {
			pkey = value.(int64)
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

// PkeyCount 主要索引數量
func (this *LayoutData) PkeyCount() int {
	count := 0

	for _, itor := range this.layouts {
		if itor.Field.IsPkey() {
			count++
		} // if
	} // for

	return count
}

// isExclude 是否排除標籤
func isExclude(tag string, excludes []string) bool {
	for _, itor := range excludes {
		if itor != "" && itor == tag { // 空標籤是不能被排除的
			return true
		} // if
	} // for

	return false
}