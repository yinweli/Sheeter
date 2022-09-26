package layouts

import (
	"fmt"
	"sort"

	"github.com/emirpasic/gods/stacks/arraystack"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
)

// NewLayoutType 建立類型布局器
func NewLayoutType() *LayoutType {
	return &LayoutType{
		types: map[string]*layoutType{},
		stack: arraystack.New(),
	}
}

// LayoutType 類型布局器
type LayoutType struct {
	types map[string]*layoutType // 類型列表
	stack *arraystack.Stack      // 類型堆疊
}

// layoutType 類型資料
type layoutType struct {
	excel  string            // excel檔案名稱
	sheet  string            // excel表格名稱
	reader bool              // 是否要產生讀取器
	fields map[string]*Field // 欄位列表
}

// Field 欄位資料
type Field struct {
	Name  string       // 欄位名稱
	Note  string       // 欄位註解
	Field fields.Field // 欄位類型
	Alter string       // 欄位類型別名
	Array bool         // 陣列旗標
}

// Type 提供給外部使用的類型資料
type Type struct {
	Excel  string   // excel檔案名稱
	Sheet  string   // excel表格名稱
	Reader bool     // 是否要產生讀取器
	Fields []*Field // 欄位列表
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
			this.types[typeName] = &layoutType{
				excel:  source.excel,
				sheet:  source.sheet,
				reader: source.reader,
				fields: map[string]*Field{},
			}
		} // if

		target := this.types[typeName]

		for fieldName, field := range source.fields {
			if _, ok := target.fields[fieldName]; ok == false {
				target.fields[fieldName] = &Field{
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

// Types 取得類型資料
func (this *LayoutType) Types(name string) *Type {
	if value, ok := this.types[name]; ok {
		field := []*Field{}

		for _, itor := range value.fields {
			field = append(field, itor)
		} // for

		sort.Slice(field, func(r, l int) bool {
			return field[r].Name < field[l].Name
		})
		return &Type{
			Excel:  value.excel,
			Sheet:  value.sheet,
			Reader: value.reader,
			Fields: field,
		}
	} // if

	return nil
}

// TypeNames 取得類型名稱列表
func (this *LayoutType) TypeNames() (results []string) {
	for itor := range this.types {
		results = append(results, itor)
	} // for

	sort.Slice(results, func(r, l int) bool {
		return results[r] < results[l]
	})
	return results
}

// FieldNames 取得類型欄位名稱列表
func (this *LayoutType) FieldNames(name string) (results []string) {
	if value, ok := this.types[name]; ok {
		for _, itor := range value.fields {
			results = append(results, itor.Name)
		} // for

		sort.Slice(results, func(r, l int) bool {
			return results[r] < results[l]
		})
	} // if

	return results
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

	this.types[name] = &layoutType{
		excel:  excel,
		sheet:  sheet,
		reader: reader,
		fields: map[string]*Field{},
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

	type_.fields[name] = &Field{
		Name:  name,
		Note:  note,
		Field: field,
		Alter: alter,
		Array: array,
	}
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
