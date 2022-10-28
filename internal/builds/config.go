package builds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/internal"
)

// Config 設定資料
type Config struct {
	Global Global   `yaml:"global"` // 全域設定
	Inputs []string `yaml:"inputs"` // 輸入列表
}

// Global 全域設定
type Global struct {
	ExportJson      bool     `yaml:"exportJson"`      // 是否產生json檔案
	ExportProto     bool     `yaml:"exportProto"`     // 是否產生proto檔案
	ExportEnum      bool     `yaml:"exportEnum"`      // 是否產生enum檔案
	SimpleNamespace bool     `yaml:"simpleNamespace"` // 是否用簡單的命名空間名稱
	LineOfName      int      `yaml:"lineOfName"`      // 名稱行號(1為起始行)
	LineOfNote      int      `yaml:"lineOfNote"`      // 註解行號(1為起始行)
	LineOfField     int      `yaml:"lineOfField"`     // 欄位行號(1為起始行)
	LineOfLayer     int      `yaml:"lineOfLayer"`     // 階層行號(1為起始行)
	LineOfData      int      `yaml:"lineOfData"`      // 資料行號(1為起始行)
	LineOfEnum      int      `yaml:"lineOfEnum"`      // 列舉行號(1為起始行)
	Excludes        []string `yaml:"excludes"`        // 排除標籤列表
}

// Sheet 表單資料
type Sheet struct {
	ExcelName string // excel名稱
	SheetName string // sheet名稱
}

// Initialize 初始化設定
func (this *Config) Initialize(cmd *cobra.Command) error {
	flags := cmd.Flags()

	if flags.Changed(flagConfig) {
		if value, err := flags.GetString(flagConfig); err == nil {
			file, err := os.ReadFile(value)

			if err != nil {
				return fmt.Errorf("config initialize failed: %w", err)
			} // if

			if err = yaml.Unmarshal(file, this); err != nil {
				return fmt.Errorf("config initialize failed: %w", err)
			} // if
		} // if
	} // if

	if flags.Changed(flagExportJson) {
		if value, err := flags.GetBool(flagExportJson); err == nil {
			this.Global.ExportJson = value
		} // if
	} // if

	if flags.Changed(flagExportProto) {
		if value, err := flags.GetBool(flagExportProto); err == nil {
			this.Global.ExportProto = value
		} // if
	} // if

	if (this.Global.ExportJson || this.Global.ExportProto) == false { // 如果json與proto都沒輸出, 就會變成通通都輸出
		this.Global.ExportJson = true
		this.Global.ExportProto = true
	} // if

	if flags.Changed(flagExportEnum) {
		if value, err := flags.GetBool(flagExportEnum); err == nil {
			this.Global.ExportEnum = value
		} // if
	} // if

	if flags.Changed(flagSimpleNamespace) {
		if value, err := flags.GetBool(flagSimpleNamespace); err == nil {
			this.Global.SimpleNamespace = value
		} // if
	} // if

	if flags.Changed(flagLineOfName) {
		if value, err := flags.GetInt(flagLineOfName); err == nil {
			this.Global.LineOfName = value
		} // if
	} // if

	if flags.Changed(flagLineOfNote) {
		if value, err := flags.GetInt(flagLineOfNote); err == nil {
			this.Global.LineOfNote = value
		} // if
	} // if

	if flags.Changed(flagLineOfField) {
		if value, err := flags.GetInt(flagLineOfField); err == nil {
			this.Global.LineOfField = value
		} // if
	} // if

	if flags.Changed(flagLineOfLayer) {
		if value, err := flags.GetInt(flagLineOfLayer); err == nil {
			this.Global.LineOfLayer = value
		} // if
	} // if

	if flags.Changed(flagLineOfData) {
		if value, err := flags.GetInt(flagLineOfData); err == nil {
			this.Global.LineOfData = value
		} // if
	} // if

	if flags.Changed(flagLineOfEnum) {
		if value, err := flags.GetInt(flagLineOfEnum); err == nil {
			this.Global.LineOfEnum = value
		} // if
	} // if

	if flags.Changed(flagExcludes) {
		if value, err := flags.GetStringSlice(flagExcludes); err == nil {
			this.Global.Excludes = append(this.Global.Excludes, value...)
		} // if
	} // if

	if flags.Changed(flagInputs) {
		if value, err := flags.GetStringSlice(flagInputs); err == nil {
			this.Inputs = append(this.Inputs, value...)
		} // if
	} // if

	return nil
}

// Path 從輸入列表取得路徑列表
func (this *Config) Path() []string {
	result := []string{}

	for _, itor := range this.Inputs {
		info, err := os.Stat(itor)

		if err != nil {
			continue
		} // if

		if info.IsDir() == false {
			continue
		} // if

		result = append(result, itor)
	} // for

	return result
}

// Excel 從輸入列表取得excel列表
func (this *Config) Excel() []string {
	result := []string{}

	for _, itor := range this.Inputs {
		info, err := os.Stat(itor)

		if err != nil {
			continue
		} // if

		if info.IsDir() {
			continue
		} // if

		if filepath.Ext(itor) != internal.ExcelExt {
			continue
		} // if

		result = append(result, itor)
	} // for

	return result
}

// Sheet 從輸入列表取得sheet列表
func (this *Config) Sheet() []Sheet {
	result := []Sheet{}

	for _, itor := range this.Inputs {
		before, after, ok := strings.Cut(itor, internal.SeparateSheet)

		if ok == false {
			continue
		} // if

		info, err := os.Stat(before)

		if err != nil {
			continue
		} // if

		if info.IsDir() {
			continue
		} // if

		if filepath.Ext(before) != internal.ExcelExt {
			continue
		} // if

		result = append(result, Sheet{
			ExcelName: before,
			SheetName: after,
		})
	} // for

	return result
}

// Check 檢查設定
func (this *Config) Check() error {
	if this.Global.LineOfName <= 0 {
		return fmt.Errorf("config check failed: lineOfName <= 0")
	} // if

	if this.Global.LineOfNote <= 0 {
		return fmt.Errorf("config check failed: lineOfNote <= 0")
	} // if

	if this.Global.LineOfField <= 0 {
		return fmt.Errorf("config check failed: lineOfField <= 0")
	} // if

	if this.Global.LineOfLayer <= 0 {
		return fmt.Errorf("config check failed: lineOfLayer <= 0")
	} // if

	if this.Global.LineOfData <= 0 {
		return fmt.Errorf("config check failed: lineOfData <= 0")
	} // if

	if this.Global.LineOfEnum <= 0 {
		return fmt.Errorf("config check failed: lineOfEnum <= 0")
	} // if

	if this.Global.LineOfName >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfName(%d) >= lineOfData(%d)", this.Global.LineOfName, this.Global.LineOfData)
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfNote(%d) >= lineOfData(%d)", this.Global.LineOfNote, this.Global.LineOfData)
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfField(%d) >= lineOfData(%d)", this.Global.LineOfField, this.Global.LineOfData)
	} // if

	if this.Global.LineOfLayer >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfLayer(%d) >= lineOfData(%d)", this.Global.LineOfLayer, this.Global.LineOfData)
	} // if

	return nil
}
