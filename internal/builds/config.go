package builds

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Config 設定資料
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目設定列表
}

// Global 全域設定
type Global struct {
	GenerateJson  bool     `yaml:"generateJson"`  // 是否產生json檔案
	GenerateProto bool     `yaml:"generateProto"` // 是否產生proto檔案
	LineOfField   int      `yaml:"lineOfField"`   // 欄位行號(1為起始行)
	LineOfLayer   int      `yaml:"lineOfLayer"`   // 階層行號(1為起始行)
	LineOfNote    int      `yaml:"lineOfNote"`    // 註解行號(1為起始行)
	LineOfData    int      `yaml:"lineOfData"`    // 資料行號(1為起始行)
	Excludes      []string `yaml:"excludes"`      // 排除標籤列表
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

	if flags.Changed(flagGenerateJson) {
		if value, err := flags.GetBool(flagGenerateJson); err == nil {
			this.Global.GenerateJson = value
		} // if
	} // if

	if flags.Changed(flagGenerateProto) {
		if value, err := flags.GetBool(flagGenerateProto); err == nil {
			this.Global.GenerateProto = value
		} // if
	} // if

	if (this.Global.GenerateJson || this.Global.GenerateProto) == false { // 如果json與proto都沒輸出, 就會變成通通都輸出
		this.Global.GenerateJson = true
		this.Global.GenerateProto = true
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

	if flags.Changed(flagLineOfNote) {
		if value, err := flags.GetInt(flagLineOfNote); err == nil {
			this.Global.LineOfNote = value
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
	if this.Global.LineOfField <= 0 {
		return fmt.Errorf("config check failed: lineOfField <= 0")
	} // if

	if this.Global.LineOfLayer <= 0 {
		return fmt.Errorf("config check failed: lineOfLayer <= 0")
	} // if

	if this.Global.LineOfNote <= 0 {
		return fmt.Errorf("config check failed: lineOfNote <= 0")
	} // if

	if this.Global.LineOfData <= 0 {
		return fmt.Errorf("config check failed: lineOfData <= 0")
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfField(%d) >= lineOfData(%d)", this.Global.LineOfField, this.Global.LineOfData)
	} // if

	if this.Global.LineOfLayer >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfLayer(%d) >= lineOfData(%d)", this.Global.LineOfLayer, this.Global.LineOfData)
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return fmt.Errorf("config check failed: lineOfNote(%d) >= lineOfData(%d)", this.Global.LineOfNote, this.Global.LineOfData)
	} // if

	for _, itor := range this.Elements {
		if utils.NameCheck(utils.FileName(itor.Excel)) == false {
			return fmt.Errorf("config check failed: invalid excel name")
		} // if

		if utils.NameCheck(itor.Sheet) == false {
			return fmt.Errorf("config check failed: invalid sheet name")
		} // if
	} // for

	return nil
}
