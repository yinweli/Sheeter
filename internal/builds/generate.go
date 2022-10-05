package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Generate 產生程式碼
func Generate(runtime *Runtime, config *Config) (errs []error) {
	tasks := []func(*RuntimeStruct) error{}

	if config.Global.GenerateJson {
		tasks = append(tasks, generateJsonCsStruct, generateJsonCsReader, generateJsonGoStruct, generateJsonGoReader)
	} // if

	if config.Global.GenerateProto {
		tasks = append(tasks, generateProtoSchema, generateProtoCsReader, generateProtoGoReader)
	} // if

	itemCount := len(runtime.Struct)
	taskCount := len(tasks)
	totalCount := itemCount * taskCount
	errors := make(chan error, itemCount) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := utils.NewWaitGroup(itemCount)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(signaler))
	progressbar := progress.AddBar(
		int64(totalCount),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("generate "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	for _, itor := range runtime.Struct {
		runtimeStruct := itor // 多執行緒需要使用中間變數

		go func() {
			defer signaler.Done()

			for _, itor := range tasks {
				if err := itor(runtimeStruct); err != nil {
					errors <- fmt.Errorf("generate failed: %w", err)
				} // if

				progressbar.Increment()
			} // for
		}()
	} // for

	progress.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
