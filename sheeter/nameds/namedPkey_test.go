package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/fields"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestPkey(t *testing.T) {
	suite.Run(t, new(SuitePkey))
}

type SuitePkey struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuitePkey) SetupSuite() {
	this.TBegin("test-nameds-pkey", "")
}

func (this *SuitePkey) TearDownSuite() {
	this.TFinal()
}

func (this *SuitePkey) TestPkey() {
	target := &Pkey{
		Pkey: &fields.Pkey{},
	}
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.PkeyCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.PkeyGo())
}
