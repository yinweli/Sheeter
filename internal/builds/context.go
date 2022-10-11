package builds

import (
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
)

// Context 執行資料
type Context struct {
	*Config                  // 設定資料
	Sector  []*ContextSector // 區段資料列表
	Struct  []*ContextStruct // 結構資料列表
}

// Close 關閉資料
func (this *Context) Close() {
	for _, itor := range this.Sector {
		itor.Close()
	} // for
}

// ContextSector 執行區段資料
type ContextSector struct {
	Element                            // 項目設定
	excel        *excels.Excel         // excel物件
	layoutJson   *layouts.LayoutJson   // json布局器
	layoutType   *layouts.LayoutType   // 類型布局器
	layoutDepend *layouts.LayoutDepend // 依賴布局器
}

// Close 關閉資料
func (this *ContextSector) Close() {
	if this.excel != nil {
		this.excel.Close()
		this.excel = nil
	} // if
}

// ContextStruct 執行結構資料
type ContextStruct struct {
	types  *layouts.Type // 類型資料
	depend []string      // 依賴列表
}
