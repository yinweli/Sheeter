package buildall

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/build/tasks"
	"github.com/yinweli/Sheeter/internal/util"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buildall [config]",
		Short: "build all sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		Run:   execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	startTime := time.Now()

	if util.ShellExist("go") == false {
		cmd.Println(fmt.Errorf("`go` not installed"))
		return
	} // if

	if util.ShellExist("quicktype") == false {
		cmd.Println(fmt.Errorf("`quicktype` not installed"))
		return
	} // if

	bytes, err := os.ReadFile(args[0])

	if err != nil {
		cmd.Println(fmt.Errorf("read config failed: %w", err))
		return
	} // if

	config := &Config{}

	if err = yaml.Unmarshal(bytes, config); err != nil {
		cmd.Println(fmt.Errorf("read config failed: %w", err))
		return
	} // if

	count := len(config.Elements)
	errors := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := util.NewMpb(&signaler)

	signaler.Add(count)

	for _, itor := range config.Elements {
		global := config.Global
		element := itor

		go func() {
			defer signaler.Done()
			task := tasks.Task{
				Path:        global.Path,
				Bom:         global.Bom,
				LineOfField: global.LineOfField,
				LineOfLayer: global.LineOfLayer,
				LineOfNote:  global.LineOfNote,
				LineOfData:  global.LineOfData,
				Excel:       element.Excel,
				Sheet:       element.Sheet,
				Progress:    progress,
			}
			errors <- task.Run()
		}()
	} // for

	signaler.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			cmd.Println(itor)
		} // if
	} // for

	cmd.Printf("%s done, usage time=%s\n", internal.Title, durafmt.Parse(time.Since(startTime)))
}

// Config 編譯設定
type Config struct {
	Global   Global    `yaml:"global"`   // 全域設定
	Elements []Element `yaml:"elements"` // 項目設定列表
}

// Global 全域設定
type Global struct {
	Path        string `yaml:"excelPath"`   // 來源excel路徑 TODO: yaml:"excelPath"改名為yaml:"path"
	Bom         bool   `yaml:"bom"`         // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int    `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfLayer int    `yaml:"lineOfLayer"` // 階層行號(1為起始行)
	LineOfNote  int    `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfData  int    `yaml:"lineOfData"`  // 資料起始行號(1為起始行)
}

// Element 項目設定
type Element struct {
	Excel string `yaml:"excel"` // excel檔案名稱
	Sheet string `yaml:"sheet"` // excel表單名稱
}
