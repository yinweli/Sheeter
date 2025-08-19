package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/fields"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestPkey(t *testing.T) {
	suite.Run(t, new(SuitePkey))
}

type SuitePkey struct {
	suite.Suite
	testdata.Env
}

func (this *SuitePkey) SetupSuite() {
	this.Env = testdata.EnvSetup("test-nameds-pkey")
}

func (this *SuitePkey) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuitePkey) TestPkey() {
	target := &Pkey{
		Pkey: &fields.Pkey{},
	}
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.PkeyCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.PkeyGo())
}
