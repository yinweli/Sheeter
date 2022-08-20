package buildall

import (
	"fmt"
	"os"
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
const flagExcels = "excels"           // 旗標名稱: excel檔案名稱列表
const flagSheets = "sheets"           // 旗標名稱: excel表單名稱列表

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buildall [config]",
		Short: "build all sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			duration, errs := execute(args[0])

			for _, itor := range errs {
				cmd.Println(itor)
			} // for

			cmd.Printf("usage time=%s\n", durafmt.Parse(duration))
		},
	}
	cmd.Flags().String(flagConfig, "", "config filepath")
	cmd.Flags().String(flagBom, "", "bom")
	cmd.Flags().String(flagLineOfField, "", "line of field")
	cmd.Flags().String(flagLineOfLayer, "", "line of layer")
	cmd.Flags().String(flagLineOfNote, "", "line of note")
	cmd.Flags().String(flagLineOfData, "", "line of data")
	cmd.Flags().String(flagExcels, "", "excel lists")
	cmd.Flags().String(flagSheets, "", "sheet lists")
	return cmd
}

// execute 執行命令
func execute(fileName string) (duration time.Duration, errs []error) {
	startTime := time.Now()

	if util.ShellExist("go") == false {
		return time.Since(startTime), []error{fmt.Errorf("build all failed, `go` not installed")}
	} // if

	if util.ShellExist("quicktype") == false {
		return time.Since(startTime), []error{fmt.Errorf("build all failed, `quicktype` not installed")}
	} // if

	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return time.Since(startTime), []error{fmt.Errorf("build all failed, read config failed: %w", err)}
	} // if

	config := &config{}

	if err = yaml.Unmarshal(bytes, config); err != nil {
		return time.Since(startTime), []error{fmt.Errorf("build all failed, read config failed: %w", err)}
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
			errs = append(errs, itor)
		} // if
	} // for

	return time.Since(startTime), errs
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
