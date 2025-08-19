package pipelines

import (
	"fmt"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/v3/sheeter"
)

// Pipeline 管線執行
func Pipeline[T any](name string, material []T, execute []Execute[T]) (result []any, err []error) {
	count := len(material) * len(execute)

	if count <= 0 {
		return nil, nil
	} // if

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
	output := make(chan Output, count)

	for _, itor := range material {
		temp := itor // 多執行緒需要使用中間變數

		go func() {
			succ := true

			for _, exec := range execute {
				if succ { // 如果管線中有執行失敗, 則就不能再執行下去
					o := exec(temp)
					succ = o.Error == nil
					output <- o
				} // if

				waitGroup.Done()
				progressbar.Increment()
			} // for
		}()
	} // for

	go func() {
		waitGroup.Wait()
		close(output)
	}()

	for itor := range output {
		if itor.Error == nil {
			result = append(result, itor.Result...)
		} else {
			err = append(err, itor.Error)
		} // if
	} // for

	progress.Wait()
	return result, err
}

// Execute 管線執行函式類型
type Execute[T any] func(material T) Output

// Output 管線結果資料
type Output struct {
	Result []any // 結果列表
	Error  error // 錯誤物件
}
