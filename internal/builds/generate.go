package builds

import (
	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Generate 產生處理
func Generate(context *Context) []error {
	return pipelines.Execute("generate ", context.Generate, []pipelines.Executor{
		GenerateJsonStructCs,
		GenerateJsonReaderCs,
		GenerateJsonStructGo,
		GenerateJsonReaderGo,
		GenerateProtoSchema,
		GenerateProtoReaderCs,
		GenerateProtoReaderGo,
	})
}
