package buildall

import (
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/build/tasks"
	"github.com/yinweli/Sheeter/internal/build/thirdParty"
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

	config, err := tasks.ReadConfig(args[0])

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
