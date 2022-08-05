package build

// duplLayer 階層重複檢查器
type duplLayer struct {
	datas map[string]int // 資料列表
}

// Check 重複檢查
func (this *duplLayer) Check(layers ...Layer) bool {
	for _, itor := range layers {
		if type_, ok := this.datas[itor.Name]; ok {
			return type_ == itor.Type
		} // if

		this.datas[itor.Name] = itor.Type
	} // for

	return true
}

// NewDuplLayer 建立階層重複檢查器
func NewDuplLayer() *duplLayer {
	return &duplLayer{
		datas: map[string]int{},
	}
}
