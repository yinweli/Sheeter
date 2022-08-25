package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestParser(t *testing.T) {
	suite.Run(t, new(SuiteParser))
}

type SuiteParser struct {
	suite.Suite
}

func (this *SuiteParser) TestParser() {
	name, field, err := Parser("real#bool")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "real", name)
	assert.Equal(this.T(), (&Bool{}).Type(), field.Type())

	_, _, err = Parser("fa-ke#fake")
	assert.NotNil(this.T(), err)

	_, _, err = Parser("fake#fake")
	assert.NotNil(this.T(), err)

	_, _, err = Parser("fake#fake#fake")
	assert.NotNil(this.T(), err)

	_, _, err = Parser(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
