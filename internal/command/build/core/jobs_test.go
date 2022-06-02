package core

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestJobs_Execute(t *testing.T) {
	job1 := &testJob{longName: "cmd1", err: fmt.Errorf("err1")}
	job2 := &testJob{longName: "cmd2", err: fmt.Errorf("err2")}
	job3 := &testJob{longName: "cmd3"}
	jobs := NewJobs(&cobra.Command{}, []Job{job1, job2, job3})

	err := jobs.cmd.Flags().Set(job1.LongName(), "true")
	assert.Nil(t, err)
	assert.Equal(t, job1.err, jobs.Execute(nil))

	err = jobs.cmd.Flags().Set(job1.LongName(), "false")
	assert.Nil(t, err)
	err = jobs.cmd.Flags().Set(job2.LongName(), "true")
	assert.Nil(t, err)
	assert.Equal(t, job2.err, jobs.Execute(nil))

	err = jobs.cmd.Flags().Set(job2.LongName(), "false")
	assert.Nil(t, err)
	err = jobs.cmd.Flags().Set(job3.LongName(), "true")
	assert.Nil(t, err)
	assert.Nil(t, jobs.Execute(nil))
}

func TestJobs_Progress(t *testing.T) {
	job1 := &testJob{longName: "cmd1", multi: 1}
	job2 := &testJob{longName: "cmd2", multi: 1}
	job3 := &testJob{longName: "cmd3", multi: 0}
	jobs := NewJobs(&cobra.Command{}, []Job{job1, job2, job3})
	sheetSize := 1

	assert.Equal(t, 0, jobs.Progress(sheetSize))

	err := jobs.cmd.Flags().Set(job1.LongName(), "true")
	assert.Nil(t, err)
	assert.Equal(t, 1, jobs.Progress(sheetSize))

	err = jobs.cmd.Flags().Set(job2.LongName(), "true")
	assert.Nil(t, err)
	assert.Equal(t, 2, jobs.Progress(sheetSize))

	err = jobs.cmd.Flags().Set(job3.LongName(), "true")
	assert.Nil(t, err)
	assert.Equal(t, 2, jobs.Progress(sheetSize))
}

func TestJobs_State(t *testing.T) {
	job1 := &testJob{longName: "cmd"}
	jobs := NewJobs(&cobra.Command{}, []Job{job1})

	err := jobs.cmd.Flags().Set(job1.LongName(), "true")
	assert.Nil(t, err)
	assert.True(t, jobs.state(job1.LongName()))
	assert.False(t, jobs.state("????"))
}

func TestNewJobs(t *testing.T) {
	executor := NewJobs(&cobra.Command{}, []Job{})
	assert.NotNil(t, executor)
}

type testJob struct {
	longName  string
	shortName string
	note      string
	multi     int
	path      string
	err       error
}

func (this *testJob) LongName() string {
	return this.longName
}

func (this *testJob) ShortName() string {
	return this.shortName
}

func (this *testJob) Note() string {
	return this.note
}

func (this *testJob) Progress(sheetSize int) int {
	return sheetSize * this.multi
}

func (this *testJob) Execute(cargo *Cargo) (filePath string, err error) {
	return this.path, this.err
}
