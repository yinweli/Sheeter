package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Initialize 初始化
func Initialize(config *Config, runtime *Runtime) (errs []error) {
	for _, itor := range config.Elements {
		runtime.Sector = append(runtime.Sector, &RuntimeSector{
			Global:  config.Global,
			Element: itor,
		})
	} // for

	tasks := []func(*RuntimeSector) error{ // 工作函式列表
		initializeSector,
	}
	itemCount := len(runtime.Sector)
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
			decor.Name("initialize "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	for _, itor := range runtime.Sector {
		runtimeSector := itor // 多執行緒需要使用中間變數

		go func() {
			defer signaler.Done()

			for _, itor := range tasks {
				if err := itor(runtimeSector); err != nil {
					errors <- fmt.Errorf("initialize failed: %w", err)
				} // if

				progressbar.Increment()
			} // for
		}()
	} // for

	progress.Wait()

	if err := initializeStruct(runtime); err != nil {
		errors <- err
	} // if

	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
