package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestLong(t *testing.T) {
	suite.Run(t, new(SuiteLong))
}

type SuiteLong struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteLong) SetupSuite() {
	this.Change("test-field-long")
}

func (this *SuiteLong) TearDownSuite() {
	this.Restore()
}

func (this *SuiteLong) target() *Long {
	return &Long{}
}

func (this *SuiteLong) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"long"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenLongCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenLongGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenLongProto, target.ToTypeProto())
}

func (this *SuiteLong) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
