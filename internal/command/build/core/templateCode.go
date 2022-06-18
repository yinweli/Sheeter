package core

import (
	"bytes"
	"text/template"
)

// TemplateCode 程式碼產生器樣板
type TemplateCode struct {
	Template
	Namespace    string    // 命名空間名稱
	JsonFileName string    // json檔名
	Columns      []*Column // 欄位列表, 由於要用到程式碼樣板中, 所以得要公開
}

// SetLine 設置行數
func (this *TemplateCode) SetLine() string {
	maxline := -1 // 從-1開始是為了避免多換一次新行

	for _, itor := range this.Columns {
		if itor.Field.IsShow() {
			maxline++
		} // if
	} // for

	return this.Template.SetLine(maxline)
}

// Generate 程式碼產生
func (this *TemplateCode) Generate(code string) (results []byte, err error) {
	temp, err := template.New("templateCode").Parse(code)

	if err != nil {
		return nil, err
	} // if

	buffer := &bytes.Buffer{}
	err = temp.Execute(buffer, this)

	if err != nil {
		return nil, err
	} // if

	return buffer.Bytes(), nil
}

// NewTemplateCode 建立程式碼產生器樣板
func NewTemplateCode(originalName string, structName string, namespace string, jsonFileName string) TemplateCode {
	return TemplateCode{
		Template:     NewTemplate(originalName, structName),
		Namespace:    namespace,
		JsonFileName: jsonFileName,
		Columns:      nil,
	}
}
