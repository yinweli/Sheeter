package core

import (
	"encoding/json"
	"fmt"

	"Sheeter/internal/util"
)

// WriteJson 寫入json
func WriteJson(cargo *Cargo) (filePath string, err error) {
	var jsonMaps []JsonMap

	for _, itor := range cargo.Columns {
		if itor.Field.Show() {
			for row, data := range itor.Datas {
				value, err := itor.Field.Transform(data)

				if err != nil {
					return "", fmt.Errorf("convert value failed: %s [%s(%d) : %s]", cargo.LogName(), itor.Name, row, err)
				} // if

				if len(jsonMaps) <= row {
					jsonMaps = append(jsonMaps, JsonMap{})
				} // if

				cargo.Progress.Add(1)
				jsonMaps[row][itor.Name] = value
			} // for
		} // if
	} // for

	bytes, err := json.MarshalIndent(jsonMaps, JsonPrefix, JsonIndent)

	if err != nil {
		return "", fmt.Errorf("convert json failed: %s [%s]", cargo.LogName(), err)
	} // if

	filePath, err = util.FileWrite(OutputPathJson, cargo.JsonFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to json failed: %s [%s]", cargo.LogName(), err)
	} // if

	return filePath, nil
}

type JsonMap map[string]interface{} // json資料列表型態
