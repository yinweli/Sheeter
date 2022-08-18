package builds

import (
	"fmt"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

// TODO: 單元測試還是讓他們有分隔行比較好(參考buildLayout_test.go)
// TODO: 檢查全部檔案的import順序
// TODO: 單元測試資料的測試設定檔
// TODO: 所有檔案的tab/space格式化

const maxTask = 8                      // 最大工作數量
const drawTime = 10 * time.Millisecond // 進度條繪製時間

// Build 進行表格轉換
func Build(content *Content) error {
	bar := content.Progress.AddBar(
		maxTask,
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%-20s", content.TargetName())),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)
	defer content.close()

	if err := readExcel(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := buildLayout(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := writeSchema(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := writeJson(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := writeJsonCs(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := writeJsonCsReader(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := writeJsonGo(content); err != nil {
		return err
	} // if

	bar.Increment()

	if err := writeJsonGoReader(content); err != nil {
		return err
	} // if

	bar.Increment()

	if bar != nil { // 讓進度條顯示完成並且有時間畫圖
		bar.SetTotal(maxTask, true)
		time.Sleep(drawTime)
	} // if

	return nil
}
