package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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
		Layout: &layouts.LayoutData{
			Tag:   "",
			Name:  "name",
			Note:  "note",
			Field: &fields.Int{},
		},
	}
	assert.Equal(this.T(), "Name", target.FieldName())
	assert.Equal(this.T(), "note", target.FieldNote())
	assert.Equal(this.T(), "Int32", target.FieldTypeCs())
	assert.Equal(this.T(), "int32", target.FieldTypeGo())
}
