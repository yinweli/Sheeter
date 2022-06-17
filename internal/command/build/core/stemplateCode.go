package core

import (
	"bytes"
	"text/template"
)

const CsNamespace = "Sheeter" // c#命名空間名稱
const GoPackage = "sheeter"   // go包名

type STemplateCode struct {
	STemplate
	JsonFileName string    // json檔名
	Columns      []*Column // 欄位列表, 由於要用到程式碼樣板中, 所以得要公開
}

// CsNamespace 取得c#命名空間名稱
func (this *STemplateCode) CsNamespace() string {
	return CsNamespace
}

// GoPackage 取得go包名
func (this *STemplateCode) GoPackage() string {
	return GoPackage
}

// SetLine 設置行數
func (this *STemplateCode) SetLine() string {
	maxline := -1 // 從-1開始是為了避免多換一次新行

	for _, itor := range this.Columns {
		if itor.Field.IsShow() {
			maxline++
		} // if
	} // for

	return this.STemplate.SetLine(maxline)
}

// Generate 程式碼產生
func (this *STemplateCode) Generate(code string) (results []byte, err error) {
	temp, err := template.New("stemplateCode").Parse(code)

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

func NewSTemplateCode(originalName string, structName string) *STemplateCode {
	return &STemplateCode{
		STemplate:    *NewSTemplate(originalName, structName),
		JsonFileName: "",
		Columns:      nil,
	}
}
