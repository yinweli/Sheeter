package core

import (
	"github.com/schollz/progressbar/v3"
)

// Cargo 工作箱
type Cargo struct {
	Progress *progressbar.ProgressBar // 進度條
	Global   *Global                  // 全域設定
	Element  *Element                 // 項目設定
	Columns  []*Column                // 行資料列表
}

// Column 行資料
type Column struct {
	Note  string   // 欄位註解
	Name  string   // 欄位名稱
	Field Field    // 欄位類型
	Datas []string // 資料列表
}
