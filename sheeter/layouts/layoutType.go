package layouts

import (
	"fmt"
	"sort"

	"github.com/emirpasic/gods/stacks/arraystack"

	"github.com/yinweli/Sheeter/sheeter/fields"
	"github.com/yinweli/Sheeter/sheeter/layers"
)

// NewLayoutType 建立類型布局器
func NewLayoutType() *LayoutType {
	return &LayoutType{
		types: map[string]*Type{},
		stack: arraystack.New(),
	}
}

// LayoutType 類型布局器
type LayoutType struct {
	types map[string]*Type  // 類型列表
	stack *arraystack.Stack // 類型堆疊
}

// Type 類型資料
type Type struct {
	Excel  string            // excel檔案名稱
	Sheet  string            // excel表格名稱
	Reader bool              // 是否要產生讀取器
	Fields map[string]*Field // 欄位列表
	Pkey   fields.Field      // 主要索引欄位
}

// Field 欄位資料
type Field struct {
	Name  string       // 欄位名稱
	Note  string       // 欄位註解
	Field fields.Field // 欄位類型
	Alter string       // 欄位類型別名
	Array bool         // 陣列旗標
}

// Begin 開始類型紀錄
func (this *LayoutType) Begin(name, excel, sheet string) error {
	if this.Closure() == false {
		return fmt.Errorf("layoutType begin failed: not closed")
	} // if

	if this.pushType(name, excel, sheet, true) == false {
		return fmt.Errorf("layoutType begin failed: pushType failed")
	} // if

	return nil
}

// End 結束類型紀錄
func (this *LayoutType) End() error {
	if this.pop(1) == false {
		return fmt.Errorf("layoutType end failed: pop failed")
	} // if

	if this.Closure() == false {
		return fmt.Errorf("layoutType end failed: not closed")
	} // if

	return nil
}

// Add 新增欄位
func (this *LayoutType) Add(name, note string, field fields.Field, layer []layers.Layer, back int) error {
	for _, itor := range layer {
		if itor.Type != layers.LayerDivider { // layers.LayerDivider不必處理, 因為結構/陣列未結束
			if this.pushField(itor.Name, "", nil, itor.Name, itor.Type == layers.LayerArray) == false {
				return fmt.Errorf("layoutType add failed: pushField failed")
			} // if

			if this.pushType(itor.Name, itor.Name, "", false) == false {
				return fmt.Errorf("layoutType add failed: pushType failed")
			} // if
		} // if
	} // for

	if this.pushField(name, note, field, "", false) == false {
		return fmt.Errorf("layoutType add failed: pushField failed")
	} // if

	if this.pop(back) == false {
		return fmt.Errorf("layoutType add failed: pop failed")
	} // if

	return nil
}

// Merge 合併類型布局器
func (this *LayoutType) Merge(merge *LayoutType) error {
	if merge.Closure() == false {
		return fmt.Errorf("layoutType merge failed: source not closed")
	} // if

	for typeName, source := range merge.types {
		if _, ok := this.types[typeName]; ok == false {
			this.types[typeName] = &Type{
				Excel:  source.Excel,
				Sheet:  source.Sheet,
				Reader: source.Reader,
				Fields: map[string]*Field{},
				Pkey:   source.Pkey,
			}
		} // if

		target := this.types[typeName]

		for fieldName, field := range source.Fields {
			if _, ok := target.Fields[fieldName]; ok == false {
				target.Fields[fieldName] = &Field{
					Name:  field.Name,
					Note:  field.Note,
					Field: field.Field,
					Alter: field.Alter,
					Array: field.Array,
				}
			} // if
		} // for
	} // for

	if this.Closure() == false {
		return fmt.Errorf("layoutType merge failed: target not closed")
	} // if

	return nil
}

// TypeNames 取得類型名稱列表
func (this *LayoutType) TypeNames() (result []string) {
	for itor := range this.types {
		result = append(result, itor)
	} // for

	sort.Slice(result, func(r, l int) bool {
		return result[r] < result[l]
	})
	return result
}

// Type 取得類型資料
func (this *LayoutType) Type(name string) *Type {
	if value, ok := this.types[name]; ok {
		return value
	} // if

	return nil
}

// Fields 取得類型欄位列表
func (this *LayoutType) Fields(name string) (result []*Field) {
	if value, ok := this.types[name]; ok {
		for _, itor := range value.Fields {
			result = append(result, itor)
		} // for

		sort.Slice(result, func(r, l int) bool {
			return result[r].Name < result[l].Name
		})
	} // if

	return result
}

// FieldNames 取得類型欄位名稱列表
func (this *LayoutType) FieldNames(name string) (result []string) {
	if value, ok := this.types[name]; ok {
		for _, itor := range value.Fields {
			result = append(result, itor.Name)
		} // for

		sort.Slice(result, func(r, l int) bool {
			return result[r] < result[l]
		})
	} // if

	return result
}

// Closure 取得是否閉合
func (this *LayoutType) Closure() bool {
	return this.stack.Empty()
}

// pushType 推入類型
func (this *LayoutType) pushType(name, excel, sheet string, reader bool) bool {
	if _, ok := this.types[name]; ok {
		return false
	} // if

	this.types[name] = &Type{
		Excel:  excel,
		Sheet:  sheet,
		Reader: reader,
		Fields: map[string]*Field{},
	}
	this.stack.Push(name)
	return true
}

// pushField 推入欄位
func (this *LayoutType) pushField(name, note string, field fields.Field, alter string, array bool) bool {
	if field != nil && field.IsShow() == false {
		return true
	} // if

	level, ok := this.stack.Peek()

	if ok == false {
		return false
	} // if

	type_, ok := this.types[level.(string)]

	if ok == false {
		return false
	} // if

	type_.Fields[name] = &Field{
		Name:  name,
		Note:  note,
		Field: field,
		Alter: alter,
		Array: array,
	}

	if field != nil && field.IsPkey() {
		type_.Pkey = field
	} // if

	return true
}

// pop 彈出類型
func (this *LayoutType) pop(count int) bool {
	for i := 0; i < count; i++ {
		if _, ok := this.stack.Pop(); ok == false {
			return false
		} // if
	} // for

	return true
}
