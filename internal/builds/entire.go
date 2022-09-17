package builds

import (
	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Entire 全域資料
type Entire struct {
	*layouts.Type
}

// FieldName 取得欄位名稱
func (this *Entire) FieldName(field *layouts.Field) string {
	return utils.FirstUpper(field.Name)
}

// FieldNote 取得欄位註解
func (this *Entire) FieldNote(field *layouts.Field) string {
	return field.Note
}

// FieldTypeCs 取得cs欄位類型
func (this *Entire) FieldTypeCs(field *layouts.Field) string {
	name := ""

	if field.Field != nil {
		name += field.Field.ToTypeCs()
	} else {
		name += field.Alter
	} // if

	if field.Array {
		name += internal.TokenArray
	} // if

	return name
}

// FieldTypeGo 取得go欄位類型
func (this *Entire) FieldTypeGo(field *layouts.Field) string {
	name := ""

	if field.Array {
		name += internal.TokenArray
	} // if

	if field.Field != nil {
		name += field.Field.ToTypeGo()
	} else {
		name += field.Alter
	} // if

	return name
}
