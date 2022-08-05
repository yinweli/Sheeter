package build

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// runJson 輸出json
func (this *Task) runJson() error {
	row := this.global.LineOfData
	rows, err := this.getRows(row)

	if err != nil {
		return fmt.Errorf("generate json failed: %s\ndata line not found", this.originalName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := util.JsonObjs{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		pkey := ""
		obj := util.JsonObj{}

		for col, itor := range this.columns {
			if itor.Field.IsShow() == false {
				continue
			} // if

			data := ""

			if col >= 0 && col < len(datas) { // 資料的數量可能因為空白格的關係會短缺, 所以要檢查一下
				data = datas[col]
			} // if

			if itor.Field.IsPkey() {
				pkey = data
			} // if

			value, err := itor.Field.ToJsonValue(data)

			if err != nil {
				return fmt.Errorf("generate json failed: %s [%s:%d]\n%w", this.originalName(), itor.Name, row, err)
			} // if

			obj[itor.Name] = value
		} // for

		objs[pkey] = obj
		row++
	} // for

	err = util.JsonWrite(this.jsonFilePath(), objs, this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate json failed: %s\n%w", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
