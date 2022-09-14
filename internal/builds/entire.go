package builds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
)

// Entire 全域資料
type Entire struct {
	*layouts.Type
}

// AppName 取得程式名稱
func (this *Entire) AppName() string {
	return internal.AppName
}

// Namespace 取得命名空間名稱
func (this *Entire) Namespace() string {
	return internal.AppName
}

// FileJsonCsCode 取得json-cs程式碼檔名路徑
func (this *Entire) FileJsonCsCode() string {
	name := this.combine(internal.PathJsonCs, this.StructName, internal.ExtCs)
	return name
}

// FileJsonCsReader 取得json-cs讀取器檔名路徑
func (this *Entire) FileJsonCsReader() string {
	name := this.combine(internal.PathJsonCs, this.ReaderName, internal.ExtCs)
	return name
}

// FileJsonGoCode 取得json-go程式碼檔名路徑
func (this *Entire) FileJsonGoCode() string {
	name := this.combine(internal.PathJsonGo, this.StructName, internal.ExtGo)
	return name
}

// FileJsonGoReader 取得json-go讀取器檔名路徑
func (this *Entire) FileJsonGoReader() string {
	name := this.combine(internal.PathJsonGo, this.ReaderName, internal.ExtGo)
	return name
}

// FileJsonCode 取得程式碼可用的json檔名路徑
func (this *Entire) FileJsonCode() string {
	path := filepath.ToSlash(this.FileJson) // 因為要把路徑寫到程式碼中, 所以要改變分隔符號的方式
	return path
}

// combine 取得組合名稱
func (this *Entire) combine(path, name, ext string) string {
	return filepath.Join(path, name+"."+ext)
}
