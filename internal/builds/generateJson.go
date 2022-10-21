package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateJson 產生json資料
type generateJson struct {
	*Global                        // 全域設定
	*nameds.Named                  // 命名工具
	*nameds.Field                  // 欄位命名工具
	*nameds.Json                   // json命名工具
	Reader        bool             // 是否要產生讀取器
	Fields        []*layouts.Field // 欄位列表
}

// GenerateJsonStructCs 產生json結構cs
func GenerateJsonStructCs(material any) error {
	data, ok := material.(*generateJson)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonStructCsPath(), tmpls.JsonStructCs.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-struct-cs failed: %w", structName, err)
	} // if

	return nil
}

// GenerateJsonReaderCs 產生json讀取器cs
func GenerateJsonReaderCs(material any) error {
	data, ok := material.(*generateJson)

	if ok == false {
		return nil
	} // if

	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonReaderCsPath(), tmpls.JsonReaderCs.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-reader-cs failed: %w", structName, err)
	} // if

	return nil
}

// GenerateJsonStructGo 產生json結構go
func GenerateJsonStructGo(material any) error {
	data, ok := material.(*generateJson)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonStructGoPath(), tmpls.JsonStructGo.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-struct-go failed: %w", structName, err)
	} // if

	return nil
}

// GenerateJsonReaderGo 產生json讀取器go
func GenerateJsonReaderGo(material any) error {
	data, ok := material.(*generateJson)

	if ok == false {
		return nil
	} // if

	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonReaderGoPath(), tmpls.JsonReaderGo.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-reader-go failed: %w", structName, err)
	} // if

	return nil
}
