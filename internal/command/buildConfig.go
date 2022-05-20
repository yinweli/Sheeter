package command

import (
	"Sheeter/internal/logger"
	"Sheeter/internal/util"
)

// BuildConfig 建立表格設定資料
type BuildConfig struct {
	Global   Global    `yaml:"global"`   // 全域設定資料
	Elements []Element `yaml:"elements"` // 項目設定資料列表
}

// Check 檢查設定資料是否正確
func (this *BuildConfig) Check() (result bool) {
	checker := util.NewChecker()
	checker.Add(this.Global.ExcelPath != "", "Global: excelPath empty")
	checker.Add(this.Global.OutputPathJson != "", "Global: outputPathJson empty")
	checker.Add(this.Global.OutputPathGo != "", "Global: outputPathGo empty")
	checker.Add(this.Global.OutputPathCs != "", "Global: outputPathCs empty")
	checker.Add(this.Global.OutputPathCpp != "", "Global: outputPathCpp empty")
	checker.Add(this.Global.LineOfData > this.Global.LineOfNote, "Global: line of note can't greater than line of data")
	checker.Add(this.Global.LineOfData > this.Global.LineOfField, "Global: line of field can't greater than line of data")
	checker.Add(len(this.Elements) != 0, "element: element empty")

	for _, itor := range this.Elements {
		checker.Add(itor.ExcelName != "", "element: excelName empty")
		checker.Add(itor.SheetName != "", "element: sheetName empty")
	} // for

	for _, itor := range checker.Errors() {
		logger.Error(itor)
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

// Element 項目設定資料
type Element struct {
	ExcelName string `yaml:"excelName"` // Excel檔名
	SheetName string `yaml:"sheetName"` // Excel表單名稱
}
