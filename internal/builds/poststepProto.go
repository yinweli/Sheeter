package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoCsBat 後製proto-cs.bat
func poststepProtoCsBat(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.FileProtoCsBat, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.FileProtoCsBat, err)
	} // if

	return nil
}

// poststepProtoCsSh 後製proto-cs.sh
func poststepProtoCsSh(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.FileProtoCsSh, tmpls.ProtoCsSh.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.FileProtoCsSh, err)
	} // if

	return nil
}

// poststepProtoGoBat 後製proto-go.bat
func poststepProtoGoBat(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.FileProtoGoBat, tmpls.ProtoGoBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.FileProtoGoBat, err)
	} // if

	return nil
}

// poststepProtoGoSh 後製proto-go.sh
func poststepProtoGoSh(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.FileProtoGoSh, tmpls.ProtoGoSh.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.FileProtoGoSh, err)
	} // if

	return nil
}
