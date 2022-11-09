package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter/layouts"
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/tmpls"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// generateEnum 產生enum資料
type generateEnum struct {
	*Global                       // 全域設定
	*nameds.Named                 // 命名工具
	*nameds.Enum                  // enum命名工具
	Enums         []*layouts.Enum // 列舉列表
}

// GenerateEnumSchema 產生enum架構檔案
func GenerateEnumSchema(material any, _ chan any) error {
	data, ok := material.(*generateEnum)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.EnumPath(), tmpls.EnumSchema.Data, data); err != nil {
		return fmt.Errorf("%s: generate enum schema failed: %w", structName, err)
	} // if

	return nil
}
