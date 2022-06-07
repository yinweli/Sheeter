package core

import "github.com/spf13/cobra"

// Jobs 工作列表
type Jobs struct {
	jobs []Job // 工作列表
}

// Job 工作介面
type Job interface {
	// Long 取得長名稱
	Long() string

	// Short 取得短名稱
	Short() string

	// Note 取得註解
	Note() string

	// Calc 計算進度值
	Calc(sheetSize int) int

	// Execute 執行工作
	Execute(cargo *Cargo) (filePath string, err error)
}

// Add 新增工作
func (this *Jobs) Add(job Job) {
	this.jobs = append(this.jobs, job)
}

// Flag 設定命令旗標
func (this *Jobs) Flag(cmd *cobra.Command) {
	for _, itor := range this.jobs {
		cmd.Flags().BoolP(itor.Long(), itor.Short(), false, itor.Note())
	} // for
}

// Calc 計算進度值
func (this *Jobs) Calc(cmd *cobra.Command, sheetSize int) int {
	value := 0

	for _, itor := range this.jobs {
		result, err := cmd.Flags().GetBool(itor.Long())

		if result && err == nil {
			value = value + itor.Calc(sheetSize)
		} // if
	} // for

	return value
}

// Execute 執行工作
func (this *Jobs) Execute(cmd *cobra.Command, cargo *Cargo) error {
	for _, itor := range this.jobs {
		result, err := cmd.Flags().GetBool(itor.Long())

		if result && err == nil {
			_, err := itor.Execute(cargo)

			if err != nil {
				return err
			} // if
		} // if
	} // for

	return nil
}

// NewJobs 建立工作列表
func NewJobs() *Jobs {
	jobs := &Jobs{}
	jobs.Add(&WriteJson{})
	jobs.Add(&WriteCpp{})
	jobs.Add(&WriteCs{})
	jobs.Add(&WriteGo{})

	return jobs
}
