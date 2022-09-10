package builds

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/builds/layouts"
	"github.com/yinweli/Sheeter/internal/util"
)

// BuildSector 區段建置
func BuildSector(config *Config, code *Code) (errs []error) {
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
			Code:    code,
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

// Sector 區段資料
type Sector struct {
	Global                   // 全域設定
	Element                  // 項目設定
	*Code                    // 模板資料
	excel   *excelize.File   // excel物件
	builder *layouts.Builder // 布局建造器
}

// AppName 取得程式名稱
func (this *Sector) AppName() string {
	return internal.AppName
}

// Namespace 取得命名空間名稱
func (this *Sector) Namespace() string {
	name := this.combine(params{})
	return name
}

// StructName 取得結構名稱
func (this *Sector) StructName() string {
	return internal.Struct
}

// ReaderName 取得讀取器名稱
func (this *Sector) ReaderName() string {
	return internal.Reader
}

// FileJson 取得json檔名路徑
func (this *Sector) FileJson() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJson,
	})
	return filepath.Join(internal.PathJson, name)
}

// FileJsonSchema 取得json架構檔名路徑
func (this *Sector) FileJsonSchema() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJson,
	})
	return filepath.Join(internal.PathJsonSchema, name)
}

// FileJsonCsCode 取得json-cs程式碼檔名路徑
func (this *Sector) FileJsonCsCode() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtCs,
	})
	return filepath.Join(internal.PathJsonCs, name)
}

// FileJsonCsReader 取得json-cs讀取器檔名路徑
func (this *Sector) FileJsonCsReader() string {
	name := this.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
	})
	return filepath.Join(internal.PathJsonCs, name)
}

// FileJsonGoCode 取得json-go程式碼檔名路徑
func (this *Sector) FileJsonGoCode() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtGo,
	})
	return filepath.Join(internal.PathJsonGo, name)
}

// FileJsonGoReader 取得json-go讀取器檔名路徑
func (this *Sector) FileJsonGoReader() string {
	name := this.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
	})
	return filepath.Join(internal.PathJsonGo, name)
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	middle     string // excel與sheet的中間字串
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}

// combine 取得組合名稱
func (this *Sector) combine(params params) string {
	excel := util.FileName(this.Excel)

	if params.excelUpper {
		excel = util.FirstUpper(excel)
	} else {
		excel = util.FirstLower(excel)
	} // if

	sheet := this.Sheet

	if params.sheetUpper {
		sheet = util.FirstUpper(sheet)
	} else {
		sheet = util.FirstLower(sheet)
	} // if

	items := []string{excel, params.middle, sheet, params.last}

	if params.ext != "" {
		items = append(items, ".", params.ext)
	} // if

	return strings.Join(items, "")
}

// Close 關閉excel物件
func (this *Sector) Close() {
	if this.excel != nil {
		_ = this.excel.Close()
	} // if
}

// GetRows 取得表格行資料, line從1起算; 如果該行不存在, 回傳成功並取得最後一行物件
func (this *Sector) GetRows(line int) (rows *excelize.Rows, err error) {
	if line <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("get row failed, row <= 0")
	} // if

	rows, err = this.excel.Rows(this.Sheet)

	if err != nil {
		return nil, fmt.Errorf("get row failed: %w", err)
	} // if

	for l := 0; l < line; l++ {
		rows.Next()
	} // for

	return rows, nil
}

// GetColumns 取得表格行內容, line從1起算; 如果該行不存在, 回傳失敗
func (this *Sector) GetColumns(line int) (cols []string, err error) {
	if line <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("get columns failed, row <= 0")
	} // if

	rows, err := this.excel.Rows(this.Sheet)

	if err != nil {
		return nil, fmt.Errorf("get columns failed: %w", err)
	} // if

	defer func() { _ = rows.Close() }()

	for l := 0; l < line; l++ {
		if rows.Next() == false {
			return nil, fmt.Errorf("get columns failed, row not found")
		} // if
	} // for

	cols, err = rows.Columns()

	if err != nil {
		return nil, fmt.Errorf("get columns failed, invalid columns: %w", err)
	} // if

	if cols == nil {
		cols = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return cols, nil
}
