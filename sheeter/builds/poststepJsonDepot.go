package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/tmpls"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// poststepJsonDepot 後製json倉庫資料
type poststepJsonDepot struct {
	*Global                            // 全域設定
	*nameds.Named                      // 命名工具
	*nameds.Json                       // json命名工具
	Struct        []poststepJsonStruct // 結構列表
}

// poststepJsonStruct 後製json結構資料
type poststepJsonStruct struct {
	*nameds.Named      // 命名工具
	Reader        bool // 是否要產生讀取器
}

// PoststepJsonDepotCs 後製json倉庫cs
func PoststepJsonDepotCs(material any, _ chan any) error {
	data, ok := material.(*poststepJsonDepot)

	if ok == false {
		return nil
	} // if

	filename := data.JsonDepotCsPath()

	if err := utils.WriteTmpl(filename, tmpls.JsonDepotCs.Data, data); err != nil {
		return fmt.Errorf("%s: poststep json-depot-cs failed: %w", filename, err)
	} // if

	return nil
}

// PoststepJsonDepotGo 後製json倉庫go
func PoststepJsonDepotGo(material any, _ chan any) error {
	data, ok := material.(*poststepJsonDepot)

	if ok == false {
		return nil
	} // if

	filename := data.JsonDepotGoPath()

	if err := utils.WriteTmpl(filename, tmpls.JsonDepotGo.Data, data); err != nil {
		return fmt.Errorf("%s: poststep json-depot-go failed: %w", filename, err)
	} // if

	return nil
}
