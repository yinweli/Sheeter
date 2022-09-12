package builds

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/internal/util"
)

const flagConfig = "config"           // 旗標名稱: 設定檔案路徑
const flagBom = "bom"                 // 旗標名稱: 順序標記
const flagLineOfField = "lineOfField" // 旗標名稱: 欄位行號
const flagLineOfLayer = "lineOfLayer" // 旗標名稱: 階層行號
const flagLineOfNote = "lineOfNote"   // 旗標名稱: 註解行號
const flagLineOfData = "lineOfData"   // 旗標名稱: 資料行號
const flagElements = "elements"       // 旗標名稱: 項目列表
const separateElement = ":"           // 項目字串以':'符號分割為檔案名稱與表單名稱

// InitializeFlags 初始化命令旗標
func InitializeFlags(cmd *cobra.Command) *cobra.Command {
	flags := cmd.Flags()
	flags.String(flagConfig, "", "config file path")
	flags.Bool(flagBom, false, "bom")
	flags.Int(flagLineOfField, 0, "line of field")
	flags.Int(flagLineOfLayer, 0, "line of layer")
	flags.Int(flagLineOfNote, 0, "line of note")
	flags.Int(flagLineOfData, 0, "line of data")
	flags.StringSlice(flagElements, []string{}, "element lists(excel:sheet,excel:sheet,excel:sheet,...)")
	return cmd
}

// Config 設定資料
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目設定列表
}

// Global 全域設定
type Global struct {
	Bom         bool `yaml:"bom"`         // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int  `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfLayer int  `yaml:"lineOfLayer"` // 階層行號(1為起始行)
	LineOfNote  int  `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfData  int  `yaml:"lineOfData"`  // 資料行號(1為起始行)
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
				return fmt.Errorf("new config failed, read config failed: %w", err)
			} // if

			if err = yaml.Unmarshal(datas, this); err != nil {
				return fmt.Errorf("new config failed, read config failed: %w", err)
			} // if
		} // if
	} // if

	if flags.Changed(flagBom) {
		if bom, err := flags.GetBool(flagBom); err == nil {
			this.Global.Bom = bom
		} // if
	} // if

	if flags.Changed(flagLineOfField) {
		if lineOfField, err := flags.GetInt(flagLineOfField); err == nil {
			this.Global.LineOfField = lineOfField
		} // if
	} // if

	if flags.Changed(flagLineOfLayer) {
		if lineOfLayer, err := flags.GetInt(flagLineOfLayer); err == nil {
			this.Global.LineOfLayer = lineOfLayer
		} // if
	} // if

	if flags.Changed(flagLineOfNote) {
		if lineOfNote, err := flags.GetInt(flagLineOfNote); err == nil {
			this.Global.LineOfNote = lineOfNote
		} // if
	} // if

	if flags.Changed(flagLineOfData) {
		if lineOfData, err := flags.GetInt(flagLineOfData); err == nil {
			this.Global.LineOfData = lineOfData
		} // if
	} // if

	if flags.Changed(flagElements) {
		if elements, err := flags.GetStringSlice(flagElements); err == nil {
			for _, itor := range elements {
				if before, after, ok := strings.Cut(itor, separateElement); ok {
					this.Elements = append(this.Elements, Element{
						Excel: before,
						Sheet: after,
					})
				} // if
			} // for
		} // if
	} // if

	// 幫項目列表排序, 可以保證輸出的讀取器內容為有序的, 對於使用版本控制的專案, 會有幫助
	sort.Slice(this.Elements, func(r, l int) bool {
		rightName := this.Elements[r].Excel + separateElement + this.Elements[r].Sheet
		leftName := this.Elements[l].Excel + separateElement + this.Elements[l].Sheet
		return rightName < leftName
	})

	return nil
}

// Check 檢查設定
func (this *Config) Check() error {
	if this.Global.LineOfField <= 0 {
		return fmt.Errorf("config check failed, lineOfField <= 0")
	} // if

	if this.Global.LineOfLayer <= 0 {
		return fmt.Errorf("config check failed, lineOfLayer <= 0")
	} // if

	if this.Global.LineOfNote <= 0 {
		return fmt.Errorf("config check failed, lineOfNote <= 0")
	} // if

	if this.Global.LineOfData <= 0 {
		return fmt.Errorf("config check failed, lineOfData <= 0")
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return fmt.Errorf("config check failed, lineOfField(%d) >= lineOfData(%d)", this.Global.LineOfField, this.Global.LineOfData)
	} // if

	if this.Global.LineOfLayer >= this.Global.LineOfData {
		return fmt.Errorf("config check failed, lineOfLayer(%d) >= lineOfData(%d)", this.Global.LineOfLayer, this.Global.LineOfData)
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return fmt.Errorf("config check failed, lineOfNote(%d) >= lineOfData(%d)", this.Global.LineOfNote, this.Global.LineOfData)
	} // if

	for _, itor := range this.Elements {
		if util.NameCheck(util.FileName(itor.Excel)) == false {
			return fmt.Errorf("config check failed, invalid excel name")
		} // if

		if util.NameCheck(itor.Sheet) == false {
			return fmt.Errorf("config check failed, invalid sheet name")
		} // if
	} // for

	return nil
}