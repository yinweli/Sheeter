package build

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"Sheeter/internal"
	"Sheeter/internal/command/build/core"

	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		Run:   execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	startTime := time.Now()
	err := exec.Command("go", "version").Run() // 檢查是否有安裝go

	if err != nil {
		cmd.Println(err)
		return
	} // if

	err = exec.Command("protoc").Run() // 檢查是否有安裝protoc

	if err != nil {
		cmd.Println(err)
		return
	} // if

	config, err := core.ReadConfig(args[0])

	if err != nil {
		cmd.Println(err)
		return
	} // if

	count := len(config.Elements)
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(40), mpb.WithWaitGroup(&signaler))
	messenger := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來

	signaler.Add(count)

	for _, itor := range config.Elements {
		global := config.Global
		element := itor
		name := fmt.Sprintf("%s(%s)", element.Excel, element.Sheet)
		bar := progress.AddBar(
			core.TaskProgress,
			mpb.PrependDecorators(
				decor.Name(fmt.Sprintf("%-20s", name)),
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				decor.OnComplete(decor.Spinner(nil), "complete"),
			),
		)

		go func() {
			defer signaler.Done()
			messenger <- core.NewTask(&global, &element, bar).Execute()
		}()
	} // for

	signaler.Wait()
	close(messenger) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range messenger {
		if itor != nil {
			cmd.Println(itor)
		} // if
	} // for

	cmd.Printf("%s done, usage time=%s\n", internal.Title, time.Since(startTime))
}
