package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoCsDepot 後製proto-cs倉庫程式碼
func poststepProtoCsDepot(data *poststepData) error {
	filename := data.ProtoCsDepotPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsDepot.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoDepot 後製proto-go倉庫程式碼
func poststepProtoGoDepot(data *poststepData) error {
	filename := data.ProtoGoDepotPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoDepot.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", filename); err != nil {
		return fmt.Errorf("%s: poststep failed: gofmt error: %w", filename, err)
	} // if

	return nil
}

// poststepProtoCsBat 後製proto-cs.bat
func poststepProtoCsBat(data *poststepData) error {
	filename := data.ProtoCsBatFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsBat.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoCsSh 後製proto-cs.sh
func poststepProtoCsSh(data *poststepData) error {
	filename := data.ProtoCsShFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsSh.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoBat 後製proto-go.bat
func poststepProtoGoBat(data *poststepData) error {
	filename := data.ProtoGoBatFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoBat.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoSh 後製proto-go.sh
func poststepProtoGoSh(data *poststepData) error {
	filename := data.ProtoGoShFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoSh.Data, data); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}
