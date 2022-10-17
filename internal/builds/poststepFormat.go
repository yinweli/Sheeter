package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/utils"
)

// poststepFormat 後製格式化程式碼
func poststepFormat(data *poststepData) error {
	if err := utils.Format(); err != nil {
		return fmt.Errorf("poststep format failed: %w", err)
	} // if

	return nil
}
