package pipelines

import (
	"fmt"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
)

// PipelineFunc 管線執行函式類型
type PipelineFunc = func(material any, result chan any) error

// Pipeline 管線執行
func Pipeline(name string, material []any, funcs []PipelineFunc) (result []any, errs []error) {
	if len(material) == 0 || len(funcs) == 0 {
		return []any{}, []error{}
	} // if

	count := len(material) * len(funcs)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(count)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(waitGroup))
	progressbar := progress.AddBar(
		int64(count),
		mpb.AppendDecorators(
			decor.Name(fmt.Sprintf("%-20s ", name)),
			decor.CountersNoUnit("(%6d/%6d) ", decor.WCSyncWidth),
			decor.NewPercentage("%d "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)
	output := make(chan any, count)

	for _, itor := range material {
		temp := itor // 多執行緒需要使用中間變數

		go func() {
			for _, itor := range funcs {
				if err := itor(temp, output); err != nil {
					output <- err
				} // if

				waitGroup.Done()
				progressbar.Increment()
			} // for
		}()
	} // for

	go func() {
		for {
			select {
			case value := <-output:
				if err, ok := value.(error); ok {
					errs = append(errs, err)
				} else {
					result = append(result, value)
				} // if

			default:
				if progressbar.Completed() {
					return
				} // if
			} // switch
		} // for
	}()

	progress.Wait()
	return result, errs
}
