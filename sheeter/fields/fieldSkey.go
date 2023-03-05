package fields

import (
	"github.com/yinweli/Sheeter/sheeter"
)

// Skey 主要整數索引
type Skey struct {
}

// Field 取得excel欄位類型列表
func (this *Skey) Field() []string {
	return []string{"skey"}
}

// IsShow 是否顯示
func (this *Skey) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Skey) IsPkey() bool {
	return true
}

// ToTypeCs 取得cs類型字串
func (this *Skey) ToTypeCs() string {
	return sheeter.TokenSkeyCs
}

// ToTypeGo 取得go類型字串
func (this *Skey) ToTypeGo() string {
	return sheeter.TokenSkeyGo
}

// ToTypeProto 取得proto類型字串
func (this *Skey) ToTypeProto() string {
	return sheeter.TokenSkeyProto
}

// ToJsonValue 轉換為json值
func (this *Skey) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil
}
