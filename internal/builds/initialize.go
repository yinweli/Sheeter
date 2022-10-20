package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Initialize 初始化處理
func Initialize(config *Config) (context *Context, errs []error) {
	executor := []pipelines.Executor{
		InitializeElement,
		InitializeEnum,
	}
	context = InitializeContext(config)

	if errs = pipelines.Execute("initialize ", context.Element, executor); len(errs) > 0 {
		return nil, errs
	} // if

	if err := InitializePick(context); err != nil {
		return nil, []error{fmt.Errorf("build failed: %w", err)}
	} // if

	return context, errs
}
