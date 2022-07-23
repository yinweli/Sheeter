package util

import (
	"testing"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDuplicate(t *testing.T) {
	suite.Run(t, new(SuiteDuplicate))
}

type SuiteDuplicate struct {
	suite.Suite
	item1 string
	item2 string
	item3 string
}

func (this *SuiteDuplicate) SetupSuite() {
	this.item1 = "001"
	this.item2 = "001/002"
	this.item3 = "001/002/003"
}

func (this *SuiteDuplicate) target() *Duplicate {
	return &Duplicate{
		datas: hashset.New(),
	}
}

func (this *SuiteDuplicate) TestDuplicate() {
	target := this.target()

	assert.True(this.T(), target.Check(this.item1))
	assert.True(this.T(), target.Check(this.item2))
	assert.True(this.T(), target.Check(this.item3))
	assert.False(this.T(), target.Check(this.item1))
	assert.False(this.T(), target.Check(this.item2))
	assert.False(this.T(), target.Check(this.item3))
}

func (this *SuiteDuplicate) TestNewDuplicate() {
	assert.NotNil(this.T(), NewDuplicate())
}
