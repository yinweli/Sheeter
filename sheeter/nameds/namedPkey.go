package nameds

import (
	"github.com/yinweli/Sheeter/v2/sheeter/fields"
)

// Pkey 主索引命名工具
type Pkey struct {
	Pkey fields.Field // 主索引欄位
}

// PkeyCs 取得cs主索引類型字串
func (this *Pkey) PkeyCs() string {
	return this.Pkey.ToTypeCs()
}

// PkeyGo 取得go主索引類型字串
func (this *Pkey) PkeyGo() string {
	return this.Pkey.ToTypeGo()
}
