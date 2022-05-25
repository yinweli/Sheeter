package version

import (
	"fmt"
	"testing"

	"Sheeter/internal"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	buffer, command := testdata.FakeCommand()

	execute(command, []string{})
	assert.Equal(t, fmt.Sprintf("%s %s", internal.Title, internal.Version), buffer.String(), "execute failed")
}
