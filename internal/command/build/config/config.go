package config

// Config 編譯設定
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目設定列表
}

// Check 檢查設定是否正確
func (this *Config) Check() (result bool, errors []string) {

	if this.Global.ExcelPath == "" {
		errors = append(errors, "Global: excelPath empty")
	} // if

	if this.Global.OutputPathJson == "" {
		errors = append(errors, "Global: outputPathJson empty")
	} // if

	if this.Global.OutputPathCpp == "" {
		errors = append(errors, "Global: outputPathCpp empty")
	} // if

	if this.Global.OutputPathCs == "" {
		errors = append(errors, "Global: outputPathCs empty")
	} // if

	if this.Global.OutputPathGo == "" {
		errors = append(errors, "Global: outputPathGo empty")
	} // if

	if this.Global.GoPackage == "" {
		errors = append(errors, "Global: goPackage empty")
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		errors = append(errors, "Global: line of note can't greater than line of data")
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		errors = append(errors, "Global: line of field can't greater than line of data")
	} // if

	if len(this.Elements) <= 0 {
		errors = append(errors, "element: element empty")
	} // if

	for _, itor := range this.Elements {
		if itor.ExcelName == "" {
			errors = append(errors, "element: excelName empty")
		} // if

		if itor.SheetName == "" {
			errors = append(errors, "element: sheetName empty")
		} // if

	} // for

	return len(errors) <= 0, errors
}

// Global 全域設定
type Global struct {
	ExcelPath      string `yaml:"excelPath"`      // 來源Excel路徑
	OutputPathJson string `yaml:"outputPathJson"` // 輸出路徑: json
	OutputPathCpp  string `yaml:"outputPathCpp"`  // 輸出路徑: c++
	OutputPathCs   string `yaml:"outputPathCs"`   // 輸出路徑: c#
	OutputPathGo   string `yaml:"outputPathGo"`   // 輸出路徑: go
	CppLibraryPath string `yaml:"cppLibraryPath"` // cpp函式庫路徑
	GoPackage      string `yaml:"goPackage"`      // go包名
	Bom            bool   `yaml:"bom"`            // 輸出的檔案是否使用順序標記(BOM)
	LineOfNote     int    `yaml:"lineOfNote"`     // 註解行號
	LineOfField    int    `yaml:"lineOfField"`    // 欄位行號
	LineOfData     int    `yaml:"lineOfData"`     // 資料起始行號
}

// Element 項目設定
type Element struct {
	ExcelName string `yaml:"excelName"` // Excel檔名
	SheetName string `yaml:"sheetName"` // Excel表單名稱
}
