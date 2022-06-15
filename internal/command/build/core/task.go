package core

import (
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/xuri/excelize/v2"
)

// TaskProgress 工作進度值
const TaskProgress = taskProgressM + taskProgressS + taskProgressS + taskProgressL + taskProgressM + taskProgressM + taskProgressM
const taskProgressS = 1 // 小型工作進度值
const taskProgressM = 3 // 中型工作進度值
const taskProgressL = 5 // 大型工作進度值

// Task 工作資料
type Task struct {
	global  *Global        // 全域設定
	element *Element       // 項目設定
	bar     *mpb.Bar       // 進度條物件
	excel   *excelize.File // excel物件
	columns []*Column      // 欄位列表
}

// Execute 執行工作
func (this *Task) Execute() error {
	defer this.close()
	err := this.executeExcel()

	if err != nil {
		return err
	} // if

	err = this.executeFields()

	if err != nil {
		return err
	} // if

	err = this.executeNotes()

	if err != nil {
		return err
	} // if

	err = this.executeJson()

	if err != nil {
		return err
	} // if

	err = this.executeJsonCs()

	if err != nil {
		return err
	} // if

	err = this.executeJsonGo()

	if err != nil {
		return err
	} // if

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
func NewTask(global *Global, element *Element, bar *mpb.Bar) *Task {
	return &Task{
		global:  global,
		element: element,
		bar:     bar,
	}
}

// Column 欄位資料
type Column struct {
	Name  string // 欄位名稱
	Note  string // 欄位註解
	Field Field  // 欄位類型
}
