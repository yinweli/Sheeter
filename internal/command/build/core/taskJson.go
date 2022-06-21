package core

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

// executeJson 輸出json
func (this *Task) executeJson() error {
	type jsonObj = map[string]interface{} // json物件型態
	type jsonObjs = map[string]jsonObj    // json列表型態

	rows := this.getRows(this.global.LineOfData)

	if rows == nil {
		return fmt.Errorf("generate json failed: %s\nsheet is empty", this.originalName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := make(jsonObjs)
	row := this.global.LineOfData

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		obj := make(jsonObj)
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

	err := jsonWrite(objs, this.jsonFilePath(), this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate json failed: %s\n%s", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}

// jsonWrite 寫入json檔案, 如果有需要會建立目錄
func jsonWrite(value any, filePath string, bom bool) error {
	bytes, err := json.MarshalIndent(value, "", "    ")

	if err != nil {
		return err
	} // if

	err = os.MkdirAll(path.Dir(filePath), os.ModePerm)

	if err != nil {
		return err
	} // if

	if bom {
		bytes = append([]byte{0xEF, 0xBB, 0xBF}[:], bytes[:]...)
	} // if

	err = ioutil.WriteFile(filePath, bytes, fs.ModePerm)

	if err != nil {
		return err
	} // if

	return nil
}
