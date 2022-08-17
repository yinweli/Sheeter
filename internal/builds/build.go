package builds

import (
	"fmt"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/xuri/excelize/v2"
	"github.com/yinweli/Sheeter/internal/builds/fields"
	"github.com/yinweli/Sheeter/internal/builds/layers"
	"github.com/yinweli/Sheeter/internal/builds/layouts"
	"github.com/yinweli/Sheeter/internal/util"
)

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

	if err := writeJson(content); err != nil {
		return err
	} // if

	bar.Increment()

	if bar != nil { // 讓進度條顯示完成並且有時間畫圖
		bar.SetTotal(maxTask, true)
		time.Sleep(drawTime)
	} // if

	return nil
}

// readExcel 讀取excel檔案
func readExcel(content *Content) error {
	excel, err := excelize.OpenFile(content.ExcelFilePath())

	if err != nil {
		return fmt.Errorf("%s: read excel failed", content.TargetName())
	} // if

	if excel.GetSheetIndex(content.Sheet) == -1 {
		return fmt.Errorf("%s: read excel failed, sheet not found", content.TargetName())
	} // if

	content.excel = excel
	return nil
}

// buildLayout 建立布局資料
func buildLayout(content *Content) error {
	fieldLine, err := content.getColumns(content.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: build layout failed, field line not found", content.TargetName())
	} // if

	layerLine, err := content.getColumns(content.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: build layout failed, layer line not found", content.TargetName())
	} // if

	noteLine, err := content.getColumns(content.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: build layout failed, note line not found", content.TargetName())
	} // if

	builder := layouts.NewBuilder()

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: build layout failed: %w", content.TargetName(), err)
		} // if

		layer, back, err := layers.Parser(catch(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: build layout failed: %w", content.TargetName(), err)
		} // if

		note := catch(noteLine, col)

		if err := builder.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: build layout failed: %w", content.TargetName(), err)
		} // if
	} // for

	pkeyCount := builder.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: build layout failed, pkey duplicate", content.TargetName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: build layout failed, pkey not found", content.TargetName())
	} // if

	content.builder = builder
	return nil
}

// writeJson 輸出json
func writeJson(content *Content) error {
	rows, err := content.getRows(content.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: write json failed, data line not found", content.TargetName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := map[string]interface{}{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		packs, pkey, err := content.builder.Pack(datas)

		if err != nil {
			return fmt.Errorf("%s: write json failed, %w", content.TargetName(), err)
		} // if

		objs[pkey] = packs
	} // for

	err = util.JsonWrite(content.JsonFilePath(), objs, content.Bom)

	if err != nil {
		return fmt.Errorf("%s: write json failed, %w", content.TargetName(), err)
	} // if

	return nil
}

// writeSchema 輸出json架構
func writeSchema(content *Content) error {
	objs := map[string]interface{}{}
}

// catch 從列表中取得項目
func catch(lists []string, index int) string {
	if index >= 0 && index < len(lists) { // 列表的數量可能因為空白格的關係會短缺, 所以要檢查一下
		return lists[index]
	} // if

	return ""
}
