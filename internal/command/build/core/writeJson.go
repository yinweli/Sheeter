package core

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteJson 寫入json
func WriteJson(cargo *Cargo) error {
	type jsonMap map[string]interface{}

	var jsonMaps []jsonMap

	for _, itor := range cargo.Columns {
		for row, data := range itor.Datas {
			value, err := itor.Field.Transform(data)

			if err != nil {
				return fmt.Errorf("convert value failed: %s [%s(%d) : %s]", cargo.Element.GetFullName(), itor.Name, row, err)
			} // if

			if len(jsonMaps) <= row {
				jsonMaps = append(jsonMaps, jsonMap{})
			} // if

			jsonMaps[row][itor.Name] = value
			_ = cargo.Progress.Add(1)
		} // for
	} // for

	bytes, err := json.MarshalIndent(jsonMaps, "", "    ")

	if err != nil {
		return fmt.Errorf("convert json failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)

	filePath := filepath.Join(OutputPathJson, cargo.JsonFileName())
	err = os.MkdirAll(OutputPathJson, os.ModePerm)

	if err != nil {
		return fmt.Errorf("write to file failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	err = ioutil.WriteFile(filePath, bytes, fs.ModePerm)

	if err != nil {
		return fmt.Errorf("write to file failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)

	return nil
}
