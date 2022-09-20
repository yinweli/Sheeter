package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
)

// initializeStruct 初始化結構
func initializeStruct(runtime *Runtime) error {
	layoutType := layouts.NewLayoutType()

	for _, itor := range runtime.Sector {
		if err := layoutType.Merge(itor.layoutType); err != nil {
			return fmt.Errorf("initialize struct failed, layoutType merge failed: %w", err)
		} // if
	} // for

	layoutDepend := layouts.NewLayoutDepend()

	for _, itor := range runtime.Sector {
		if err := layoutDepend.Merge(itor.layoutDepend); err != nil {
			return fmt.Errorf("initialize struct failed, layoutDepend merge failed: %w", err)
		} // if
	} // for

	for _, itor := range layoutType.TypeNames() {
		types := layoutType.Types(itor)
		depend := layoutDepend.Depends(itor)
		runtime.Struct = append(runtime.Struct, &RuntimeStruct{
			Mixed:  mixeds.NewMixed(types.Excel, types.Sheet),
			Type:   types,
			Depend: depend,
		})
	} // for

	return nil
}
