package builds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal/codes"
	"github.com/yinweli/Sheeter/internal/utils"
)

// SectorJsonGoCode 輸出json-go程式碼
func SectorJsonGoCode(sector *Sector) error {
	if err := os.MkdirAll(filepath.Dir(sector.FileJsonGoCode()), os.ModePerm); err != nil {
		return fmt.Errorf("sector json-go code failed: %w", err)
	} // if

	options := []string{
		"--src", sector.FileJsonSchema(),
		"--src-lang", "json",
		"--out", sector.FileJsonGoCode(),
		"--lang", "go",
		"--package", sector.Namespace(),
		"--top-level", sector.StructName(),
		"--just-types-and-package",
	}

	if err := utils.ShellRun("quicktype", options...); err != nil {
		return fmt.Errorf("sector json-go code failed, quicktype error: %w", err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", sector.FileJsonGoCode()); err != nil {
		return fmt.Errorf("sector json-go code failed, gofmt error: %w", err)
	} // if

	return nil
}

// SectorJsonGoReader 輸出json-go讀取器
func SectorJsonGoReader(sector *Sector) error {
	if err := utils.WriteTmpl(sector.FileJsonGoReader(), codes.JsonGoReader.Code, sector); err != nil {
		return fmt.Errorf("sector json-go reader failed: %w", err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", sector.FileJsonGoReader()); err != nil {
		return fmt.Errorf("sector json-go reader failed, gofmt error: %w", err)
	} // if

	return nil
}
