package core

import (
	"fmt"
	"strings"
)

const (
	LayerArray  = iota // 陣列階層
	LayerStruct        // 結構階層
)

const (
	modeBegin = iota // 開始模式
	modeEnd          // 結束模式
)

const tokenArray = "{[]" // 階層字串以'{[]'符號開始, 表示為陣列的開始
const tokenStruct = "{"  // 階層字串以'{'符號開始, 表示為結構的開始
const tokenEnd = "}"     // 階層字串以'}'符號開始, 表示為結構/陣列的結束

// ParseLayer 解析字串為階層, 格式為 {name 或 {[]name 或 }
func ParseLayer(input string) (layer []Layer, back int, err error) { // TODO: layerParse還沒做玩
	tokens := strings.Fields(input)
	mode := modeBegin

	for _, itor := range tokens {
		switch {
		case mode == modeBegin && strings.HasPrefix(itor,
			tokenArray): // 由於tokenStruct可能會被辨別為tokenArray的一種, 所以tokenArray要先判斷
			layer = append(layer, Layer{
				Name: strings.TrimPrefix(itor, tokenArray),
				Type: LayerArray,
			})

		case mode == modeBegin && strings.HasPrefix(itor, tokenStruct):
			layer = append(layer, Layer{
				Name: strings.TrimPrefix(itor, tokenStruct),
				Type: LayerStruct,
			})

		case strings.HasPrefix(itor, tokenEnd):
			mode = modeEnd
			back++

		default:
			return nil, 0, fmt.Errorf("layer format failed: %s", input)
		} // switch
	} // for

	return layer, back, nil
}

// Layer 階層資料
type Layer struct {
	Name string // 階層名稱
	Type int    // 階層類型
}
