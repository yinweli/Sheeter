package core

import (
	"encoding/json"
	"fmt"

	"Sheeter/internal/util"
)

// WriteJson 寫入json
func WriteJson(cargo *Cargo) (filePath string, err error) {
	type jsonMap map[string]interface{}

	var jsonMaps []jsonMap

	for _, itor := range cargo.Columns {
		for row, data := range itor.Datas {
			value, err := itor.Field.Transform(data)

			if err != nil {
				return "", fmt.Errorf("convert value failed: %s [%s(%d) : %s]", cargo.Element.GetFullName(), itor.Name, row, err)
			} // if

			if len(jsonMaps) <= row {
				jsonMaps = append(jsonMaps, jsonMap{})
			} // if

			_ = cargo.Progress.Add(1)
			jsonMaps[row][itor.Name] = value
		} // for
	} // for

	_ = cargo.Progress.Add(1)
	bytes, err := json.MarshalIndent(jsonMaps, "", "    ")

	if err != nil {
		return "", fmt.Errorf("convert json failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathJson, cargo.JsonFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to json failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	return filePath, nil
}
