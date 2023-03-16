package builds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Config 設定資料
type Config struct {
	Global Global   `yaml:"global"` // 全域設定
	Source []string `yaml:"source"` // 來源列表
	Output string   `yaml:"output"` // 輸出路徑
}

// Global 全域設定
type Global struct {
	Tag         string `yaml:"tag"`         // 標籤列表
	AutoKey     bool   `yaml:"autoKey"`     // 自動選取索引
	LineOfTag   int    `yaml:"lineOfTag"`   // 標籤行號(1為起始行)
	LineOfName  int    `yaml:"lineOfName"`  // 名稱行號(1為起始行)
	LineOfNote  int    `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfField int    `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfData  int    `yaml:"lineOfData"`  // 資料行號(1為起始行)
}

// Initialize 初始化設定
func (this *Config) Initialize(cmd *cobra.Command) error {
	flag := cmd.Flags()

	if flag.Changed(flagConfig) {
		if value, err := flag.GetString(flagConfig); err == nil {
			file, err := os.ReadFile(value)

			if err != nil {
				return fmt.Errorf("config initialize: %w", err)
			} // if

			if err = yaml.Unmarshal(file, this); err != nil {
				return fmt.Errorf("config initialize: %w", err)
			} // if
		} // if
	} // if

	if flag.Changed(flagSource) {
		if value, err := flag.GetStringSlice(flagSource); err == nil {
			this.Source = append(this.Source, value...)
		} // if
	} // if

	if flag.Changed(flagOutput) {
		if value, err := flag.GetString(flagOutput); err == nil {
			this.Output = value
		} // if
	} // if

	if flag.Changed(flagTag) {
		if value, err := flag.GetString(flagTag); err == nil {
			this.Global.Tag = value
		} // if
	} // if

	if flag.Changed(flagAutoKey) {
		if value, err := flag.GetBool(flagAutoKey); err == nil {
			this.Global.AutoKey = value
		} // if
	} // if

	if flag.Changed(flagLineOfTag) {
		if value, err := flag.GetInt(flagLineOfTag); err == nil {
			this.Global.LineOfTag = value
		} // if
	} // if

	if flag.Changed(flagLineOfName) {
		if value, err := flag.GetInt(flagLineOfName); err == nil {
			this.Global.LineOfName = value
		} // if
	} // if

	if flag.Changed(flagLineOfNote) {
		if value, err := flag.GetInt(flagLineOfNote); err == nil {
			this.Global.LineOfNote = value
		} // if
	} // if

	if flag.Changed(flagLineOfField) {
		if value, err := flag.GetInt(flagLineOfField); err == nil {
			this.Global.LineOfField = value
		} // if
	} // if

	if flag.Changed(flagLineOfData) {
		if value, err := flag.GetInt(flagLineOfData); err == nil {
			this.Global.LineOfData = value
		} // if
	} // if

	return nil
}

// File 從來源列表取得檔案列表
func (this *Config) File() []string {
	result := []string{}

	for _, itor := range this.Source {
		info, err := os.Stat(itor)

		if err != nil {
			continue
		} // if

		if info.IsDir() {
			continue
		} // if

		if filepath.Ext(itor) != sheeter.ExcelExt {
			continue
		} // if

		result = append(result, itor)
	} // for

	return utils.GetUnique(result)
}

// Path 從來源列表取得路徑列表
func (this *Config) Path() []string {
	result := []string{}

	for _, itor := range this.Source {
		info, err := os.Stat(itor)

		if err != nil {
			continue
		} // if

		if info.IsDir() == false {
			continue
		} // if

		result = append(result, itor)
	} // for

	return utils.GetUnique(result)
}

// Check 檢查設定
func (this *Config) Check() error {
	if this.Global.LineOfTag <= 0 {
		return fmt.Errorf("config check: lineOfTag <= 0")
	} // if

	if this.Global.LineOfName <= 0 {
		return fmt.Errorf("config check: lineOfName <= 0")
	} // if

	if this.Global.LineOfNote <= 0 {
		return fmt.Errorf("config check: lineOfNote <= 0")
	} // if

	if this.Global.LineOfField <= 0 {
		return fmt.Errorf("config check: lineOfField <= 0")
	} // if

	if this.Global.LineOfData <= 0 {
		return fmt.Errorf("config check: lineOfData <= 0")
	} // if

	if this.Global.LineOfTag >= this.Global.LineOfData {
		return fmt.Errorf("config check: lineOfTag(%d) >= lineOfData(%d)", this.Global.LineOfTag, this.Global.LineOfData)
	} // if

	if this.Global.LineOfName >= this.Global.LineOfData {
		return fmt.Errorf("config check: lineOfName(%d) >= lineOfData(%d)", this.Global.LineOfName, this.Global.LineOfData)
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return fmt.Errorf("config check: lineOfNote(%d) >= lineOfData(%d)", this.Global.LineOfNote, this.Global.LineOfData)
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return fmt.Errorf("config check: lineOfField(%d) >= lineOfData(%d)", this.Global.LineOfField, this.Global.LineOfData)
	} // if

	return nil
}
