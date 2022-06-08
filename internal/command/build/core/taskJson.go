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
	var jsonMap []JsonMap

	for _, itor := range ctx.Columns {
		if itor.Field.IsShow() {
			for row, data := range itor.Datas {
				value, err := itor.Field.Transform(data)

				if err != nil {
					return fmt.Errorf("convert value failed: %s [%s(%d) : %s]", ctx.LogName(), itor.Name, row, err)
				} // if

				if len(jsonMap) <= row {
					jsonMap = append(jsonMap, JsonMap{})
				} // if

				_ = ctx.Progress.Add(1)
				jsonMap[row][itor.Name] = value
			} // for
		} // if
	} // for

	bytes, err := json.MarshalIndent(jsonMap, jsonPrefix, jsonIndent)

	if err != nil {
		return fmt.Errorf("convert json failed: %s [%s]", ctx.LogName(), err)
	} // if

	err = util.FileWrite(ctx.JsonFilePath(), bytes)

	if err != nil {
		return fmt.Errorf("write to json failed: %s [%s]", ctx.LogName(), err)
	} // if

	_ = ctx.Progress.Add(1)
	return nil
}
