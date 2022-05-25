package cargo

import (
	"Sheeter/internal/command/build/config"
	"Sheeter/internal/command/build/field"

	"github.com/schollz/progressbar/v3"
)

// Cargo 執行資料
type Cargo struct {
	Progress *progressbar.ProgressBar // 進度條
	Global   *config.Global           // 全域設定
	Element  *config.Element          // 項目設定
	Fields   []field.Field            // 欄位列表
	Columns  []Column                 // 行資料列表
}

// Column 行資料
type Column struct {
	Note  string       // 欄位註解
	Name  string       // 欄位名稱
	Field *field.Field // 欄位類型
	Datas []string     // 資料列表
}
