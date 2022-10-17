package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepJsonDepotCs 後製json倉庫cs程式碼
func poststepJsonDepotCs(data *poststepData) error {
	filename := data.JsonDepotCsPath()

	if err := utils.WriteTmpl(filename, tmpls.JsonDepotCs.Data, data); err != nil {
		return fmt.Errorf("%s: poststep json-depot-cs failed: %w", filename, err)
	} // if

	return nil
}

// poststepJsonDepotGo 後製json-go倉庫程式碼
func poststepJsonDepotGo(data *poststepData) error {
	filename := data.JsonDepotGoPath()

	if err := utils.WriteTmpl(filename, tmpls.JsonDepotGo.Data, data); err != nil {
		return fmt.Errorf("%s: poststep json-go-depot failed: %w", filename, err)
	} // if

	return nil
}
