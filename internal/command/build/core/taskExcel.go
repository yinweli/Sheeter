package core

import (
	"fmt"
	"io"
	"os"

	"Sheeter/internal/util"

	"github.com/schollz/progressbar/v3"
	"github.com/xuri/excelize/v2"
)

// TaskExcel 讀取excel檔案並獲取表格列表
func TaskExcel(ctx *Context, writer io.Writer) error {
	file, err := os.Open(ctx.ExcelFilePath())

	if err != nil {
		return fmt.Errorf("file not found: %s", ctx.LogName())
	} // if

	defer func() {
		_ = file.Close()
	}()

	stat, err := file.Stat()

	if err != nil {
		return fmt.Errorf("file size error: %s", ctx.LogName())
	} // if

	// TODO: 這個進度條有問題

	reader := progressbar.NewReader(file, util.NewProgress(stat.Size(), ctx.LogName(), writer)) // 建立可以同時讀取檔案並且推進進度條的讀取器
	excel, err := excelize.OpenReader(&reader)

	if err != nil {
		return fmt.Errorf("excel read failed: %s", ctx.LogName())
	} // if

	defer func() {
		_ = excel.Close()
	}()

	// TODO: 結果花時間的地方在這邊
	// TODO: 10萬條紀錄太多了 = =

	sheets, err := excel.GetRows(ctx.Element.Sheet)

	if err != nil {
		return fmt.Errorf("sheet not found: %s", ctx.LogName())
	} // if

	if len(sheets) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return fmt.Errorf("sheet have too less line: %s", ctx.LogName())
	} // if

	// 進度值的算法:
	// 讀取檔案(檔案大小) +
	// 建立欄位(1) + 建立註解(1) + 建立資料(表格數量) + 主要索引檢查(1) +
	// json轉換(表格數量) + json存檔(1) + c++存檔(1) + cs存檔(1) + go存檔(1)
	ctx.Progress = util.NewProgress(int64(ctx.Sheets.Size()*2)+10, ctx.LogName(), writer)
	ctx.Sheets = sheets
	return nil
}
