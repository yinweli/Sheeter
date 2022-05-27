package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// WriteJson 寫入json
func WriteJson(cargo *Cargo) error {
	jboxes, err := buildJBoxes(cargo)

	if err != nil {
		return err
	} // if

	jsons, err := buildJson(cargo, jboxes)

	if err != nil {
		return err
	} // if

	err = writeFile(cargo, jsons)

	if err != nil {
		return err
	} // if

	return nil
}

// buildJBoxes 建立json箱列表
func buildJBoxes(cargo *Cargo) (result []jbox, err error) {
	for _, itor := range cargo.Columns {
		for row, data := range itor.Datas {
			value, err := itor.Field.Transform(data)

			if err != nil {
				return nil, fmt.Errorf("convert value failed: %s [%s(%d) : %s]", cargo.Element.GetFullName(), itor.Name, row, err)
			} // if

			if len(result) <= row {
				result = append(result, jbox{})
			} // if

			result[row][itor.Name] = value
			_ = cargo.Progress.Add(1)
		} // for
	} // for

	return result, nil
}

// buildJson 建立json字串
func buildJson(cargo *Cargo, jboxes []jbox) (result []byte, err error) {
	result, err = json.MarshalIndent(jboxes, "", "    ")

	if err != nil {
		return nil, fmt.Errorf("convert json failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)

	return result, nil
}

// writeFile 把json字串寫入檔案
func writeFile(cargo *Cargo, jsons []byte) error {
	excelName := strings.TrimSuffix(cargo.Element.Excel, filepath.Ext(cargo.Element.Excel))
	fileName := fmt.Sprintf("%s%s.json", strings.ToLower(excelName), strings.ToLower(cargo.Element.Sheet))
	fullPath := filepath.Join(OutputPathJson, fileName)
	err := ioutil.WriteFile(fullPath, jsons, 0666)

	if err != nil {
		return fmt.Errorf("write to file failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)

	return nil
}

// jbox json箱
type jbox map[string]interface{}
