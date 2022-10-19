package builds

import (
	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Poststep 後製處理
func Poststep(context *Context) []error {
	return pipelines.Execute("poststep ", context.Poststep, []pipelines.Executor{
		PoststepJsonDepotCs,
		PoststepJsonDepotGo,
		PoststepProtoDepotCs,
		PoststepProtoDepotGo,
		PoststepConvertCs,
		PoststepConvertGo,
		PoststepFormatCs,
		PoststepFormatGo,
		PoststepFormatProto,
	})
}
