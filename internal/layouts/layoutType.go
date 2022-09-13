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
		level: arraystack.New(),
	}
}

// Merge 合併類型布局器
func Merge(layoutTypes ...*LayoutType) (result *LayoutType, err error) {
	result = NewLayoutType()

	for _, itor := range layoutTypes {
		if itor.Closure() == false {
			return nil, fmt.Errorf("layoutType merge failed, not closed")
		} // if

		for key, source := range itor.types {
			if _, ok := result.types[key]; ok == false {
				result.types[key] = &layoutType{
					fields: map[string]*Field{},
					reader: source.reader,
				}
			} // if

			target := result.types[key]

			for name, field := range source.fields {
				if _, ok := target.fields[name]; ok == false {
					target.fields[name] = &Field{
						Name:  field.Name,
						Note:  field.Note,
						Field: field.Field,
						Alter: field.Alter,
						Array: field.Array,
					}
				} // if
			} // for
		} // for
	} // for

	if result.Closure() == false {
		return nil, fmt.Errorf("layoutType merge failed, not closed")
	} // if

	return result, nil
}

// LayoutType 類型布局器
type LayoutType struct {
	types map[string]*layoutType // 類型列表
	level *arraystack.Stack      // 類型堆疊
}

// layoutType 布局資料
type layoutType struct {
	fields map[string]*Field // 欄位列表
	reader bool              // 是否需要建立讀取器
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
func (this *LayoutType) Begin(name string) error {
	if this.Closure() == false {
		return fmt.Errorf("layoutType begin failed, not closed")
	} // if

	this.pushType(name, true)
	return nil
}

// End 結束類型紀錄
func (this *LayoutType) End() error {
	if this.pop(1) == false {
		return fmt.Errorf("layoutType end failed, pop failed")
	} // if

	return nil
}

// Add 新增欄位
func (this *LayoutType) Add(name, note string, field fields.Field, layer []layers.Layer, back int) error {
	for _, itor := range layer {
		if itor.Type != layers.LayerDivider { // layers.LayerDivider不必處理, 因為結構/陣列未結束
			if this.pushField(itor.Name, "", nil, itor.Name, itor.Type == layers.LayerArray) == false {
				return fmt.Errorf("layoutType add failed, pushField failed")
			} // if

			if this.pushType(itor.Name, false) == false {
				return fmt.Errorf("layoutType add failed, pushType failed")
			} // if
		} // if
	} // for

	if this.pushField(name, note, field, "", false) == false {
		return fmt.Errorf("layoutType add failed, pushField failed")
	} // if

	if this.pop(back) == false {
		return fmt.Errorf("layoutType add failed, pop failed")
	} // if

	return nil
}

// TypeNames 取得類型名稱列表
func (this *LayoutType) TypeNames() (results []string) {
	for key := range this.types {
		results = append(results, key)
	} // for

	sort.Slice(results, func(r, l int) bool {
		return results[r] < results[l]
	})
	return results
}

// FieldNames 取得類型欄位名稱列表
func (this *LayoutType) FieldNames(name string) (results []string) {
	if types, ok := this.types[name]; ok {
		for _, itor := range types.fields {
			results = append(results, itor.Name)
		} // for

		sort.Slice(results, func(r, l int) bool {
			return results[r] < results[l]
		})
	} // if

	return results
}

// FieldDetails 取得類型欄位詳細列表
func (this *LayoutType) FieldDetails(name string) (results []*Field, reader bool, err error) {
	types, ok := this.types[name]

	if ok == false {
		return nil, false, fmt.Errorf("layoutType fieldDetails failed, type not exist")
	} // if

	for _, itor := range types.fields {
		results = append(results, itor)
	} // for

	sort.Slice(results, func(r, l int) bool {
		return results[r].Name < results[l].Name
	})
	return results, types.reader, nil
}

// Closure 取得是否閉合
func (this *LayoutType) Closure() bool {
	return this.level.Empty()
}

// pushType 推入類型
func (this *LayoutType) pushType(name string, reader bool) bool {
	if _, ok := this.types[name]; ok {
		return false
	} // if

	this.types[name] = &layoutType{
		fields: map[string]*Field{},
		reader: reader,
	}
	this.level.Push(name)
	return true
}

// pushField 推入欄位
func (this *LayoutType) pushField(name, note string, field fields.Field, alter string, array bool) bool {
	level, ok := this.level.Peek()

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
		if _, ok := this.level.Pop(); ok == false {
			return false
		} // if
	} // for

	return true
}
