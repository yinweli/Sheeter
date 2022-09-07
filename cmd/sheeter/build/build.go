package build

import (
	"fmt"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal/builds"
	"github.com/yinweli/Sheeter/internal/util"
)

const barWidth = 40  // 進度條寬度
const taskItem = 3   // 單項編譯的工作數量
const taskEntire = 4 // 全域編譯的工作數量

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	return builds.InitFlags(&cobra.Command{
		Use:   "build",
		Short: "build sheet",
		Long:  "build sheet",
		Run:   execute,
	})
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	startTime := time.Now()

	if util.ShellExist("gofmt") == false {
		cmd.Println(fmt.Errorf("build failed, `gofmt` not installed"))
		return
	} // if

	if util.ShellExist("quicktype") == false {
		cmd.Println(fmt.Errorf("build failed, `quicktype` not installed"))
		return
	} // if

	config, err := builds.NewConfig(cmd)

	if err != nil {
		cmd.Println(fmt.Errorf("build failed, config read failed: %w", err))
		return
	} // if

	if err := config.Check(); err != nil {
		cmd.Println(fmt.Errorf("build failed, config check failed: %w", err))
		return
	} // if

	contents := config.ToContents()

	if buildItem(cmd, contents) == false {
		cmd.Println(fmt.Errorf("build failed, build item failed"))
		return
	} // if

	if buildEntire(cmd, contents, config.Global.Bom) == false {
		cmd.Println(fmt.Errorf("build failed, build entire failed"))
		return
	} // if

	cmd.Printf("usage time=%s\n", durafmt.Parse(time.Since(startTime)))
}

// buildItem 單項編譯
func buildItem(cmd *cobra.Command, contents *builds.Contents) bool {
	count := len(contents.Contents)
	errors := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(barWidth), mpb.WithWaitGroup(&signaler))

	signaler.Add(count)

	for _, itor := range contents.Contents {
		content := itor // 多執行緒需要中間變數

		go func() {
			defer signaler.Done()
			defer content.Close()

			bar := progress.AddBar(
				taskItem,
				mpb.PrependDecorators(
					decor.Name(fmt.Sprintf("%-20s", content.StructName())),
					decor.Percentage(decor.WCSyncSpace),
				),
				mpb.AppendDecorators(
					decor.OnComplete(decor.Spinner(nil), "complete"),
				),
			)

			if err := builds.Initialize(content); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := builds.OutputJson(content); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := builds.OutputJsonSchema(content); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()
		}()
	} // for

	progress.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			cmd.Println(itor)
		} // if
	} // for

	return len(errors) == 0
}

// buildEntire 全域編譯
func buildEntire(cmd *cobra.Command, contents *builds.Contents, bom bool) bool {
	progress := mpb.New(mpb.WithWidth(barWidth))
	bar := progress.AddBar(
		taskEntire,
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%-20s", "generate code")),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	if err := builds.OutputJsonCsCode(contents); err != nil {
		cmd.Println(err)
		return false
	} // if

	bar.Increment()

	if err := builds.OutputJsonCsReader(contents, bom); err != nil {
		cmd.Println(err)
		return false
	} // if

	bar.Increment()

	if err := builds.OutputJsonGoCode(contents); err != nil {
		cmd.Println(err)
		return false
	} // if

	bar.Increment()

	if err := builds.OutputJsonGoReader(contents, bom); err != nil {
		cmd.Println(err)
		return false
	} // if

	bar.Increment()
	progress.Wait()
	return true
}
