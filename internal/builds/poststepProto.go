package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoCsBat 後製proto-cs.bat
func poststepProtoCsBat(runtime *Runtime) error {
	path := runtime.FileProtoCsBat()

	if err := utils.WriteTmpl(path, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", path, err)
	} // if

	return nil
}

// poststepProtoCsSh 後製proto-cs.sh
func poststepProtoCsSh(runtime *Runtime) error {
	path := runtime.FileProtoCsSh()

	if err := utils.WriteTmpl(path, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", path, err)
	} // if

	return nil
}

// poststepProtoGoBat 後製proto-go.bat
func poststepProtoGoBat(runtime *Runtime) error {
	path := runtime.FileProtoGoBat()

	if err := utils.WriteTmpl(path, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", path, err)
	} // if

	return nil
}

// poststepProtoGoSh 後製proto-go.sh
func poststepProtoGoSh(runtime *Runtime) error {
	path := runtime.FileProtoGoSh()

	if err := utils.WriteTmpl(path, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", path, err)
	} // if

	return nil
}
