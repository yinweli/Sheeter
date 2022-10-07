package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoCsDepot 後製proto-cs倉庫程式碼
func poststepProtoCsDepot(runtime *Runtime) error {
	return nil
}

// poststepProtoGoDepot 後製proto-go倉庫程式碼
func poststepProtoGoDepot(runtime *Runtime) error {
	return nil
}

// poststepProtoCsBat 後製proto-cs.bat
func poststepProtoCsBat(runtime *Runtime) error {
	filename := runtime.ProtoCsBatFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoCsSh 後製proto-cs.sh
func poststepProtoCsSh(runtime *Runtime) error {
	filename := runtime.ProtoCsShFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoCsSh.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoBat 後製proto-go.bat
func poststepProtoGoBat(runtime *Runtime) error {
	filename := runtime.ProtoGoBatFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", filename, err)
	} // if

	return nil
}

// poststepProtoGoSh 後製proto-go.sh
func poststepProtoGoSh(runtime *Runtime) error {
	filename := runtime.ProtoGoShFile()

	if err := utils.WriteTmpl(filename, tmpls.ProtoGoSh.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", filename, err)
	} // if

	return nil
}
