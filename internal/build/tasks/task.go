package tasks

import (
	"fmt"
	"time"

	"github.com/yinweli/Sheeter/internal/build/fields"

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
	Name  string       // 欄位名稱
	Note  string       // 欄位註解
	Field fields.Field // 欄位類型
}

// Global 全域設定
type Global struct {
	ExcelPath   string `yaml:"excelPath"`   // 來源excel路徑
	Bom         bool   `yaml:"bom"`         // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int    `yaml:"lineOfField"` // 欄位行號(1為起始行)
	LineOfLayer int    `yaml:"lineOfLayer"` // 階層行號(1為起始行)
	LineOfNote  int    `yaml:"lineOfNote"`  // 註解行號(1為起始行)
	LineOfData  int    `yaml:"lineOfData"`  // 資料起始行號(1為起始行)
}

// Element 項目設定
type Element struct {
	Excel string `yaml:"excel"` // excel檔案名稱
	Sheet string `yaml:"sheet"` // excel表單名稱
}
