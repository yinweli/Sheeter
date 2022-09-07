package builds

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const separateElement = ":" // 項目字串以':'符號分割為檔案名稱與表單名稱

// NewConfig 建立編譯設定
func NewConfig(cmd *cobra.Command) (config *Config, err error) {
	flags := cmd.Flags()
	config = &Config{}

	if flags.Changed(flagConfig) {
		if filepath, err := flags.GetString(flagConfig); err == nil {
			bytes, err := os.ReadFile(filepath)

			if err != nil {
				return nil, fmt.Errorf("new config failed, read config failed: %w", err)
			} // if

			if err = yaml.Unmarshal(bytes, config); err != nil {
				return nil, fmt.Errorf("new config failed, read config failed: %w", err)
			} // if
		} // if
	} // if

	if flags.Changed(flagBom) {
		if bom, err := flags.GetBool(flagBom); err == nil {
			config.Global.Bom = bom
		} // if
	} // if

	if flags.Changed(flagLineOfField) {
		if lineOfField, err := flags.GetInt(flagLineOfField); err == nil {
			config.Global.LineOfField = lineOfField
		} // if
	} // if

	if flags.Changed(flagLineOfLayer) {
		if lineOfLayer, err := flags.GetInt(flagLineOfLayer); err == nil {
			config.Global.LineOfLayer = lineOfLayer
		} // if
	} // if

	if flags.Changed(flagLineOfNote) {
		if lineOfNote, err := flags.GetInt(flagLineOfNote); err == nil {
			config.Global.LineOfNote = lineOfNote
		} // if
	} // if

	if flags.Changed(flagLineOfData) {
		if lineOfData, err := flags.GetInt(flagLineOfData); err == nil {
			config.Global.LineOfData = lineOfData
		} // if
	} // if

	if flags.Changed(flagElements) {
		if elements, err := flags.GetStringSlice(flagElements); err == nil {
			for _, itor := range elements {
				if before, after, ok := strings.Cut(itor, separateElement); ok {
					config.Elements = append(config.Elements, Element{
						Excel: before,
						Sheet: after,
					})
				} // if
			} // for
		} // if
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
