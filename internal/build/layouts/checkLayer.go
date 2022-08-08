package layouts

import "github.com/yinweli/Sheeter/internal/build/layers"

// checkLayer 階層檢查器
type checkLayer map[string]int

// check 階層檢查
func (this *checkLayer) check(layer ...layers.Layer) bool {
	for _, itor := range layer {
		if type_, ok := (*this)[itor.Name]; ok {
			return type_ == itor.Type
		} // if

		(*this)[itor.Name] = itor.Type
	} // for

	return true
}
