package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/tmpls"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// poststepProtoDepot 後製proto倉庫資料
type poststepProtoDepot struct {
	*Global                             // 全域設定
	*nameds.Named                       // 命名工具
	*nameds.Proto                       // proto命名工具
	Struct        []poststepProtoStruct // 結構列表
}

// poststepProtoStruct 後製proto結構資料
type poststepProtoStruct struct {
	*nameds.Named      // 命名工具
	Reader        bool // 是否要產生讀取器
}

// PoststepProtoDepotCs 後製proto倉庫cs
func PoststepProtoDepotCs(material any, _ chan any) error {
	data, ok := material.(*poststepProtoDepot)

	if ok == false {
		return nil
	} // if

	filename := data.ProtoDepotCsPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoDepotCs.Data, data); err != nil {
		return fmt.Errorf("%s: poststep proto-depot-cs failed: %w", filename, err)
	} // if

	return nil
}

// PoststepProtoDepotGo 後製proto倉庫go
func PoststepProtoDepotGo(material any, _ chan any) error {
	data, ok := material.(*poststepProtoDepot)

	if ok == false {
		return nil
	} // if

	filename := data.ProtoDepotGoPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoDepotGo.Data, data); err != nil {
		return fmt.Errorf("%s: poststep proto-depot-go failed: %w", filename, err)
	} // if

	return nil
}
