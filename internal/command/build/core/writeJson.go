package core

import (
	"encoding/json"
	"fmt"

	"Sheeter/internal/util"
)

// WriteJson 寫入json
type WriteJson struct {
}

// LongName 取得長名稱
func (this *WriteJson) LongName() string {
	return "json"
}

// ShortName 取得短名稱
func (this *WriteJson) ShortName() string {
	return "j"
}

// Note 取得註解
func (this *WriteJson) Note() string {
	return "generate json file"
}

// Progress 取得進度值
func (this *WriteJson) Progress(sheetSize int) int {
	return sheetSize + 2
}

// Execute 執行工作
func (this *WriteJson) Execute(cargo *Cargo) (filePath string, err error) {
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

	cargo.Progress.Add(1)
	bytes, err := json.MarshalIndent(jsonMaps, JsonPrefix, JsonIndent)

	if err != nil {
		return "", fmt.Errorf("convert json failed: %s [%s]", cargo.LogName(), err)
	} // if

	cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathJson, cargo.JsonFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to json failed: %s [%s]", cargo.LogName(), err)
	} // if

	return filePath, nil
}

type JsonMap map[string]interface{} // json資料列表型態
