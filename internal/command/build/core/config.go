package core

import (
	"fmt"
)

// Config 編譯設定
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目設定列表
}

// Check 檢查設定是否正確
func (this *Config) Check() error {
	if this.Global.CppLibraryPath == "" {
		return fmt.Errorf("global: cppLibraryPath empty")
	} // if

	if this.Global.LineOfNote <= 0 {
		return fmt.Errorf("global: LineOfNote <= 0")
	} // if

	if this.Global.LineOfField <= 0 {
		return fmt.Errorf("global: LineOfField <= 0")
	} // if

	if this.Global.LineOfData <= 0 {
		return fmt.Errorf("global: LineOfData <= 0")
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return fmt.Errorf("global: LineOfNote(%d) >= LineOfData(%d)", this.Global.LineOfNote, this.Global.LineOfData)
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return fmt.Errorf("global: LineOfField(%d) >= LineOfData(%d)", this.Global.LineOfField, this.Global.LineOfData)
	} // if

	if len(this.Elements) <= 0 {
		return fmt.Errorf("elements empty")
	} // if

	for _, itor := range this.Elements {
		if itor.Excel == "" {
			return fmt.Errorf("element: excel empty")
		} // if

		if itor.Sheet == "" {
			return fmt.Errorf("element: sheet empty")
		} // if
	} // for

	return nil
}

// Global 全域設定
type Global struct {
	ExcelPath      string `yaml:"excelPath"`      // 來源Excel路徑
	CppLibraryPath string `yaml:"cppLibraryPath"` // c++函式庫路徑
	Bom            bool   `yaml:"bom"`            // 輸出的檔案是否使用順序標記(BOM)
	LineOfNote     int    `yaml:"lineOfNote"`     // 註解行號(1為起始行)
	LineOfField    int    `yaml:"lineOfField"`    // 欄位行號(1為起始行)
	LineOfData     int    `yaml:"lineOfData"`     // 資料起始行號(1為起始行)
}

// GetLineOfNote 取得註解行號
func (this *Global) GetLineOfNote() int {
	return this.LineOfNote - 1
}

// GetLineOfField 取得欄位行號
func (this *Global) GetLineOfField() int {
	return this.LineOfField - 1
}

// GetLineOfData 取得資料起始行號
func (this *Global) GetLineOfData() int {
	return this.LineOfData - 1
}

// Element 項目設定
type Element struct {
	Excel string `yaml:"excel"` // Excel檔案名稱
	Sheet string `yaml:"sheet"` // Excel表單名稱
}

// GetFullName 取得完整名稱
func (this *Element) GetFullName() string {
	return fmt.Sprintf("%s(%s)", this.Excel, this.Sheet)
}
