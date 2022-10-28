package builds

import (
	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Generate 產生處理
func Generate(context *Context) []error {
	_, errs := pipelines.Pipeline("generate", context.Generate, []pipelines.PipelineFunc{
		GenerateJsonStructCs,
		GenerateJsonReaderCs,
		GenerateJsonStructGo,
		GenerateJsonReaderGo,
		GenerateProtoSchema,
		GenerateProtoReaderCs,
		GenerateProtoReaderGo,
		GenerateEnumSchema,
	})
	return errs
}
