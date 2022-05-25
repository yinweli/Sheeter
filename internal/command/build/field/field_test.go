package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFields(t *testing.T) {
	assert.NotNil(t, NewFields(), "new fields failed")
}
