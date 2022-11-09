package builds

import (
	"github.com/yinweli/Sheeter/sheeter/pipelines"
)

// Encoding 編碼處理
func Encoding(context *Context) []error {
	_, errs := pipelines.Pipeline("encoding", context.Encoding, []pipelines.PipelineFunc{
		EncodingJson,
		EncodingProto,
	})
	return errs
}
