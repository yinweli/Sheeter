package buildoo

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
)

// initializeElement 初始化項目資料
type initializeElement struct {
	*Global                            // 全域設定
	Element                            // 項目設定, 這裡複製會比取用指標好
	excel        *excels.Excel         // excel物件
	layoutJson   *layouts.LayoutJson   // json布局器
	layoutType   *layouts.LayoutType   // 類型布局器
	layoutDepend *layouts.LayoutDepend // 依賴布局器
}

// InitializeElement 初始化項目
func InitializeElement(material any) error {
	data, ok := material.(*initializeElement)

	if ok == false {
		return nil
	} // if

	fmt.Println(data) // TODO: initializeElement
	return nil
}
