package builds

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingProto 編碼proto資料
type encodingProto struct {
	*Global                           // 全域設定
	*nameds.Named                     // 命名工具
	*nameds.Proto                     // proto命名工具
	excel         *excels.Excel       // excel物件
	layoutData    *layouts.LayoutData // 資料布局器
}

// EncodingProto 編碼proto資料
func EncodingProto(material any, _ chan any) error {
	data, ok := material.(*encodingProto)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()
	sheet, err := data.excel.Get(data.Named.SheetName)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: sheet not found", structName)
	} // if

	sheet.Nextn(data.LineOfData)
	json, err := layouts.JsonPack(sheet, data.layoutData, data.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	filename := data.ProtoName()
	message := data.StorerMessage(data.SimpleNamespace)
	imports := []string{filepath.Join(internal.ProtoPath, internal.SchemaPath)}
	proto, err := utils.JsonToProto(filename, message, imports, json)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(data.ProtoDataPath(), proto); err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	return nil
}
