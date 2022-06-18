package core

import (
	"fmt"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/xuri/excelize/v2"
)

// Task 工作資料
type Task struct {
	global  *Global        // 全域設定
	element *Element       // 項目設定
	bar     *mpb.Bar       // 進度條物件
	excel   *excelize.File // excel物件
	columns []*Column      // 欄位列表
}

// Execute 執行工作
func (this *Task) Execute(progress *mpb.Progress) error {
	defer this.close()

	this.bar = progress.AddBar(
		7, // 目前Task中有7項任務要進行
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%-20s", this.originalName())),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	err := this.executeExcel()

	if err != nil {
		return err
	} // if

	err = this.executeColumn()

	if err != nil {
		return err
	} // if

	err = this.executeSchema()

	if err != nil {
		return err
	} // if

	// TODO: schema -> json -> json cs -> json go -> lua

	if this.bar != nil { // 讓進度條顯示完成並且有時間畫圖
		this.bar.SetTotal(100, true)
		time.Sleep(10 * time.Millisecond)
	} // if

	return nil
}

// close 結束工作
func (this *Task) close() {
	if this.excel != nil {
		_ = this.excel.Close()
	} // if
}

// NewTask 建立工作資料
func NewTask(global *Global, element *Element) *Task {
	return &Task{
		global:  global,
		element: element,
	}
}

// Column 欄位資料
type Column struct {
	Name  string // 欄位名稱
	Note  string // 欄位註解
	Field Field  // 欄位類型
}
