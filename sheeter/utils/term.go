package utils

import (
	"strings"

	"github.com/yinweli/Sheeter/v3/sheeter"
)

// MergeTerm 合併名稱
type MergeTerm string

// Name 取得合併名稱
func (this MergeTerm) Name() string {
	if name, _, ok := strings.Cut(string(this), sheeter.TokenName); ok {
		return name
	} // if

	return ""
}

// Member 取得合併成員
func (this MergeTerm) Member() []SheetTerm {
	result := []SheetTerm{}

	if _, member, ok := strings.Cut(string(this), sheeter.TokenName); ok {
		for _, itor := range strings.Split(member, sheeter.TokenTerm) {
			result = append(result, SheetTerm(itor))
		} // for
	} // if

	return result
}

// SheetTerm 表格名稱
type SheetTerm string

// Name 取得表格名稱
func (this SheetTerm) Name() (excel, sheet string) {
	if e, s, ok := strings.Cut(string(this), sheeter.TokenExcel); ok {
		return e, s
	} // if

	return "", ""
}

// Match 名稱是否匹配
func (this SheetTerm) Match(excel, sheet string) bool {
	if e, s, ok := strings.Cut(string(this), sheeter.TokenExcel); ok {
		if strings.EqualFold(e, excel) && strings.EqualFold(s, sheet) {
			return true
		} // if
	} // if

	return false
}
