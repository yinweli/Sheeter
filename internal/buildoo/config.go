package buildoo

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/internal"
)

// Config 設定資料
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目列表
}

// Global 全域設定
type Global struct {
	ExportJson      bool     `yaml:"exportJson"`      // 是否產生json檔案
	ExportProto     bool     `yaml:"exportProto"`     // 是否產生proto檔案
	Format          bool     `yaml:"format"`          // 是否格式化程式碼
	SimpleNamespace bool     `yaml:"simpleNamespace"` // 是否用簡單的命名空間名稱
	LineOfName      int      `yaml:"lineOfName"`      // 名稱行號(1為起始行)
	LineOfNote      int      `yaml:"lineOfNote"`      // 註解行號(1為起始行)
	LineOfField     int      `yaml:"lineOfField"`     // 欄位行號(1為起始行)
	LineOfLayer     int      `yaml:"lineOfLayer"`     // 階層行號(1為起始行)
	LineOfData      int      `yaml:"lineOfData"`      // 資料行號(1為起始行)
	Excludes        []string `yaml:"excludes"`        // 排除標籤列表
}

// Element 項目設定
type Element struct {
	Excel string `yaml:"excel"` // excel檔案名稱
	Sheet string `yaml:"sheet"` // excel表單名稱
}

// Initialize 初始化設定
func (this *Config) Initialize(cmd *cobra.Command) error {
	flags := cmd.Flags()

	if flags.Changed(flagConfig) {
		if filepath, err := flags.GetString(flagConfig); err == nil {
			datas, err := os.ReadFile(filepath)

			if err != nil {
				return fmt.Errorf("config initialize failed: %w", err)
			} // if

			if err = yaml.Unmarshal(datas, this); err != nil {
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

	if flags.Changed(flagFormat) {
		if value, err := flags.GetBool(flagFormat); err == nil {
			this.Global.Format = value
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

	if flags.Changed(flagExcludes) {
		if items, err := flags.GetStringSlice(flagExcludes); err == nil {
			this.Global.Excludes = append(this.Global.Excludes, items...)
		} // if
	} // if

	if flags.Changed(flagElements) {
		if items, err := flags.GetStringSlice(flagElements); err == nil {
			for _, itor := range items {
				if before, after, ok := strings.Cut(itor, internal.SeparateElement); ok {
					this.Elements = append(this.Elements, Element{
						Excel: before,
						Sheet: after,
					})
				} // if
			} // for
		} // if
	} // if

	return nil
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
