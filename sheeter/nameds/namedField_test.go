package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/fields"
	"github.com/yinweli/Sheeter/v2/sheeter/layouts"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteField) SetupSuite() {
	this.TBegin("test-nameds-field", "")
}

func (this *SuiteField) TearDownSuite() {
	this.TFinal()
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
	assert.Equal(this.T(), "note", target.FieldNote())
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.FieldTypeCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.FieldTypeGo())
}
