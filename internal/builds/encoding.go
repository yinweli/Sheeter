package builds

import (
	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Encoding 編碼處理
func Encoding(context *Context) []error {
	return pipelines.Execute("encoding", context.Encoding, []pipelines.Executor{
		EncodingJson,
		EncodingProto,
	})
}
