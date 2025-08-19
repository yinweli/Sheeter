package pipelines

import (
	"fmt"
	"sync"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/v3/sheeter"
)

// PipelineFunc 管線執行函式類型
type PipelineFunc[T any] func(material T, result chan any) error

// Pipeline 管線執行
func Pipeline[T any](name string, material []T, pipeline []PipelineFunc[T]) (result []any, err []error) {
	if len(material) == 0 || len(pipeline) == 0 {
		return []any{}, []error{}
	} // if

	count := len(material) * len(pipeline)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(count)
	progress := mpb.New(mpb.WithWidth(sheeter.BarWidth), mpb.WithWaitGroup(waitGroup))
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
			wrong := false

			for _, f := range pipeline {
				if wrong == false { // 如果管線中有執行失敗, 則就不能再執行下去
					if err := f(temp, output); err != nil {
						output <- err
						wrong = true
					} // if
				} // if

				time.Sleep(time.Millisecond) // 用來預防資料來不及添加到結果/錯誤列表中
				waitGroup.Done()
				progressbar.Increment()
			} // for
		}()
	} // for

	go func() {
		for {
			select {
			case value := <-output:
				if e, ok := value.(error); ok {
					err = append(err, e)
				} else {
					result = append(result, value)
				} // if

			default:
				if progressbar.Completed() {
					close(output)
					return
				} // if
			} // switch
		} // for
	}()

	progress.Wait()
	return result, err
}
