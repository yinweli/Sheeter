package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// executeJson 輸出json
func (this *Task) executeJson() error {
	rows := this.getRows(this.global.LineOfData)

	if rows == nil {
		return fmt.Errorf("generate json failed: %s\nsheet is empty", this.originalName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := make(util.JsonObjs)
	row := this.global.LineOfData

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		obj := make(util.JsonObj)
		pkey := ""

		for col, itor := range this.columns {
			if itor.Field.IsShow() == false {
				continue
			} // if

			var data string

			if col >= 0 && col < len(datas) { // 資料的數量可能因為空白格的關係會短缺, 所以要檢查一下
				data = datas[col]
			} // if

			if itor.Field.IsPkey() {
				pkey = data
			} // if

			value, err := itor.Field.ToJsonValue(data)

			if err != nil {
				return fmt.Errorf("generate json failed: %s [%d(%s)]\n%s", this.originalName(), row, itor.Name, err)
			} // if

			obj[itor.Name] = value
		} // for

		objs[pkey] = obj
		row++
	} // for

	err := util.JsonWrite(objs, this.jsonFilePath(), this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate json failed: %s\n%s", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
