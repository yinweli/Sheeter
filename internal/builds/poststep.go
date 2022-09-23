package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
)

// Poststep 後製步驟
func Poststep(runtime *Runtime) error {
	tasks := []func(*Runtime) error{ // 工作函式列表
	}
	totalCount := len(tasks)

	progress := mpb.New(mpb.WithWidth(internal.BarWidth))
	progressbar := progress.AddBar(
		int64(totalCount),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("poststep "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	for _, itor := range tasks {
		if err := itor(runtime); err != nil {
			return fmt.Errorf("poststep failed: %w", err)
		} // if

		progressbar.Increment()
	} // for

	progress.Wait()
	return nil
}
