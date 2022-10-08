package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
)

// Runtime 執行期資料
type Runtime struct {
	Global                        // 全域設定
	mixeds.Mixed                  // 綜合工具
	Sector       []*RuntimeSector // 區段資料列表
	Struct       []*RuntimeStruct // 結構資料列表
}

// RuntimeSector 執行期區段資料
type RuntimeSector struct {
	Global                              // 全域設定
	Element                             // 項目設定
	*mixeds.Mixed                       // 綜合工具
	excel         *excels.Excel         // excel物件
	layoutJson    *layouts.LayoutJson   // json布局器
	layoutType    *layouts.LayoutType   // 類型布局器
	layoutDepend  *layouts.LayoutDepend // 依賴布局器
}

// OpenExcel 開啟excel
func (this *RuntimeSector) OpenExcel() error {
	if err := this.excel.Open(this.Element.Excel); err != nil {
		return fmt.Errorf("open excel failed: %w", err)
	} // if

	if this.excel.SheetExist(this.Element.Sheet) == false {
		return fmt.Errorf("open excel failed: sheet not found")
	} // if

	return nil
}

// CloseExcel 關閉excel
func (this *RuntimeSector) CloseExcel() {
	this.excel.Close()
}

// GetExcelLine 取得excel行資料, index從1起算; 當行不存在時不會發生錯誤
func (this *RuntimeSector) GetExcelLine(index int) (line *excels.Line, err error) {
	line, err = this.excel.GetLine(this.Element.Sheet, index)

	if err != nil {
		return nil, fmt.Errorf("get excel line failed: %w", err)
	} // if

	return line, nil
}

// GetExcelData 取得excel行資料列表, index從1起算; 當行不存在時會發生錯誤
func (this *RuntimeSector) GetExcelData(index int) (data []string, err error) {
	data, err = this.excel.GetData(this.Element.Sheet, index)

	if err != nil {
		return nil, fmt.Errorf("get excel data failed: %w", err)
	} // if

	return data, nil
}

// RuntimeStruct 執行期結構資料
type RuntimeStruct struct {
	Global                   // 全域設定
	*mixeds.Mixed            // 綜合工具
	*layouts.Type            // 類型資料
	SimpleNamespace bool     // 是否用簡單的命名空間名稱
	Depend          []string // 依賴列表
}
