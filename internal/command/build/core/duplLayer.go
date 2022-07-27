package core

// DuplLayer 階層重複檢查器
type DuplLayer struct {
	dupls map[string]int // 重複列表
}

// Check 重複檢查, true表示正常, false則否
func (this *DuplLayer) Check(layers ...Layer) bool {
	for _, itor := range layers {
		if type_, ok := this.dupls[itor.Name]; ok {
			return type_ == itor.Type
		} // if

		this.dupls[itor.Name] = itor.Type
	} // for

	return true
}

// NewDuplLayer 建立階層重複檢查器
func NewDuplLayer() *DuplLayer {
	return &DuplLayer{
		dupls: map[string]int{},
	}
}
