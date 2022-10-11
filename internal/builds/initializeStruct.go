package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
)

// initializeStruct 初始化結構
func initializeStruct(context *Context) error {
	layoutType := layouts.NewLayoutType()

	for _, itor := range context.Sector {
		if err := layoutType.Merge(itor.layoutType); err != nil {
			return fmt.Errorf("initialize struct failed: %w", err)
		} // if
	} // for

	layoutDepend := layouts.NewLayoutDepend()

	for _, itor := range context.Sector {
		if err := layoutDepend.Merge(itor.layoutDepend); err != nil {
			return fmt.Errorf("initialize struct failed: %w", err)
		} // if
	} // for

	for _, itor := range layoutType.TypeNames() {
		context.Struct = append(context.Struct, &ContextStruct{
			types:  layoutType.Types(itor),
			depend: layoutDepend.Depends(itor),
		})
	} // for

	return nil
}
