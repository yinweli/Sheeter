package build

import (
	"fmt"
	"os"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := NewCommand()

	assert.NotNil(t, command, "command nil")
}

func TestExecute(t *testing.T) {
	buffer, command := testdata.MockCommand()

	fmt.Println(os.Getwd())

	execute(command, []string{testdata.Path(testdata.RealYaml), ""})
	assert.Equal(t, "", buffer.String(), "read real config failed")

	execute(command, []string{testdata.Path(testdata.FakeYaml), ""})
	assert.NotEqual(t, "", buffer.String(), "read fake config failed")

	execute(command, []string{testdata.Path(testdata.DefectYml), ""})
	assert.NotEqual(t, "", buffer.String(), "read defect config failed")
}
