package config

import "errors"

// Config 編譯設定
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目設定列表
}

// Check 檢查設定是否正確
func (this *Config) Check() error {
	if this.Global.ExcelPath == "" {
		return errors.New("global: excelPath empty")
	} // if

	if this.Global.CppLibraryPath == "" {
		return errors.New("global: cppLibraryPath empty")
	} // if

	if this.Global.CsNamespace == "" {
		return errors.New("global: csNamespace empty")
	} // if

	if this.Global.GoPackage == "" {
		return errors.New("global: goPackage empty")
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return errors.New("global: line of note can't greater than line of data")
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return errors.New("global: line of field can't greater than line of data")
	} // if

	if len(this.Elements) <= 0 {
		return errors.New("elements empty")
	} // if

	for _, itor := range this.Elements {
		if itor.Excel == "" {
			return errors.New("element: excel empty")
		} // if

		if itor.Sheet == "" {
			return errors.New("element: sheet empty")
		} // if
	} // for

	return nil
}

// Global 全域設定
type Global struct {
	ExcelPath      string `yaml:"excelPath"`      // 來源Excel路徑
	CppLibraryPath string `yaml:"cppLibraryPath"` // c++函式庫路徑
	CsNamespace    string `yaml:"csNamespace"`    // c#命名空間名稱
	GoPackage      string `yaml:"goPackage"`      // go包名
	Bom            bool   `yaml:"bom"`            // 輸出的檔案是否使用順序標記(BOM)
	LineOfNote     int    `yaml:"lineOfNote"`     // 註解行號(1為起始行)
	LineOfField    int    `yaml:"lineOfField"`    // 欄位行號(1為起始行)
	LineOfData     int    `yaml:"lineOfData"`     // 資料起始行號(1為起始行)
}

// Element 項目設定
type Element struct {
	Excel string `yaml:"excel"` // Excel檔案名稱
	Sheet string `yaml:"sheet"` // Excel表單名稱
}
