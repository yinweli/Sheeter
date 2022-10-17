package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoDepotCs 後製proto倉庫cs程式碼
func poststepProtoDepotCs(data *poststepData) error {
	filename := data.ProtoDepotCsPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoDepotCs.Data, data); err != nil {
		return fmt.Errorf("%s: poststep proto-depot-cs failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoDepotGo 後製proto倉庫go程式碼
func poststepProtoDepotGo(data *poststepData) error {
	filename := data.ProtoDepotGoPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoDepotGo.Data, data); err != nil {
		return fmt.Errorf("%s: poststep proto-depot-go failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoBatCs 後製proto-bat-cs
func poststepProtoBatCs(data *poststepData) error {
	filename := data.ProtoBatCsFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoBatCs.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoShCs 後製proto-sh-cs
func poststepProtoShCs(data *poststepData) error {
	filename := data.ProtoShCsFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoShCs.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoBatGo 後製proto-bat-go
func poststepProtoBatGo(data *poststepData) error {
	filename := data.ProtoBatGoFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoBatGo.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoShGo 後製proto-sh-go
func poststepProtoShGo(data *poststepData) error {
	filename := data.ProtoShGoFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoShGo.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}
