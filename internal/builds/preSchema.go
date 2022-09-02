package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// OutputJsonSchema 輸出json架構
func OutputJsonSchema(content *Content) error {
	packs, _, err := content.builder.Pack([]string{}, true)

	if err != nil {
		return fmt.Errorf("%s: output json schema failed: %w", content.ShowName(), err)
	} // if

	if err = util.JsonWrite(content.SchemaPath(), packs, content.Bom); err != nil { // TODO: json schema的副檔名還是要用.json
		return fmt.Errorf("%s: output json schema failed: %w", content.ShowName(), err)
	} // if

	return nil
}
