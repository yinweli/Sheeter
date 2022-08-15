package tasks

import (
	"fmt"
	"time"

	"github.com/yinweli/Sheeter/internal/build/layouts"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/xuri/excelize/v2"
)

const maxTask = 8                      // 最大工作數量
const drawTime = 10 * time.Millisecond // 畫圖時間

// Task 工作資料
type Task struct {
	Path        string           // 來源excel路徑
	Bom         bool             // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int              // 欄位行號(1為起始行)
	LineOfLayer int              // 階層行號(1為起始行)
	LineOfNote  int              // 註解行號(1為起始行)
	LineOfData  int              // 資料起始行號(1為起始行)
	Excel       string           // excel檔案名稱
	Sheet       string           // excel表單名稱
	Progress    *mpb.Progress    // 進度條產生器
	bar         *mpb.Bar         // 進度條物件
	xlsfile     *excelize.File   // excel物件
	builder     *layouts.Builder // 布局建造器
}

// Run 執行工作
func (this *Task) Run() error {
	defer this.close()

	if err := this.check(); err != nil {
		return err
	} // if

	this.bar = this.Progress.AddBar(
		maxTask,
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%-20s", this.targetName())),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	if err := this.excel(); err != nil {
		return err
	} // if

	if err := this.column(); err != nil {
		return err
	} // if

	if err := this.json(); err != nil {
		return err
	} // if

	if err := this.jsonSchema(); err != nil {
		return err
	} // if

	if err := this.jsonCs(); err != nil {
		return err
	} // if

	if err := this.jsonCsReader(); err != nil {
		return err
	} // if

	if err := this.jsonGo(); err != nil {
		return err
	} // if

	if err := this.jsonGoReader(); err != nil {
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
	if this.xlsfile != nil {
		_ = this.xlsfile.Close()
	} // if
}

// check 檢查工作
func (this *Task) check() error {
	if this.LineOfField <= 0 {
		return fmt.Errorf("lineOfField <= 0")
	} // if

	if this.LineOfLayer <= 0 {
		return fmt.Errorf("lineOfLayer <= 0")
	} // if

	if this.LineOfNote <= 0 {
		return fmt.Errorf("lineOfNote <= 0")
	} // if

	if this.LineOfData <= 0 {
		return fmt.Errorf("lineOfData <= 0")
	} // if

	if this.LineOfField >= this.LineOfData {
		return fmt.Errorf("lineOfField(%d) >= lineOfData(%d)", this.LineOfField, this.LineOfData)
	} // if

	if this.LineOfLayer >= this.LineOfData {
		return fmt.Errorf("lineOfLayer(%d) >= lineOfData(%d)", this.LineOfLayer, this.LineOfData)
	} // if

	if this.LineOfNote >= this.LineOfData {
		return fmt.Errorf("lineOfNote(%d) >= lineOfData(%d)", this.LineOfNote, this.LineOfData)
	} // if

	if this.Excel == "" {
		return fmt.Errorf("excel empty")
	} // if

	if this.Sheet == "" {
		return fmt.Errorf("sheet empty")
	} // if

	if this.Progress == nil {
		return fmt.Errorf("progress nil")
	} // if

	return nil
}
