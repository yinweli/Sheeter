package core

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
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

// ParseLayer 解析字串為階層, 格式為'{[]name'或'{name'或'}', 以空格分隔
func ParseLayer(input string) (layer []Layer, back int, err error) {
	tokens := strings.Fields(input)
	mode := modeBegin

	for _, itor := range tokens {
		if mode == modeBegin && strings.HasPrefix(itor, tokenArray) { // tokenArray要先判斷, 不然會有錯誤
			if name := strings.TrimPrefix(itor, tokenArray); util.VariableCheck(name) {
				layer = append(layer, Layer{
					Name: name,
					Type: LayerArray,
				})
				continue
			} // if
		} // if

		if mode == modeBegin && strings.HasPrefix(itor, tokenStruct) {
			if name := strings.TrimPrefix(itor, tokenStruct); util.VariableCheck(name) {
				layer = append(layer, Layer{
					Name: name,
					Type: LayerStruct,
				})
				continue
			} // if
		} // if

		if strings.HasPrefix(itor, tokenEnd) && util.AllSame(itor) {
			mode = modeEnd
			back += len(itor)
			continue
		} // if

		goto failed
	} // for

	for _, itor := range layer {
		if itor.Name == "" {
			goto failed
		} // if
	} // for

	return layer, back, nil

failed:
	return nil, 0, fmt.Errorf("layer format failed: %s", input)
}

// Layer 階層資料
type Layer struct {
	Name string // 階層名稱
	Type int    // 階層類型
}
