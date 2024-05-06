package nameds

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Merge 合併命名工具
type Merge struct {
	Name   string   // 結構名稱
	Member []*Named // 成員列表
}

// StructName 取得結構名稱
func (this *Merge) StructName() string {
	return utils.FirstUpper(this.Name)
}

// StructNote 取得結構說明
func (this *Merge) StructNote() string {
	maximum := len(this.Member) - 1
	builder := strings.Builder{}
	builder.WriteString("merge by ")

	for i, itor := range this.Member {
		builder.WriteString(fmt.Sprintf("%v#%v", itor.ExcelName, itor.SheetName))

		if i < maximum {
			builder.WriteString(", ")
		} // if
	} // for

	return builder.String()
}

// ReaderName 取得讀取器名稱
func (this *Merge) ReaderName() string {
	return this.Member[0].ReaderName()
}

// MemberName 取得成員名稱
func (this *Merge) MemberName() []string {
	result := []string{}

	for _, itor := range this.Member {
		result = append(result, itor.StructName())
	} // for

	return result
}
