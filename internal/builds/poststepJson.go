package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepJsonCsDepot 後製json-cs倉庫程式碼
func poststepJsonCsDepot(data *poststepData) error {
	filename := data.JsonCsDepotPath()

	if err := utils.WriteTmpl(filename, tmpls.JsonCsDepot.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepJsonGoDepot 後製json-go倉庫程式碼
func poststepJsonGoDepot(data *poststepData) error {
	filename := data.JsonGoDepotPath()

	if err := utils.WriteTmpl(filename, tmpls.JsonGoDepot.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", filename); err != nil {
		return fmt.Errorf("%s: poststep failed: gofmt error: %w", filename, err)
	} // if

	return nil
}
