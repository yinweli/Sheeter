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

	if this.Global.OutputPathJson == "" {
		return errors.New("global: outputPathJson empty")
	} // if

	if this.Global.OutputPathCpp == "" {
		return errors.New("global: outputPathCpp empty")
	} // if

	if this.Global.OutputPathCs == "" {
		return errors.New("global: outputPathCs empty")
	} // if

	if this.Global.OutputPathGo == "" {
		return errors.New("global: outputPathGo empty")
	} // if

	if this.Global.CppLibraryPath == "" {
		return errors.New("global: cppLibraryPath empty")
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
	OutputPathJson string `yaml:"outputPathJson"` // 輸出路徑: json TODO: 改內定的
	OutputPathCpp  string `yaml:"outputPathCpp"`  // 輸出路徑: c++ TODO: 改內定的
	OutputPathCs   string `yaml:"outputPathCs"`   // 輸出路徑: c# TODO: 改內定的
	OutputPathGo   string `yaml:"outputPathGo"`   // 輸出路徑: go TODO: 改內定的
	CppLibraryPath string `yaml:"cppLibraryPath"` // cpp函式庫路徑 TODO: 改內定的
	// TODO: 多一個cs命名空間
	GoPackage   string `yaml:"goPackage"`   // go包名
	Bom         bool   `yaml:"bom"`         // 輸出的檔案是否使用順序標記(BOM)
	LineOfNote  int    `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfField int    `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfData  int    `yaml:"lineOfData"`  // 資料起始行號(1為起始行)
}

// Element 項目設定
type Element struct {
	Excel string `yaml:"excel"` // Excel檔案名稱
	Sheet string `yaml:"sheet"` // Excel表單名稱
}
