package core

import (
	"github.com/spf13/cobra"
)

// Executor 執行者
type Executor struct {
	Cmd      *cobra.Command      // 命令物件
	execMaps map[string]ExecData // 執行資料列表
}

// ExecData 執行資料
type ExecData struct {
	LongName  string                                          // 長名稱
	ShortName string                                          // 短名稱
	Note      string                                          // 註解
	ExecFunc  func(cargo *Cargo) (filePath string, err error) // 執行函式
}

// Run 執行項目
func (this *Executor) Run(cargo *Cargo) error {
	for name, data := range this.execMaps {
		if this.state(name) {
			_, err := data.ExecFunc(cargo)

			if err != nil {
				return err
			} // if
		} // if
	} // for

	return nil
}

// Count 取得執行數量
func (this *Executor) Count() int {
	count := 0

	for name, _ := range this.execMaps {
		if this.state(name) {
			count++
		} // if
	} // for

	return count
}

// state 取得執行狀態
func (this *Executor) state(name string) bool {
	result, err := this.Cmd.Flags().GetBool(name)

	if err != nil {
		return false
	} // if

	return result
}

// NewExecutor 建立執行者
func NewExecutor(cmd *cobra.Command, execDatas []ExecData) *Executor {
	execMaps := make(map[string]ExecData)

	for _, itor := range execDatas {
		cmd.Flags().BoolP(itor.LongName, itor.ShortName, false, itor.Note)
		execMaps[itor.LongName] = itor
	} // for

	return &Executor{
		Cmd:      cmd,
		execMaps: execMaps,
	}
}
