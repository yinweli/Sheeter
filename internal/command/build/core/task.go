package core

import (
	"fmt"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/xuri/excelize/v2"
)

const maxTask = 9                      // 最大工作數量
const drawTime = 10 * time.Millisecond // 畫圖時間

// Task 工作資料
type Task struct {
	global  *Global        // 全域設定
	element *Element       // 項目設定
	bar     *mpb.Bar       // 進度條物件
	excel   *excelize.File // excel物件
	columns []*Column      // 欄位列表
}

// Run 執行工作
func (this *Task) Run(progress *mpb.Progress) error {
	defer this.close()

	this.bar = progress.AddBar(
		maxTask,
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%-20s", this.originalName())),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	err := this.runExcel()

	if err != nil {
		return err
	} // if

	err = this.runColumn()

	if err != nil {
		return err
	} // if

	err = this.runJson()

	if err != nil {
		return err
	} // if

	err = this.runJsonSchema()

	if err != nil {
		return err
	} // if

	err = this.runJsonCs()

	if err != nil {
		return err
	} // if

	err = this.runJsonCsReader()

	if err != nil {
		return err
	} // if

	err = this.runJsonGo()

	if err != nil {
		return err
	} // if

	err = this.runJsonGoReader()

	if err != nil {
		return err
	} // if

	err = this.runLua()

	if err != nil {
		return err
	} // if

	if this.bar != nil { // 讓進度條顯示完成並且有時間畫圖
		this.bar.SetTotal(maxTask, true)
		time.Sleep(drawTime)
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

// NewColumn 建立欄位資料
func NewColumn(name, note string, field Field) *Column {
	return &Column{
		Name:  name,
		Note:  note,
		Field: field,
	}
}
