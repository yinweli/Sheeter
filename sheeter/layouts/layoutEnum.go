package layouts

import (
	"fmt"
	"strconv"

	"github.com/yinweli/Sheeter/sheeter/utils"
)

// NewLayoutEnum 建立列舉布局器
func NewLayoutEnum() *LayoutEnum {
	return &LayoutEnum{}
}

// LayoutEnum 列舉布局器
type LayoutEnum struct {
	enums []*Enum // 列舉列表
}

// Enum 列舉資料
type Enum struct {
	Name    string // 列舉名稱
	Index   int    // 列舉編號
	Comment string // 列舉說明
}

// Add 新增列舉
func (this *LayoutEnum) Add(data []string) error {
	name := utils.GetItem(data, 0)

	if utils.NameCheck(name) == false {
		return fmt.Errorf("layoutEnum add failed: invalid name")
	} // if

	index, err := strconv.Atoi(utils.GetItem(data, 1))

	if err != nil {
		return fmt.Errorf("layoutEnum add failed: invalid index")
	} // if

	for _, itor := range this.enums {
		if itor.Name == name {
			return fmt.Errorf("layoutEnum add failed: name duplicate")
		} // if

		if itor.Index == index {
			return fmt.Errorf("layoutEnum add failed: index duplicate")
		} // if
	} // for

	this.enums = append(this.enums, &Enum{
		Name:    name,
		Index:   index,
		Comment: utils.GetItem(data, 2),
	})
	return nil
}

// Enums 取得列舉列表
func (this *LayoutEnum) Enums() []*Enum {
	return this.enums
}
