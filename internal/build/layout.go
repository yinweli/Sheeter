package build

import (
	"fmt"
)

// LayoutBuilder 布局建立器
type LayoutBuilder struct {
	duplField *duplField // 欄位重複檢查器
	duplLayer *duplLayer // 階層重複檢查器
	layouts   []Layout   // 布局列表
}

// Add 新增布局
func (this *LayoutBuilder) Add(name, note string, field Field, layers []Layer, back int) error {
	if this.duplField.Check(name) == false {
		return fmt.Errorf("field duplicate")
	} // if

	if field == nil {
		return fmt.Errorf("field nil")
	} // if

	if this.duplLayer.Check(layers...) == false {
		return fmt.Errorf("layer duplicate")
	} // if

	if back < 0 {
		return fmt.Errorf("back < 0")
	} // if

	this.layouts = append(this.layouts, Layout{
		Name:   name,
		Note:   note,
		Field:  field,
		Layers: layers,
		Back:   back,
	})
	return nil
}

// Pack 打包資料
func (this *LayoutBuilder) Pack(datas []string) (packs map[string]interface{}, pkey string, err error) {
	stacker := NewLayoutStacker()

	for i, itor := range this.layouts {
		if itor.Field.IsShow() == false {
			continue
		} // if

		data := ""

		if i >= 0 && i < len(datas) { // 資料的數量可能因為空白格的關係會短缺, 所以要檢查一下
			data = datas[i]
		} // if

		if itor.Field.IsPkey() {
			if pkey != "" {
				return nil, "", fmt.Errorf("pkey duplicate: %s", itor.Name)
			} // if

			pkey = data
		} // if

		value, err := itor.Field.ToJsonValue(data)

		if err != nil {
			return nil, "", fmt.Errorf("value error: %s\n%w", itor.Name, err)
		} // if

		for _, layer := range itor.Layers {
			if layer.Type == LayerArray {
				if stacker.PushArray(layer.Name) == false || stacker.PushStructA() == false {
					return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
				} // if
			} // if

			if layer.Type == LayerStruct {
				if stacker.PushStructS(layer.Name) == false {
					return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
				} // if
			} // if

			if layer.Type == LayerDivider {
				stacker.Pop(1, false)

				if stacker.PushStructA() == false {
					return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
				} // if
			} // if

			return nil, "", fmt.Errorf("layer unknown: %s", itor.Name)
		} // for

		if stacker.PushValue(itor.Name, value) == false {
			return nil, "", fmt.Errorf("fromat error: %s", itor.Name)
		} // if

		stacker.Pop(itor.Back, true)
	} // for

	return stacker.Result(), pkey, nil
}

// Layouts 取得布局列表
func (this *LayoutBuilder) Layouts() []Layout {
	return this.layouts
}

// Layout 布局資料
type Layout struct {
	Name   string  // 欄位名稱
	Note   string  // 欄位註解
	Field  Field   // 欄位類型
	Layers []Layer // 階層列表
	Back   int     // 倒退數量
}

// NewLayoutBuilder 建立布局建立器
func NewLayoutBuilder() *LayoutBuilder {
	return &LayoutBuilder{
		duplField: NewDuplField(),
		duplLayer: NewDuplLayer(),
		layouts:   []Layout{},
	}
}
