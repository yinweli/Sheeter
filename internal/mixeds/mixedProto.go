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
	return internal.PathProto
}

// PathProtoCs 取得proto-cs路徑
func (this *Proto) PathProtoCs() string {
	return filepath.Join(internal.PathProto, internal.PathCs)
}

// PathProtoGo 取得proto-go路徑
func (this *Proto) PathProtoGo() string {
	return filepath.Join(internal.PathProto, internal.PathGo)
}

// FileProtoSchema 取得proto架構檔名路徑
func (this *Proto) FileProtoSchema() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoSchema,
		path:       []string{internal.PathProto},
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

// FileProtoBat 取得proto-bat檔名路徑
func (this *Proto) FileProtoBat() string {
	return filepath.Join(internal.PathProto, internal.FileProtoBat)
}

// FileProtoSh 取得proto-sh檔名路徑
func (this *Proto) FileProtoSh() string {
	return filepath.Join(internal.PathProto, internal.FileProtoSh)
}

// ProtoDepend 取得proto依賴檔案名稱
func (this *Proto) ProtoDepend(name string) string {
	return name + "." + internal.ExtProtoSchema
}
