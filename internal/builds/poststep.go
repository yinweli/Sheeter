package builds

import (
	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Poststep 後製處理
func Poststep(context *Context) []error {
	_, errs := pipelines.Pipeline("poststep", context.Poststep, []pipelines.PipelineFunc{
		PoststepJsonDepotCs,
		PoststepJsonDepotGo,
		PoststepProtoDepotCs,
		PoststepProtoDepotGo,
		PoststepConvertCs,
		PoststepConvertGo,
	})
	return errs
}
