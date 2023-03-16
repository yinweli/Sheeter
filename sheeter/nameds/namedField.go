package nameds

import (
	"strings"

	"github.com/yinweli/Sheeter/v2/sheeter/layouts"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Field 欄位命名工具
type Field struct {
	Data *layouts.Data // 布局資料
}

// FieldName 取得欄位名稱
func (this *Field) FieldName() string {
	return utils.FirstUpper(this.Data.Name)
}

// FieldNote 取得欄位註解
func (this *Field) FieldNote() string {
	return strings.ReplaceAll(this.Data.Note, "\n", "") // 避免換行造成產生的程式碼錯誤
}

// FieldTypeCs 取得cs欄位類型
func (this *Field) FieldTypeCs() string {
	return this.Data.Field.ToTypeCs()
}

// FieldTypeGo 取得go欄位類型
func (this *Field) FieldTypeGo() string {
	return this.Data.Field.ToTypeGo()
}
