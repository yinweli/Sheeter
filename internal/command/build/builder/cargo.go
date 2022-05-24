package builder

import (
	"io"

	"Sheeter/internal/command/build/config"
	"Sheeter/internal/command/build/field"

	"github.com/schollz/progressbar/v3"
)

// Cargo 執行資料
type Cargo struct {
	Output   io.Writer                // 輸出物件
	Global   *config.Global           // 全域設定
	Element  *config.Element          // 項目設定
	Progress *progressbar.ProgressBar // 進度條
	Columns  []Column                 // 欄位列表
}

// Column 欄位資料
type Column struct {
	Note  string       // 欄位註解
	Name  string       // 欄位名稱
	Field *field.Field // 欄位類型
	Datas []string     // 資料列表
}
