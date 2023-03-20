package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestLong(t *testing.T) {
	suite.Run(t, new(SuiteLong))
}

type SuiteLong struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteLong) SetupSuite() {
	this.TBegin("test-fields-long", "")
}

func (this *SuiteLong) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteLong) TestField() {
	target := &Long{}
	assert.Equal(this.T(), []string{"long"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.IsType(this.T(), &Lkey{}, target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeLongCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeLongGo, target.ToTypeGo())
}

func (this *SuiteLong) TestToJsonValue() {
	target := &Long{}

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(0), result)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}
