package builds

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const FlagConfig = "config"           // 旗標名稱: 編譯設定檔案路徑
const FlagBom = "bom"                 // 旗標名稱: 順序標記
const FlagLineOfField = "lineOfField" // 旗標名稱: 欄位行號
const FlagLineOfLayer = "lineOfLayer" // 旗標名稱: 階層行號
const FlagLineOfNote = "lineOfNote"   // 旗標名稱: 註解行號
const FlagLineOfData = "lineOfData"   // 旗標名稱: 資料行號
const FlagElements = "elements"       // 旗標名稱: 項目列表
const separateElement = ":"           // 項目字串以':'符號分割為檔案名稱與表單名稱

// NewConfig 建立編譯設定
func NewConfig(cmd *cobra.Command) (config *Config, err error) {
	config = &Config{}

	if filepath, err := cmd.Flags().GetString(FlagConfig); err == nil {
		bytes, err := os.ReadFile(filepath)

		if err != nil {
			return nil, fmt.Errorf("new config failed, read config failed: %w", err)
		} // if

		if err = yaml.Unmarshal(bytes, config); err != nil {
			return nil, fmt.Errorf("new config failed, read config failed: %w", err)
		} // if
	} // if

	if bom, err := cmd.Flags().GetBool(FlagBom); err == nil {
		config.Global.Bom = bom
	} // if

	if lineOfField, err := cmd.Flags().GetInt(FlagLineOfField); err == nil {
		config.Global.LineOfField = lineOfField
	} // if

	if lineOfLayer, err := cmd.Flags().GetInt(FlagLineOfLayer); err == nil {
		config.Global.LineOfLayer = lineOfLayer
	} // if

	if lineOfNote, err := cmd.Flags().GetInt(FlagLineOfNote); err == nil {
		config.Global.LineOfNote = lineOfNote
	} // if

	if lineOfData, err := cmd.Flags().GetInt(FlagLineOfData); err == nil {
		config.Global.LineOfData = lineOfData
	} // if

	if elements, err := cmd.Flags().GetStringSlice(FlagElements); err == nil {
		for _, itor := range elements {
			if before, after, ok := strings.Cut(itor, separateElement); ok {
				config.Elements = append(config.Elements, Element{
					Excel: before,
					Sheet: after,
				})
			} // if
		} // for
	} // if

	return config, nil
}

// Config 編譯設定
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
		if itor.Excel == "" {
			return fmt.Errorf("config check failed, excel empty")
		} // if

		if itor.Sheet == "" {
			return fmt.Errorf("config check failed, sheet empty")
		} // if
	} // for

	return nil
}

// ToContents 轉換成內容列表
func (this *Config) ToContents() *Contents {
	contents := &Contents{}

	for _, itor := range this.Elements {
		contents.Contents = append(contents.Contents, &Content{
			Bom:         this.Global.Bom,
			LineOfField: this.Global.LineOfField,
			LineOfLayer: this.Global.LineOfLayer,
			LineOfNote:  this.Global.LineOfNote,
			LineOfData:  this.Global.LineOfData,
			Excel:       itor.Excel,
			Sheet:       itor.Sheet,
		})
	} // for

	// 幫內容列表排序, 可以保證不管外面的輸入是否有序, 輸出時仍然有序
	// 如此可以保證輸出的讀取器內容為有序的, 對於使用版本控制的專案, 會有幫助
	sort.Slice(contents.Contents, func(i, j int) bool {
		return contents.Contents[i].StructName() < contents.Contents[j].StructName()
	})

	return contents
}
