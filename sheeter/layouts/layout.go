package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/fields"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// NewLayout 建立資料布局器
func NewLayout(lineTag, lineName, lineNote, lineField []string) (result *Layout, failed []string) {
	result = &Layout{
		layout: map[int]*LayoutData{},
	}
	duplicate := map[string]bool{}

	for i, itor := range lineTag { // 尋訪時, 以標籤行為主
		if i == sheeter.IndexOutput { // 跳過輸出欄
			continue
		} // if

		if i != sheeter.IndexPrimary && utils.CheckIgnore(itor) { // 跳過忽略欄
			continue
		} // if

		if i != sheeter.IndexPrimary && itor == "" { // 一旦遇到空欄位, 就結束布局
			break
		} // if

		fail := []string{}
		name := utils.At(lineName, i)
		note := utils.At(lineNote, i)
		field, err := fields.Parser(utils.At(lineField, i))

		if err != nil {
			fail = append(fail, fmt.Sprintf("column %v name %v: %v", i, name, err))
		} // if

		if utils.CheckField(name) == false {
			fail = append(fail, fmt.Sprintf("column %v name %v: name invalid", i, name))
		} // if

		if duplicate[name] {
			fail = append(fail, fmt.Sprintf("column %v name %v: name duplicate", i, name))
		} // if

		if len(fail) == 0 {
			result.layout[i] = &LayoutData{
				Tag:   itor,
				Name:  name,
				Note:  note,
				Field: field,
			}
			duplicate[name] = true
		} else {
			failed = append(failed, fail...)
		} // if
	} // for

	if result.layout[sheeter.IndexPrimary] == nil {
		failed = append(failed, "primary column missing")
	} // if

	return result, failed
}

// Layout 資料布局器
type Layout struct {
	layout map[int]*LayoutData // 布局列表
}

// LayoutData 布局資料
type LayoutData struct {
	Tag   string       // 標籤字串
	Name  string       // 欄位名稱
	Note  string       // 欄位註解
	Field fields.Field // 欄位類型
}

// Primary 取得主索引布局資料
func (this *Layout) Primary() *LayoutData {
	return this.layout[sheeter.IndexPrimary]
}

// Select 取得用標籤篩選過的布局資料
func (this *Layout) Select(tag string) (result []*LayoutData) {
	for i, itor := range this.layout {
		if i == sheeter.IndexPrimary || utils.CheckTag(tag, itor.Tag) {
			result = append(result, itor)
		} // if
	} // for

	return result
}

// Pack 打包資料
func (this *Layout) Pack(tag string, data []string) (primary any, pack map[string]interface{}, err error) {
	pack = map[string]interface{}{}

	for i, itor := range this.layout {
		if i != sheeter.IndexPrimary && utils.CheckTag(tag, itor.Tag) == false {
			continue
		} // if

		value, err := itor.Field.ToJsonValue(utils.At(data, i))

		if err != nil {
			return nil, nil, fmt.Errorf("layout pack: column %v name %v: %w", i, itor.Name, err)
		} // if

		if i == sheeter.IndexPrimary {
			primary = value
		} // if

		pack[itor.Name] = value
	} // for

	return primary, pack, nil
}
