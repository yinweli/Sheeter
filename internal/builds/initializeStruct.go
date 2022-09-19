package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
)

// initializeStruct 初始化結構
func initializeStruct(runtime *Runtime) error {
	layoutType := layouts.NewLayoutType()

	for _, itor := range runtime.Sector {
		if err := layoutType.Merge(itor.layoutType); err != nil {
			return fmt.Errorf("initialize struct failed: %w", err)
		} // if
	} // for

	for _, itor := range layoutType.TypeNames() {
		runtime.Struct = append(runtime.Struct, &RuntimeStruct{
			layoutType.Types(itor),
		})
	} // for

	return nil
}
