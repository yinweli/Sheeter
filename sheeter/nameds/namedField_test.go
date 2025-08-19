package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/fields"
	"github.com/yinweli/Sheeter/v3/sheeter/layouts"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteField) SetupSuite() {
	this.Env = testdata.EnvSetup("test-nameds-field")
}

func (this *SuiteField) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteField) TestName() {
	target := &Field{
		Data: &layouts.Data{
			Tag:   "",
			Name:  "name",
			Note:  "no\nte",
			Field: &fields.Pkey{},
		},
	}
	assert.Equal(this.T(), "Name", target.FieldName())
	assert.Equal(this.T(), "no. te", target.FieldNote())
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.FieldTypeCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.FieldTypeGo())
}
