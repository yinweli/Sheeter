package builds

import (
	"fmt"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
)

// TODO: 增加Build函式, 裡面呼叫buildSector & buildEntire
// TODO: BuildSector => buildSector

// BuildSector 區段建置
func BuildSector(config *Config) (errs []error) {
	const tasks = 7 // 區段建置的工作數量

	count := len(config.Elements)
	errors := make(chan error, count) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(&signaler))

	signaler.Add(count)

	for _, itor := range config.Elements {
		sector := &Sector{
			Global:  config.Global,
			Element: itor,
		}

		go func() {
			defer signaler.Done()
			defer sector.Close()

			bar := progress.AddBar(
				tasks,
				mpb.PrependDecorators(
					decor.Name(fmt.Sprintf("%-20s", sector.StructName())),
					decor.Percentage(decor.WCSyncSpace),
				),
				mpb.AppendDecorators(
					decor.OnComplete(decor.Spinner(nil), "complete"),
				),
			)

			if err := SectorInit(sector); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := SectorJson(sector); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := SectorJsonSchema(sector); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := SectorJsonCsCode(sector); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := SectorJsonCsReader(sector); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := SectorJsonGoCode(sector); err != nil {
				errors <- err
				return
			} // if

			bar.Increment()

			if err := SectorJsonGoReader(sector); err != nil {
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
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
