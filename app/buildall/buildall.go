package buildall

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/build/tasks"
	"github.com/yinweli/Sheeter/internal/build/thirdParty"
	"gopkg.in/yaml.v3"
)

const barWidth = 40 // 進度條寬度

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

	if thirdParty.Check(cmd.Println) == false {
		return
	} // if

	config, err := ReadConfig(args[0])

	if err != nil {
		cmd.Println(err)
		return
	} // if

	count := len(config.Elements)
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(barWidth), mpb.WithWaitGroup(&signaler))
	errors := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來

	signaler.Add(count)

	for _, itor := range config.Elements {
		global := config.Global
		element := itor

		go func() {
			defer signaler.Done()
			errors <- tasks.NewTask(&global, &element).Run(progress)
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

// ReadConfig 讀取設定
func ReadConfig(fileName string) (result *Config, err error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	} // if

	result = &Config{}
	err = yaml.Unmarshal(bytes, result)

	if err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	} // if

	err = result.Check()

	if err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	} // if

	return result, nil
}

// Config 編譯設定
type Config struct {
	Global   tasks.Global    `yaml:"global"`   // 全域設定
	Elements []tasks.Element `yaml:"elements"` // 項目設定列表
}

// Check 檢查設定是否正確
func (this *Config) Check() error {
	if this.Global.LineOfField <= 0 {
		return fmt.Errorf("global: LineOfField <= 0")
	} // if

	if this.Global.LineOfLayer <= 0 {
		return fmt.Errorf("global: LineOfLayer <= 0")
	} // if

	if this.Global.LineOfNote <= 0 {
		return fmt.Errorf("global: LineOfNote <= 0")
	} // if

	if this.Global.LineOfData <= 0 {
		return fmt.Errorf("global: LineOfData <= 0")
	} // if

	if this.Global.LineOfField >= this.Global.LineOfData {
		return fmt.Errorf("global: LineOfField(%d) >= LineOfData(%d)", this.Global.LineOfField, this.Global.LineOfData)
	} // if

	if this.Global.LineOfLayer >= this.Global.LineOfData {
		return fmt.Errorf("global: LineOfLayer(%d) >= LineOfData(%d)", this.Global.LineOfLayer, this.Global.LineOfData)
	} // if

	if this.Global.LineOfNote >= this.Global.LineOfData {
		return fmt.Errorf("global: LineOfNote(%d) >= LineOfData(%d)", this.Global.LineOfNote, this.Global.LineOfData)
	} // if

	for _, itor := range this.Elements {
		if itor.Excel == "" {
			return fmt.Errorf("element: excel empty")
		} // if

		if itor.Sheet == "" {
			return fmt.Errorf("element: sheet empty")
		} // if
	} // for

	return nil
}
