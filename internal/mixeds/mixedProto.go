package mixeds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
)

// Proto proto綜合工具
type Proto struct {
	mixed *Mixed // 綜合工具
}

// PathProtoSchema 取得proto架構路徑
func (this *Proto) PathProtoSchema() string {
	return internal.PathProtoSchema
}

// PathProtoCs 取得proto-cs路徑
func (this *Proto) PathProtoCs() string {
	return internal.PathProtoCs
}

// PathProtoGo 取得proto-go路徑
func (this *Proto) PathProtoGo() string {
	return internal.PathProtoGo
}

// FileProtoDepend 取得proto依賴檔案名稱
func (this *Proto) FileProtoDepend(name string) string {
	return name + "." + internal.ExtProtoSchema
}

// FileProtoSchema 取得proto架構檔名路徑
func (this *Proto) FileProtoSchema() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoSchema,
		path:       internal.PathProtoSchema,
	})
}

// FileProtoCsReader 取得proto-cs讀取器程式碼檔名路徑
func (this *Proto) FileProtoCsReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
		path:       internal.PathProtoCs,
	})
}

// FileProtoGoReader 取得proto-go讀取器程式碼檔名路徑
func (this *Proto) FileProtoGoReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
		path:       internal.PathProtoGo,
	})
}

// FileProtoData 取得proto資料檔名路徑
func (this *Proto) FileProtoData() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoData,
		path:       internal.PathProtoData,
	})
}

// FileProtoDataCode 取得程式碼可用的proto資料檔名路徑
func (this *Proto) FileProtoDataCode() string {
	return filepath.ToSlash(this.FileProtoData()) // 因為要把路徑寫到程式碼中, 所以要改變分隔符號的方式
}

// FileProtoBat 取得proto-bat檔名路徑
func (this *Proto) FileProtoBat() string {
	return internal.PathProtoSchema + "." + internal.ExtBat
}

// FileProtoSh 取得proto-sh檔名路徑
func (this *Proto) FileProtoSh() string {
	return internal.PathProtoSchema + "." + internal.ExtSh
}
