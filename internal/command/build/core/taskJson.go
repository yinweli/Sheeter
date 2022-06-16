package core

import (
	"encoding/json"
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

const jsonPrefix = ""     // json前綴
const jsonIndent = "    " // json縮排

type box map[string]interface{} // 資料箱形態
type boxMap map[string]box      // 資料箱列表形態

// executeJson 輸出json
func (this *Task) executeJson() error {
	rows := this.getRows(this.global.LineOfData)

	if rows == nil {
		return nil // 找不到資料行, 除了錯誤, 也有可能是碰到空表格
	} // if

	defer func() { _ = rows.Close() }()
	boxMap := make(boxMap)
	row := this.global.LineOfData

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		count := len(datas)
		box := make(box)
		pkey := ""

		for col, itor := range this.columns {
			if itor.Field.IsShow() == false {
				continue
			} // if

			var data string

			if col >= 0 && col < count {
				data = datas[col]
			} // if

			if itor.Field.IsPkey() {
				pkey = data
			} // if

			value, err := itor.Field.ToJsonValue(data)

			if err != nil {
				return fmt.Errorf("convert value failed: %s [%d(%s) : %s]", this.originalName(), row, itor.Name, err)
			} // if

			box[itor.Name] = value
		} // for

		boxMap[pkey] = box
		row++
	} // for

	bytes, err := json.MarshalIndent(boxMap, jsonPrefix, jsonIndent)

	if err != nil {
		return fmt.Errorf("convert json failed: %s [%s]", this.originalName(), err)
	} // if

	err = util.FileWrite(this.jsonFilePath(), bytes, this.global.Bom)

	if err != nil {
		return fmt.Errorf("write to json failed: %s [%s]", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressL)
	} // if

	return nil
}
