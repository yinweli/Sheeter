package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
}

func (this *SuiteField) TestParseField() {
	name, field, err := ParseField("real#bool")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "real", name)
	assert.Equal(this.T(), (&FieldBool{}).Type(), field.Type())

	_, _, err = ParseField("fa-ke#fake")
	assert.NotNil(this.T(), err)

	_, _, err = ParseField("fake#fake")
	assert.NotNil(this.T(), err)

	_, _, err = ParseField("fake#fake#fake")
	assert.NotNil(this.T(), err)

	_, _, err = ParseField(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
