package buildoo

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepFormat 後製格式化資料
type poststepFormat struct {
}

// PoststepFormatCs 後製格式化cs
func PoststepFormatCs(material any) error {
	_, ok := material.(*poststepFormat)

	if ok == false {
		return nil
	} // if

	if err := utils.ShellRun("dotnet", "csharpier", "."); err != nil {
		return fmt.Errorf("poststep format cs failed: %w", err)
	} // if

	return nil
}

// PoststepFormatGo 後製格式化cs
func PoststepFormatGo(material any) error {
	_, ok := material.(*poststepFormat)

	if ok == false {
		return nil
	} // if

	if err := utils.ShellRun("gofmt", "-w", "."); err != nil {
		return fmt.Errorf("poststep format go failed: %w", err)
	} // if

	return nil
}

// PoststepFormatProto 後製格式化proto
func PoststepFormatProto(material any) error {
	_, ok := material.(*poststepFormat)

	if ok == false {
		return nil
	} // if

	if err := utils.ShellRun("buf", "format", "-w", "."); err != nil {
		return fmt.Errorf("poststep format proto failed: %w", err)
	} // if

	return nil
}
