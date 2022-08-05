package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDuplField(t *testing.T) {
	suite.Run(t, new(SuiteDuplField))
}

type SuiteDuplField struct {
	suite.Suite
	item1 string
	item2 string
	item3 string
}

func (this *SuiteDuplField) SetupSuite() {
	this.item1 = "001"
	this.item2 = "001/002"
	this.item3 = "001/002/003"
}

func (this *SuiteDuplField) target() *duplField {
	return &duplField{
		datas: map[string]bool{},
	}
}

func (this *SuiteDuplField) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.Check(this.item1))
	assert.True(this.T(), target.Check(this.item2))
	assert.True(this.T(), target.Check(this.item3))
	assert.False(this.T(), target.Check(this.item1))
	assert.False(this.T(), target.Check(this.item2))
	assert.False(this.T(), target.Check(this.item3))
}

func (this *SuiteDuplField) TestNewDuplField() {
	assert.NotNil(this.T(), NewDuplField())
}
