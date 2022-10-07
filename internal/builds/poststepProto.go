package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepProtoCsBat 後製proto-cs.bat
func poststepProtoCsBat(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.ProtoCsBatFile, tmpls.ProtoCsBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.ProtoCsBatFile, err)
	} // if

	return nil
}

// poststepProtoCsSh 後製proto-cs.sh
func poststepProtoCsSh(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.ProtoCsShFile, tmpls.ProtoCsSh.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.ProtoCsShFile, err)
	} // if

	return nil
}

// poststepProtoGoBat 後製proto-go.bat
func poststepProtoGoBat(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.ProtoGoBatFile, tmpls.ProtoGoBat.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.ProtoGoBatFile, err)
	} // if

	return nil
}

// poststepProtoGoSh 後製proto-go.sh
func poststepProtoGoSh(runtime *Runtime) error {
	if err := utils.WriteTmpl(internal.ProtoGoShFile, tmpls.ProtoGoSh.Data, runtime); err != nil {
		return fmt.Errorf("poststep %s failed: %w", internal.ProtoGoShFile, err)
	} // if

	return nil
}
