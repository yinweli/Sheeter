package config

import (
	"Sheeter/internal/util"
)

// Config 設定資料
type Config struct {
	Global   Global    `yaml:"Global"`   // 全域設定資料
	Elements []Element `yaml:"elements"` // 項目設定資料列表
}

// Check 檢查設定資料是否正確
func (this *Config) Check() (result bool) {
	checker := util.NewChecker()
	checker.Add(this.Global.Check(), "")
	checker.Add(len(this.Elements) <= 0, "element: element empty")

	for _, itor := range this.Elements {
		checker.Add(itor.Check(), "")
	} // for

	return checker.Result()
}

// Global 全域設定資料
type Global struct {
	ExcelPath      string `yaml:"excelPath"`      // 來源Excel路徑
	OutputPathJson string `yaml:"outputPathJson"` // 輸出路徑: json
	OutputPathGo   string `yaml:"outputPathGo"`   // 輸出路徑: go
	OutputPathCs   string `yaml:"outputPathCs"`   // 輸出路徑: c#
	OutputPathCpp  string `yaml:"outputPathCpp"`  // 輸出路徑: c++
	CppLibraryPath string `yaml:"cppLibraryPath"` // 函式庫路徑: cpp
	Bom            bool   `yaml:"bom"`            // 輸出的檔案是否使用順序標記(BOM)
	LineOfNote     int    `yaml:"lineOfNote"`     // 註解行號
	LineOfField    int    `yaml:"lineOfField"`    // 欄位行號
	LineOfData     int    `yaml:"lineOfData"`     // 資料起始行號
}

// Check 檢查全域設定資料是否正確
func (this *Global) Check() (result bool) {
	checker := util.NewChecker()
	checker.Add(this.ExcelPath == "", "Global: excelPath empty")
	checker.Add(this.OutputPathJson == "", "Global: outputPathJson empty")
	checker.Add(this.OutputPathGo == "", "Global: outputPathGo empty")
	checker.Add(this.OutputPathCs == "", "Global: outputPathCs empty")
	checker.Add(this.OutputPathCpp == "", "Global: outputPathCpp empty")
	checker.Add(this.LineOfNote >= this.LineOfData, "Global: line of note can't greater than line of data")
	checker.Add(this.LineOfField >= this.LineOfData, "Global: line of field can't greater than line of data")

	return checker.Result()
}

// Element 項目設定資料
type Element struct {
	ExcelName string `yaml:"excelName"` // Excel檔名
	SheetName string `yaml:"sheetName"` // Excel表單名稱
}

// Check 檢查項目設定資料是否正確
func (this *Element) Check() (result bool) {
	checker := util.NewChecker()
	checker.Add(this.ExcelName == "", "element: excelName empty")
	checker.Add(this.SheetName == "", "element: sheetName empty")

	return checker.Result()
}
