package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoCsDepot 後製proto-cs倉庫程式碼
func poststepProtoCsDepot(runtime *Runtime) error {
	filename := runtime.ProtoCsDepotPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsDepot.Data, runtime); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoDepot 後製proto-go倉庫程式碼
func poststepProtoGoDepot(runtime *Runtime) error {
	filename := runtime.ProtoGoDepotPath()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoDepot.Data, runtime); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", filename); err != nil {
		return fmt.Errorf("%s: poststep failed: gofmt error: %w", filename, err)
	} // if

	return nil
}

// poststepProtoCsBat 後製proto-cs.bat
func poststepProtoCsBat(runtime *Runtime) error {
	filename := runtime.ProtoCsBatFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoCsSh 後製proto-cs.sh
func poststepProtoCsSh(runtime *Runtime) error {
	filename := runtime.ProtoCsShFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsSh.Data, runtime); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoBat 後製proto-go.bat
func poststepProtoGoBat(runtime *Runtime) error {
	filename := runtime.ProtoGoBatFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoBat.Data, runtime); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoSh 後製proto-go.sh
func poststepProtoGoSh(runtime *Runtime) error {
	filename := runtime.ProtoGoShFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoSh.Data, runtime); err != nil {
		return fmt.Errorf("%s: poststep failed: %w", filename, err)
	} // if

	return nil
}
