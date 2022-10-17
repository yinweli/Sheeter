package buildoo

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateProto 產生proto資料
type generateProto struct {
	*Global                // 全域設定
	*nameds.Named          // 命名工具
	*nameds.Field          // 欄位命名工具
	*nameds.Proto          // proto命名工具
	*layouts.Type          // 類型資料
	Depend        []string // 依賴列表
}

// GenerateProtoSchema 產生proto架構檔案
func GenerateProtoSchema(material any) error {
	data, ok := material.(*generateProto)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoPath(), tmpls.ProtoSchema.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto schema failed: %w", structName, err)
	} // if

	return nil
}

// GenerateProtoReaderCs 產生proto讀取器cs
func GenerateProtoReaderCs(material any) error {
	data, ok := material.(*generateProto)

	if ok == false {
		return nil
	} // if

	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoReaderCsPath(), tmpls.ProtoReaderCs.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto-reader-cs failed: %w", structName, err)
	} // if

	return nil
}

// GenerateProtoReaderGo 產生proto讀取器go
func GenerateProtoReaderGo(material any) error {
	data, ok := material.(*generateProto)

	if ok == false {
		return nil
	} // if

	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoReaderGoPath(), tmpls.ProtoReaderGo.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto-reader-cs failed: %w", structName, err)
	} // if

	return nil
}
