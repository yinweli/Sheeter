package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// OutputJsonSchema 輸出json架構
func OutputJsonSchema(content *Content) error {
	packs, _, err := content.builder.Pack([]string{}, true)

	if err != nil {
		return fmt.Errorf("%s: output json schema failed: %w", content.StructName(), err)
	} // if

	if err = util.WriteJson(content.FileJsonSchema(), packs, content.Bom); err != nil {
		return fmt.Errorf("%s: output json schema failed: %w", content.StructName(), err)
	} // if

	return nil
}
