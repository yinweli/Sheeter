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
	return internal.PathSchema
}

// PathProtoCs 取得proto-cs路徑
func (this *Proto) PathProtoCs() string {
	return internal.PathCs
}

// PathProtoGo 取得proto-go路徑
func (this *Proto) PathProtoGo() string {
	return internal.PathGo
}

// FileProtoSchema 取得proto架構檔名路徑
func (this *Proto) FileProtoSchema() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoSchema,
		path:       []string{internal.PathProto, internal.PathSchema},
	})
}

// FileProtoSchemaRelative 取得proto架構檔名相對路徑
func (this *Proto) FileProtoSchemaRelative() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoSchema,
		path:       []string{internal.PathSchema},
	})
}

// FileProtoCsReader 取得proto-cs讀取器程式碼檔名路徑
func (this *Proto) FileProtoCsReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
		path:       []string{internal.PathProto, internal.PathCs},
	})
}

// FileProtoGoReader 取得proto-go讀取器程式碼檔名路徑
func (this *Proto) FileProtoGoReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
		path:       []string{internal.PathProto, internal.PathGo},
	})
}

// FileProtoDataName 取得proto資料檔案名稱
func (this *Proto) FileProtoDataName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoData,
	})
}

// FileProtoDataPath 取得proto資料檔名路徑
func (this *Proto) FileProtoDataPath() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoData,
		path:       []string{internal.PathProto, internal.PathData},
	})
}

// FileProtoCsBat 取得proto-cs-bat檔名路徑
func (this *Proto) FileProtoCsBat() string {
	return filepath.Join(internal.PathProto, internal.FileProtoCsBat)
}

// FileProtoCsSh 取得proto-cs-sh檔名路徑
func (this *Proto) FileProtoCsSh() string {
	return filepath.Join(internal.PathProto, internal.FileProtoCsSh)
}

// FileProtoGoBat 取得proto-go-bat檔名路徑
func (this *Proto) FileProtoGoBat() string {
	return filepath.Join(internal.PathProto, internal.FileProtoGoBat)
}

// FileProtoGoSh 取得proto-go-sh檔名路徑
func (this *Proto) FileProtoGoSh() string {
	return filepath.Join(internal.PathProto, internal.FileProtoGoSh)
}

// ProtoDepend 取得proto依賴檔案名稱
func (this *Proto) ProtoDepend(name string) string {
	return name + "." + internal.ExtProtoSchema
}
