package core

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestJobs(t *testing.T) {
	jobs := mockJobs()
	cmd := &cobra.Command{}

	assert.Equal(t, 3, len(jobs.jobs))
	assert.Equal(t, "job1", jobs.jobs[0].Long())
	assert.Equal(t, "job2", jobs.jobs[1].Long())
	assert.Equal(t, "job3", jobs.jobs[2].Long())

	jobs.Flag(cmd)
	assert.NotNil(t, cmd.Flags().Lookup("job1"))
	assert.NotNil(t, cmd.Flags().Lookup("job2"))
	assert.NotNil(t, cmd.Flags().Lookup("job3"))

	assert.Equal(t, 0, jobs.Calc(cmd, 5))

	_ = cmd.Flags().Set("job1", "true")
	_ = cmd.Flags().Set("job2", "false")
	_ = cmd.Flags().Set("job3", "false")
	assert.Equal(t, 0, jobs.Calc(cmd, 5))

	_ = cmd.Flags().Set("job1", "false")
	_ = cmd.Flags().Set("job2", "true")
	_ = cmd.Flags().Set("job3", "false")
	assert.Equal(t, 5, jobs.Calc(cmd, 5))

	_ = cmd.Flags().Set("job1", "false")
	_ = cmd.Flags().Set("job2", "false")
	_ = cmd.Flags().Set("job3", "true")
	assert.Equal(t, 10, jobs.Calc(cmd, 5))

	_ = cmd.Flags().Set("job1", "true")
	_ = cmd.Flags().Set("job2", "true")
	_ = cmd.Flags().Set("job3", "true")
	assert.Equal(t, 15, jobs.Calc(cmd, 5))

	_ = cmd.Flags().Set("job1", "true")
	_ = cmd.Flags().Set("job2", "false")
	_ = cmd.Flags().Set("job3", "false")
	err := jobs.Execute(cmd, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "err1", err.Error())

	_ = cmd.Flags().Set("job1", "false")
	_ = cmd.Flags().Set("job2", "true")
	_ = cmd.Flags().Set("job3", "false")
	err = jobs.Execute(cmd, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "err2", err.Error())

	_ = cmd.Flags().Set("job1", "false")
	_ = cmd.Flags().Set("job2", "false")
	_ = cmd.Flags().Set("job3", "true")
	err = jobs.Execute(cmd, nil)
	assert.Nil(t, err)
}

func TestNewJobs(t *testing.T) {
	jobs := NewJobs()
	assert.NotNil(t, jobs)
}

type testJob struct {
	long  string
	multi int
	err   error
}

func (this *testJob) Long() string {
	return this.long
}

func (this *testJob) Short() string {
	return ""
}

func (this *testJob) Note() string {
	return ""
}

func (this *testJob) Calc(sheetSize int) int {
	return sheetSize * this.multi
}

func (this *testJob) Execute(cargo *Cargo) (filePath string, err error) {
	return "", this.err
}

func mockJobs() *Jobs {
	jobs := &Jobs{}
	jobs.Add(&testJob{long: "job1", multi: 0, err: fmt.Errorf("err1")})
	jobs.Add(&testJob{long: "job2", multi: 1, err: fmt.Errorf("err2")})
	jobs.Add(&testJob{long: "job3", multi: 2})

	return jobs
}
