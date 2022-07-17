package core

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
)

const (
	LayerStruct = iota // 結構階層
	LayerArray         // 陣列階層
)

// ParseLayer 解析字串為階層, 格式為 {name 或 {[]name 或 }
func ParseLayer(input string) (layer []Layer, back int, err error) {
	const separateArray = "{[]" // 階層字串以'{[]'符號開始, 表示為陣列的開始
	const separateStruct = "{"  // 階層字串以'{'符號開始, 表示為結構的開始
	const separateEnd = "}"     // 階層字串以'}'符號開始, 表示為結構/陣列的結束

	tokens := strings.Fields(input)
	forward := true

	for _, itor := range tokens {
		if forward {
			switch {
			case strings.HasPrefix(itor, separateArray):
				layer = append(layer, Layer{
					Name: strings.TrimPrefix(itor, separateArray),
					Type: LayerArray,
				})

			case strings.HasPrefix(itor, separateStruct):
				layer = append(layer, Layer{
					Name: strings.TrimPrefix(itor, separateStruct),
					Type: LayerStruct,
				})

			case strings.HasPrefix(itor, separateEnd):
				forward = false
				back++

			default:
				return nil, 0, fmt.Errorf("layer format failed: %s", input)
			} // switch
		} else {

		} // if
	} // for

	before, after, ok := strings.Cut(input, "#") // 欄位字串以'#'符號分割為名稱與欄位

	if ok == false {
		return "", nil, fmt.Errorf("field format failed: %s", input)
	} // if

	if util.VariableCheck(before) == false {
		return "", nil, fmt.Errorf("name invalid: %s", input)
	} // if

	for _, itor := range fields {
		if itor.Type() == after {
			return before, itor, nil
		} // if
	} // for

	return "", nil, fmt.Errorf("field not found: %s", input)
}

// Layer 階層資料
type Layer struct {
	Name string // 階層名稱
	Type int    // 階層類型
}
