package layers

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
)

const (
	LayerArray   = iota // 陣列階層
	LayerStruct         // 結構階層
	LayerDivider        // 陣列隔線
)

const (
	modeBegin = iota // 開始模式
	modeEnd          // 結束模式
)

const tokenArray = "{[]" // 以'{[]'符號開始, 表示為陣列的開始
const tokenStruct = "{"  // 以'{'符號開始, 表示為結構的開始
const tokenDivider = "/" // 以'/'符號開始, 表示為陣列的分隔, 必須在階層字串的最開始, 並且只能出現一次
const tokenEnd = "}"     // 以'}'符號開始, 表示為結構/陣列的結束

// ParseLayer 解析字串為階層, 格式為'{[]name'或'/'或'{name'或'}', 以空格分隔
func ParseLayer(input string) (layers []Layer, back int, err error) {
	tokens := strings.Fields(input)
	mode := modeBegin
	divider := false

	for i, itor := range tokens {
		if mode == modeBegin && strings.HasPrefix(itor, tokenArray) { // tokenArray要先判斷, 不然會有錯誤
			if name := strings.TrimPrefix(itor, tokenArray); util.VariableCheck(name) {
				layers = append(layers, Layer{
					Name: name,
					Type: LayerArray,
				})
				continue
			} // if
		} // if

		if mode == modeBegin && strings.HasPrefix(itor, tokenStruct) {
			if name := strings.TrimPrefix(itor, tokenStruct); util.VariableCheck(name) {
				layers = append(layers, Layer{
					Name: name,
					Type: LayerStruct,
				})
				continue
			} // if
		} // if

		if mode == modeBegin && strings.HasPrefix(itor, tokenDivider) && divider == false && i == 0 {
			if name := strings.TrimPrefix(itor, tokenDivider); name == "" {
				layers = append(layers, Layer{
					Type: LayerDivider,
				})
				divider = true
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

	for _, itor := range layers {
		if itor.Name == "" && itor.Type != LayerDivider {
			goto failed
		} // if
	} // for

	return layers, back, nil

failed:
	return nil, 0, fmt.Errorf("layer format failed: %s", input)
}

// Layer 階層資料
type Layer struct {
	Name string // 階層名稱
	Type int    // 階層類型
}
