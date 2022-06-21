package build

import (
	"fmt"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/command/build/core"
	"github.com/yinweli/Sheeter/internal/util"
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

	check := true
	check = installCheck(cmd, "go") && check        // 檢查是否有安裝go
	check = installCheck(cmd, "quicktype") && check // 檢查是否有安裝quicktype

	if check == false {
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

		go func() {
			defer signaler.Done()
			messenger <- core.NewTask(&global, &element).Execute(progress)
		}()
	} // for

	signaler.Wait()
	close(messenger) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range messenger {
		if itor != nil {
			cmd.Println(itor)
		} // if
	} // for

	cmd.Printf("%s done, usage time=%s\n", internal.Title, durafmt.Parse(time.Since(startTime)))
}

// installCheck 檢查是否有安裝軟體
func installCheck(cmd *cobra.Command, name string) bool {
	err := util.ShellExist(name)

	if err != nil {
		cmd.Println(fmt.Errorf("`%s` not installed\n%s", name, err))
		return false
	} // if

	return true
}
