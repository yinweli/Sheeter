package core

import "github.com/spf13/cobra"

// Jobs 工作列表
type Jobs struct {
	cmd     *cobra.Command // 命令物件
	jobMaps map[string]Job // 工作列表
}

// Job 工作介面
type Job interface {
	// LongName 取得長名稱
	LongName() string

	// ShortName 取得短名稱
	ShortName() string

	// Note 取得註解
	Note() string

	// Progress 取得進度值
	Progress(sheetSize int) int

	// Execute 執行工作
	Execute(cargo *Cargo) (filePath string, err error)
}

// Execute 執行工作
func (this *Jobs) Execute(cargo *Cargo) error {
	for name, itor := range this.jobMaps {
		if this.state(name) {
			_, err := itor.Execute(cargo)

			if err != nil {
				return err
			} // if
		} // if
	} // for

	return nil
}

// Progress 取得進度值
func (this *Jobs) Progress(sheetSize int) int {
	count := 0

	for name, itor := range this.jobMaps {
		if this.state(name) {
			count = count + itor.Progress(sheetSize)
		} // if
	} // for

	return count
}

// state 取得執行狀態
func (this *Jobs) state(name string) bool {
	result, err := this.cmd.Flags().GetBool(name)

	if err != nil {
		return false
	} // if

	return result
}

// NewJobs 建立工作列表
func NewJobs(cmd *cobra.Command, jobs []Job) *Jobs {
	jobMaps := make(map[string]Job)

	for _, itor := range jobs {
		cmd.Flags().BoolP(itor.LongName(), itor.ShortName(), false, itor.Note())
		jobMaps[itor.LongName()] = itor
	} // for

	return &Jobs{
		cmd:     cmd,
		jobMaps: jobMaps,
	}
}
