package core

import (
	"bytes"
	"text/template"

	"Sheeter/internal/util"
)

const CppNamespace = "Sheeter" // c++命名空間名稱
const CsNamespace = "Sheeter"  // c#命名空間名稱
const GoPackage = "sheeter"    // go包名

// Coder 程式碼產生器
type Coder struct {
	Columns        []*Column // 欄位列表, 由於要用到程式碼樣板中, 所以得要公開
	cppLibraryPath string    // c++函式庫路徑
	jsonFileName   string    // json檔名
	structName     string    // 結構名稱
	maxline        int       // 最大行數
	curline        int       // 當前行數
}

// Generate 程式碼產生
func (this *Coder) Generate(code string) (results []byte, err error) {
	temp, err := template.New("coder").Parse(code)

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

// CppLibraryPath 取得c++函式庫路徑
func (this *Coder) CppLibraryPath() string {
	return this.cppLibraryPath
}

// JsonFileName 取得json檔名
func (this *Coder) JsonFileName() string {
	return this.jsonFileName
}

// StructName 取得結構名稱
func (this *Coder) StructName() string {
	return this.structName
}

// CppNamespace 取得c++命名空間名稱
func (this *Coder) CppNamespace() string {
	return CppNamespace
}

// CsNamespace 取得c#命名空間名稱
func (this *Coder) CsNamespace() string {
	return CsNamespace
}

// GoPackage 取得go包名
func (this *Coder) GoPackage() string {
	return GoPackage
}

// SetLine 重置換行計數
func (this *Coder) SetLine() string {
	this.curline = 0
	return ""
}

// NewLine 換行輸出
func (this *Coder) NewLine() string {
	defer func() { this.curline++ }()

	if this.maxline > this.curline {
		return "\n"
	} // if

	return ""
}

// ColumnName 取得欄位名稱
func (this *Coder) ColumnName(name string) string {
	return util.FirstUpper(name)
}

// NewCoder 建立程式碼產生器
func NewCoder(columns []*Column, cppLibraryPath string, jsonFileName string, structName string) *Coder {
	return &Coder{
		Columns:        columns,
		cppLibraryPath: cppLibraryPath,
		jsonFileName:   jsonFileName,
		structName:     structName,
		maxline:        calcMaxLine(columns),
	}
}

// calcMaxLine 計算最大行數
func calcMaxLine(columns []*Column) int {
	maxline := -1 // 從-1開始是為了避免多換一次新行

	for _, itor := range columns {
		if itor.Field.IsShow() {
			maxline++
		} // if
	} // for

	return maxline
}
