package builds

import (
	"fmt"
)

// encodingProto 產生proto資料
func encodingProto(runtimeSector *RuntimeSector) error {
	structName := runtimeSector.StructName()
	rows, err := runtimeSector.GetRows(runtimeSector.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: data line not found", structName)
	} // if

	_, err = packJson(rows, runtimeSector.layoutJson)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	return nil
}
