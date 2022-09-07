package builds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal/util"
)

// SectorJsonCsCode 輸出json-cs程式碼
func SectorJsonCsCode(sector *Sector) error {
	if err := os.MkdirAll(filepath.Dir(sector.FileJsonCsCode()), os.ModePerm); err != nil {
		return fmt.Errorf("sector json-cs code failed: %w", err)
	} // if

	options := []string{
		"--src", sector.FileJsonSchema(),
		"--src-lang", "json",
		"--out", sector.FileJsonCsCode(),
		"--lang", "cs",
		"--namespace", sector.Namespace(),
		"--top-level", sector.StructName(),
		"--array-type", "array",
		"--features", "attributes-only",
	}

	if err := util.ShellRun("quicktype", options...); err != nil {
		return fmt.Errorf("sector json-cs code failed, quicktype error: %w", err)
	} // if

	return nil
}

// SectorJsonCsReader 輸出json-cs讀取器
func SectorJsonCsReader(sector *Sector) error {
	if err := util.WriteTmpl(sector.FileJsonCsReader(), sector.JsonCsReader, sector, sector.Bom); err != nil {
		return fmt.Errorf("entire json-cs reader failed: %w", err)
	} // if

	return nil
}
