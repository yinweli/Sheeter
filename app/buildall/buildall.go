package buildall

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/build/tasks"
	"github.com/yinweli/Sheeter/internal/build/thirdparty"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
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
	err := thirdparty.Check()

	if err != nil {
		cmd.Println(err)
		return
	} // if

	config, err := readConfig(args[0])

	if err != nil {
		cmd.Println(err)
		return
	} // if

	count := len(config.Elements)
	signaler := sync.WaitGroup{}
	errors := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	progress := mpb.New(mpb.WithWidth(barWidth), mpb.WithWaitGroup(&signaler))

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

// readConfig 讀取設定 TODO: 考慮合併進去buildall
func readConfig(fileName string) (result *Config, err error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	} // if

	result = &Config{}

	if err = yaml.Unmarshal(bytes, result); err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	} // if

	if err = result.Global.Check(); err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	} // if

	for _, itor := range result.Elements {
		if err = itor.Check(); err != nil {
			return nil, fmt.Errorf("read config failed: %w", err)
		} // if
	} // for

	return result, nil
}

// Config 編譯設定
type Config struct {
	Global   tasks.Global    `yaml:"global"`   // 全域設定
	Elements []tasks.Element `yaml:"elements"` // 項目設定列表
}
