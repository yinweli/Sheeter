package layers

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Layer 階層資料
type Layer struct {
	Name string // 階層名稱
	Type int    // 階層類型
}

const (
	LayerArray   = iota // 陣列階層
	LayerStruct         // 結構階層
	LayerDivider        // 陣列隔線
)

const tokenIgnore = "ignore" // 表示為停用階層
const tokenArray = "{[]"     // 以'{[]'符號開始, 表示為陣列的開始
const tokenStruct = "{"      // 以'{'符號開始, 表示為結構的開始
const tokenDivider = "/"     // 以'/'符號開始, 表示為陣列的分隔, 必須在階層字串的最開始, 並且只能出現一次
const tokenEnd = "}"         // 以'}'符號開始, 表示為結構/陣列的結束

// Parser 階層解析, 格式為'{[]name', '{name', '/', '}', 以空格分隔
func Parser(input string) (layers []Layer, back int, err error) {
	tokens := strings.Fields(input)

	if len(tokens) == 0 { // 停用階層
		return layers, back, nil
	} // if

	if strings.EqualFold(tokens[0], tokenIgnore) { // 停用階層
		return layers, back, nil
	} // if

	locate := false  // 位置旗標, false表示前置位置, true表示後置位置
	divider := false // 陣列隔線旗標, false表示沒有出現, true表示已經出現

	for i, itor := range tokens {
		if locate == false && strings.HasPrefix(itor, tokenArray) { // tokenArray要比tokenStruct先判斷, 不然會有錯誤
			if name := strings.TrimPrefix(itor, tokenArray); utils.NameCheck(name) {
				layers = append(layers, Layer{
					Name: name,
					Type: LayerArray,
				})
				continue
			} // if
		} // if

		if locate == false && strings.HasPrefix(itor, tokenStruct) {
			if name := strings.TrimPrefix(itor, tokenStruct); utils.NameCheck(name) {
				layers = append(layers, Layer{
					Name: name,
					Type: LayerStruct,
				})
				continue
			} // if
		} // if

		if locate == false && strings.HasPrefix(itor, tokenDivider) && divider == false && i == 0 {
			if name := strings.TrimPrefix(itor, tokenDivider); name == "" {
				layers = append(layers, Layer{
					Type: LayerDivider,
				})
				divider = true
				continue
			} // if
		} // if

		if strings.HasPrefix(itor, tokenEnd) && utils.AllSame(itor) {
			locate = true
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
	return nil, 0, fmt.Errorf("%s: parser layer failed: invalid format", input)
}
