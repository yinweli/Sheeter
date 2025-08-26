package sheeter

const (
	BarWidth     = 40        // 進度條寬度
	MaxExcel     = 999999    // 最大開啟excel數量
	MaxSheet     = 999999    // 最大開啟sheet數量
	Application  = "sheeter" // 程式名稱
	Namespace    = "sheeter" // 命名空間名稱
	Reader       = "Reader"  // 讀取程式碼名稱, 為了程式碼組合方便, 這裡故意寫成大寫駝峰
	Sheeter      = "sheeter" // 表格程式碼名稱
	Helper       = "helper"  // 工具程式碼名稱
	PathCs       = "codeCs"  // 輸出路徑: cs
	PathGo       = "codeGo"  // 輸出路徑: go
	PathJson     = "json"    // 輸出路徑: json
	ExtExcel     = ".xlsx"   // 副檔名: excel
	ExtCs        = ".cs"     // 副檔名: cs
	ExtGo        = ".go"     // 副檔名: go
	ExtJson      = ".json"   // 副檔名: json
	TokenIgnore  = "ignore"  // 忽略符號, 當輸出或是標籤為此名稱時不輸出
	TokenArray   = ","       // 陣列分割符號, 陣列以','符號分割元素
	TokenName    = "$"       // 名稱分割符號, 以'$'來分隔名稱與後續項目
	TokenTerm    = "&"       // 項目分割符號, 以'&'來分隔項目
	TokenExcel   = "#"       // 表格分割符號, 以'#'來分隔項目中的excel與sheet名稱
	JsonPrefix   = ""        // json前綴字串
	JsonIdent    = "    "    // json縮排字串
	IndexOutput  = 0         // 輸出欄位置編號
	IndexPrimary = 1         // 主索引位置編號
)
