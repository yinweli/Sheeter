package build

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yinweli/Sheeter/internal/builds"
	"github.com/yinweli/Sheeter/internal/util"
)

const flagConfig = "config"           // 旗標名稱: 編譯設定檔案路徑
const flagBom = "bom"                 // 旗標名稱: 順序標記
const flagLineOfField = "lineOfField" // 旗標名稱: 欄位行號
const flagLineOfLayer = "lineOfLayer" // 旗標名稱: 階層行號
const flagLineOfNote = "lineOfNote"   // 旗標名稱: 註解行號
const flagLineOfData = "lineOfData"   // 旗標名稱: 資料行號
const flagElements = "elements"       // 旗標名稱: 項目列表
const separateElement = ":"           // 項目字串以':'符號分割為檔案名稱與表單名稱

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [-c config file path | -b bom | -lf line of field | -ll line of layer | -ln line of note | -ld line of data | -e element lists(excel:sheet)]",
		Short: "build sheet",
		Long:  "build sheet",
		Run:   execute,
	}
	cmd.Flags().StringP(flagConfig, "c", "", "config file path")
	cmd.Flags().StringP(flagBom, "b", "", "bom")
	cmd.Flags().StringP(flagLineOfField, "lf", "", "line of field")
	cmd.Flags().StringP(flagLineOfLayer, "ll", "", "line of layer")
	cmd.Flags().StringP(flagLineOfNote, "ln", "", "line of note")
	cmd.Flags().StringP(flagLineOfData, "ld", "", "line of data")
	cmd.Flags().StringP(flagElements, "e", "", "element lists(excel:sheet)")
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	startTime := time.Now()

	if util.ShellExist("go") == false {
		cmd.Println(fmt.Errorf("build failed, `go` not installed"))
		return
	} // if

	if util.ShellExist("quicktype") == false {
		cmd.Println(fmt.Errorf("build failed, `quicktype` not installed"))
		return
	} // if

	config := &config{}

	if filepath, err := cmd.Flags().GetString(flagConfig); err == nil {
		bytes, err := os.ReadFile(filepath)

		if err != nil {
			cmd.Println(fmt.Errorf("build failed, read config failed: %w", err))
			return
		} // if

		if err = yaml.Unmarshal(bytes, config); err != nil {
			cmd.Println(fmt.Errorf("build failed, read config failed: %w", err))
			return
		} // if
	} // if

	if bom, err := cmd.Flags().GetBool(flagBom); err == nil {
		config.Global.Bom = bom
	} // if

	if lineOfField, err := cmd.Flags().GetInt(flagLineOfField); err == nil {
		config.Global.LineOfField = lineOfField
	} // if

	if lineOfLayer, err := cmd.Flags().GetInt(flagLineOfLayer); err == nil {
		config.Global.LineOfLayer = lineOfLayer
	} // if

	if lineOfNote, err := cmd.Flags().GetInt(flagLineOfNote); err == nil {
		config.Global.LineOfNote = lineOfNote
	} // if

	if lineOfData, err := cmd.Flags().GetInt(flagLineOfData); err == nil {
		config.Global.LineOfData = lineOfData
	} // if

	if elements, err := cmd.Flags().GetStringSlice(flagElements); err == nil {
		for _, itor := range elements {
			if before, after, ok := strings.Cut(itor, separateElement); ok {
				config.Elements = append(config.Elements, element{
					Excel: before,
					Sheet: after,
				})
			} // if
		} // for
	} // if

	count := len(config.Elements)
	errors := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := util.NewMpb(&signaler)

	signaler.Add(count)

	for _, itor := range config.Elements {
		global := config.Global
		element := itor // 由於多執行緒的關係, 所以要創建中間變數會比較安全

		go func() {
			defer signaler.Done()
			content := &builds.Content{
				Bom:         global.Bom,
				LineOfField: global.LineOfField,
				LineOfLayer: global.LineOfLayer,
				LineOfNote:  global.LineOfNote,
				LineOfData:  global.LineOfData,
				Excel:       element.Excel,
				Sheet:       element.Sheet,
				Progress:    progress,
			}
			errors <- builds.Build(content)
		}()
	} // for

	signaler.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			cmd.Println(itor)
		} // if
	} // for

	cmd.Printf("usage time=%s\n", durafmt.Parse(time.Since(startTime)))
}

// config 編譯設定
type config struct {
	Global   global    `yaml:"global"`   // 全域設定
	Elements []element `yaml:"elements"` // 項目設定列表
}

// global 全域設定
type global struct {
	Bom         bool `yaml:"bom"`         // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int  `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfLayer int  `yaml:"lineOfLayer"` // 階層行號(1為起始行)
	LineOfNote  int  `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfData  int  `yaml:"lineOfData"`  // 資料行號(1為起始行)
}

// element 項目設定
type element struct {
	Excel string `yaml:"excel"` // excel檔案名稱
	Sheet string `yaml:"sheet"` // excel表單名稱
}
