package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/workflow"
)

// Initialize 初始化
func Initialize(context *Context) []error {
	for _, itor := range context.Elements {
		context.Sector = append(context.Sector, &ContextSector{
			Element: itor, // 這裡複製會比取用指標好
		})
	} // for

	totalCount := len(context.Sector)

	if totalCount <= 0 {
		return []error{}
	} // if

	work := workflow.NewWorkflow("initialize ", totalCount)

	for _, itor := range context.Sector {
		ref := itor // 多執行緒需要使用中間變數

		go func() {
			if err := initializeSector(context, ref); err != nil {
				work.Error(fmt.Errorf("initialize failed: %w", err))
			} // if

			work.Increment()
		}()
	} // for

	errs := work.End()

	if len(errs) == 0 {
		if err := initializeStruct(context); err != nil {
			errs = append(errs, err)
		} // if
	} // if

	return errs
}
