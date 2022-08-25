package util

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMpb(t *testing.T) {
	suite.Run(t, new(SuiteMpb))
}

type SuiteMpb struct {
	suite.Suite
}

func (this *SuiteMpb) TestNewMpb() {
	assert.NotNil(this.T(), NewMpb(&sync.WaitGroup{}))
}
