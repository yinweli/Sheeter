package core

import (
	"fmt"
	"strings"
)

const (
	LayerStruct = iota // 結構階層
	LayerArray         // 陣列階層
)

const tokenArray = "{[]" // 階層字串以'{[]'符號開始, 表示為陣列的開始
const tokenStruct = "{"  // 階層字串以'{'符號開始, 表示為結構的開始
const tokenEnd = "}"     // 階層字串以'}'符號開始, 表示為結構/陣列的結束

// ParseLayer 解析字串為階層, 格式為 {name 或 {[]name 或 }
func ParseLayer(input string) (layer []Layer, back int, err error) { // TODO: layerParse還沒做玩
	tokens := strings.Fields(input)
	mode := true // 階層模式, true表示是開始模式, false是結束模式; 結束模式下就不可以碰到陣列開始/結構開始, 會視為錯誤

	for _, itor := range tokens {
		switch {
		case mode && strings.HasPrefix(itor, tokenArray):
			layer = append(layer, Layer{
				Name: strings.TrimPrefix(itor, tokenArray),
				Type: LayerArray,
			})

		case mode && strings.HasPrefix(itor, tokenStruct):
			layer = append(layer, Layer{
				Name: strings.TrimPrefix(itor, tokenStruct),
				Type: LayerStruct,
			})

		case strings.HasPrefix(itor, tokenEnd):
			mode = false
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
