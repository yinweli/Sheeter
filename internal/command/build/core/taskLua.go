package core

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/yinweli/Sheeter/internal/util"
)

const luaCode = `{{$.StructName}} = { {{range $Pkey, $Data := $.Objs}}
[{{$Pkey}}] = { {{range $Name, $Value := $Data}}{{$Name}} = {{$Value}}, {{end}} },{{end}}
}`

// luaData lua產生資料
type luaData struct {
	util.TempTool
	StructName string       // 結構名稱
	Objs       util.LuaObjs // 內容資料列表
}

// runLua 輸出lua
func (this *Task) runLua() error {
	row := this.global.LineOfData
	rows, err := this.getRows(row)

	if err != nil {
		return fmt.Errorf("generate lua failed: %s\ndata line not found", this.originalName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := util.LuaObjs{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		pkey := ""
		obj := util.LuaObj{}

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

			value, err := itor.Field.ToLuaValue(data)

			if err != nil {
				return fmt.Errorf("generate lua failed: %s [%s:%d]\n%s", this.originalName(), itor.Name, row, err)
			} // if

			obj[itor.Name] = value
		} // for

		objs[pkey] = obj
		row++
	} // for

	temp, err := template.New("luaCode").Parse(luaCode)

	if err != nil {
		return fmt.Errorf("generate lua failed: %s\n%s", this.originalName(), err)
	} // if

	buffer := &bytes.Buffer{}
	data := &luaData{
		StructName: this.structName(),
		Objs:       objs,
	}
	err = temp.Execute(buffer, data)

	if err != nil {
		return fmt.Errorf("generate lua failed: %s\n%s", this.originalName(), err)
	} // if

	err = util.FileWrite(this.luaFilePath(), buffer.Bytes(), this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate lua failed: %s\n%s", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
