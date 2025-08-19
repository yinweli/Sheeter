package builds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

const (
	flagConfig      = "config"      // 旗標名稱: 設定檔案路徑
	flagSource      = "source"      // 旗標名稱: 來源列表
	flagMerge       = "merge"       // 旗標名稱: 合併列表
	flagExclude     = "exclude"     // 旗標名稱: 排除列表
	flagOutput      = "output"      // 旗標名稱: 輸出路徑
	flagTag         = "tag"         // 旗標名稱: 標籤列表
	flagLineOfTag   = "lineOfTag"   // 旗標名稱: 標籤行號
	flagLineOfName  = "lineOfName"  // 旗標名稱: 名稱行號
	flagLineOfNote  = "lineOfNote"  // 旗標名稱: 註解行號
	flagLineOfField = "lineOfField" // 旗標名稱: 欄位行號
	flagLineOfData  = "lineOfData"  // 旗標名稱: 資料行號
)

// Config 設定資料
type Config struct {
	Source      []string `yaml:"source"`      // 來源列表
	Merge       []string `yaml:"merge"`       // 合併列表
	Exclude     []string `yaml:"exclude"`     // 排除列表
	Output      string   `yaml:"output"`      // 輸出路徑
	Tag         string   `yaml:"tag"`         // 標籤列表
	LineOfTag   int      `yaml:"lineOfTag"`   // 標籤行號(1為起始行)
	LineOfName  int      `yaml:"lineOfName"`  // 名稱行號(1為起始行)
	LineOfNote  int      `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfField int      `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfData  int      `yaml:"lineOfData"`  // 資料行號(1為起始行)
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

	if flag.Changed(flagMerge) {
		if value, err := flag.GetStringSlice(flagMerge); err == nil {
			this.Merge = append(this.Merge, value...)
		} // if
	} // if

	if flag.Changed(flagExclude) {
		if value, err := flag.GetStringSlice(flagExclude); err == nil {
			this.Exclude = append(this.Exclude, value...)
		} // if
	} // if

	if flag.Changed(flagOutput) {
		if value, err := flag.GetString(flagOutput); err == nil {
			this.Output = value
		} // if
	} // if

	if flag.Changed(flagTag) {
		if value, err := flag.GetString(flagTag); err == nil {
			this.Tag = value
		} // if
	} // if

	if flag.Changed(flagLineOfTag) {
		if value, err := flag.GetInt(flagLineOfTag); err == nil {
			this.LineOfTag = value
		} // if
	} // if

	if flag.Changed(flagLineOfName) {
		if value, err := flag.GetInt(flagLineOfName); err == nil {
			this.LineOfName = value
		} // if
	} // if

	if flag.Changed(flagLineOfNote) {
		if value, err := flag.GetInt(flagLineOfNote); err == nil {
			this.LineOfNote = value
		} // if
	} // if

	if flag.Changed(flagLineOfField) {
		if value, err := flag.GetInt(flagLineOfField); err == nil {
			this.LineOfField = value
		} // if
	} // if

	if flag.Changed(flagLineOfData) {
		if value, err := flag.GetInt(flagLineOfData); err == nil {
			this.LineOfData = value
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

		if filepath.Ext(itor) != sheeter.ExtExcel {
			continue
		} // if

		result = append(result, itor)
	} // for

	return utils.Unique(result)
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

	return utils.Unique(result)
}

// Merged 從合併列表取得合併資料
func (this *Config) Merged() []utils.MergeTerm {
	result := []utils.MergeTerm{}

	for _, itor := range this.Merge {
		result = append(result, utils.MergeTerm(itor))
	} // for

	return result
}

// Excluded 檢查是否排除
func (this *Config) Excluded(excel, sheet string) bool {
	for _, itor := range this.Exclude {
		if utils.SheetTerm(itor).Match(excel, sheet) {
			return true
		} // if
	} // for

	return false
}

// Check 檢查設定
func (this *Config) Check() error {
	if this.LineOfTag <= 0 {
		return fmt.Errorf("config check: lineOfTag <= 0")
	} // if

	if this.LineOfName <= 0 {
		return fmt.Errorf("config check: lineOfName <= 0")
	} // if

	if this.LineOfNote <= 0 {
		return fmt.Errorf("config check: lineOfNote <= 0")
	} // if

	if this.LineOfField <= 0 {
		return fmt.Errorf("config check: lineOfField <= 0")
	} // if

	if this.LineOfData <= 0 {
		return fmt.Errorf("config check: lineOfData <= 0")
	} // if

	if this.LineOfTag >= this.LineOfData {
		return fmt.Errorf("config check: lineOfTag(%d) >= lineOfData(%d)", this.LineOfTag, this.LineOfData)
	} // if

	if this.LineOfName >= this.LineOfData {
		return fmt.Errorf("config check: lineOfName(%d) >= lineOfData(%d)", this.LineOfName, this.LineOfData)
	} // if

	if this.LineOfNote >= this.LineOfData {
		return fmt.Errorf("config check: lineOfNote(%d) >= lineOfData(%d)", this.LineOfNote, this.LineOfData)
	} // if

	if this.LineOfField >= this.LineOfData {
		return fmt.Errorf("config check: lineOfField(%d) >= lineOfData(%d)", this.LineOfField, this.LineOfData)
	} // if

	return nil
}

// SetFlag 設定命令旗標
func SetFlag(cmd *cobra.Command) *cobra.Command {
	flag := cmd.Flags()
	flag.String(flagConfig, "", "config file path")
	flag.StringSlice(flagSource, []string{}, "source file/folder list")
	flag.StringSlice(flagMerge, []string{}, "merge list, example: name$excel#sheet&...,...")
	flag.StringSlice(flagExclude, []string{}, "exclude list, example: excel#sheet,...")
	flag.String(flagOutput, "", "output path")
	flag.String(flagTag, "", "tag that determines the columns to be output")
	flag.Int(flagLineOfTag, 0, "line of tag")
	flag.Int(flagLineOfName, 0, "line of name")
	flag.Int(flagLineOfNote, 0, "line of note")
	flag.Int(flagLineOfField, 0, "line of field")
	flag.Int(flagLineOfData, 0, "line of data")
	return cmd
}
