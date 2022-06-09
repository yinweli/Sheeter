package core

import (
	"encoding/json"
	"fmt"

	"Sheeter/internal/util"
)

const jsonPrefix = ""     // json前綴
const jsonIndent = "    " // json縮排

type JsonMap map[string]interface{} // json資料列表型態

// TaskJson 輸出json
func TaskJson(ctx *Context) error {
	rows := ctx.GetRows(ctx.Global.LineOfData)

	if rows == nil {
		return nil // 找不到資料行, 除了錯誤, 也有可能是碰到空表格
	} // if

	defer util.SilentClose(rows)
	var jsonMap []JsonMap
	var row = 0

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			return fmt.Errorf("empty line: %s [%d]", ctx.LogName(), row)
		} // if

		count := len(datas)
		jsonMap = append(jsonMap, JsonMap{})

		for col, itor := range ctx.Columns {
			if itor.Field.IsShow() {
				var data string

				if col >= 0 && col < count {
					data = datas[col]
				} // if

				value, err := itor.Field.Transform(data)

				if err != nil {
					return fmt.Errorf("convert value failed: %s [%d(%s) : %s]", ctx.LogName(), row, itor.Name, err)
				} // if

				jsonMap[row][itor.Name] = value
			} // if
		} // for

		row++
	} // for

	bytes, err := json.MarshalIndent(jsonMap, jsonPrefix, jsonIndent)

	if err != nil {
		return fmt.Errorf("convert json failed: %s [%s]", ctx.LogName(), err)
	} // if

	err = util.FileWrite(ctx.JsonFilePath(), bytes)

	if err != nil {
		return fmt.Errorf("write to json failed: %s [%s]", ctx.LogName(), err)
	} // if

	return nil
}
