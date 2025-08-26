package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestString(t *testing.T) {
	suite.Run(t, new(SuiteString))
}

type SuiteString struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteString) SetupSuite() {
	this.Env = testdata.EnvSetup("test-utils-string")
}

func (this *SuiteString) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteString) TestFirstUpper() {
	assert.Equal(this.T(), "", FirstUpper(""))
	assert.Equal(this.T(), "T", FirstUpper("t"))
	assert.Equal(this.T(), "TestString", FirstUpper("testString"))
}

func (this *SuiteString) TestFirstLower() {
	assert.Equal(this.T(), "", FirstLower(""))
	assert.Equal(this.T(), "t", FirstLower("T"))
	assert.Equal(this.T(), "testString", FirstLower("TestString"))
}

func (this *SuiteString) TestSnakeToCamel() {
	assert.Equal(this.T(), "", SnakeToCamel(""))
	assert.Equal(this.T(), "", SnakeToCamel("_"))
	assert.Equal(this.T(), "Abc", SnakeToCamel("abc"))
	assert.Equal(this.T(), "Abc", SnakeToCamel("abc_"))
	assert.Equal(this.T(), "AbcX", SnakeToCamel("abc_x"))
	assert.Equal(this.T(), "AbcXY", SnakeToCamel("abc_x_y"))
	assert.Equal(this.T(), "AbcXaYa", SnakeToCamel("abc_xa_ya"))
	assert.Equal(this.T(), "Abc", SnakeToCamel("abc "))
	assert.Equal(this.T(), "AbcX", SnakeToCamel("abc x"))
	assert.Equal(this.T(), "AbcXY", SnakeToCamel("abc x y"))
	assert.Equal(this.T(), "AbcXaYa", SnakeToCamel("abc xa ya"))
}

func (this *SuiteString) TestAllSame() {
	assert.Equal(this.T(), true, AllSame(""))
	assert.Equal(this.T(), true, AllSame("aaaaa"))
	assert.Equal(this.T(), false, AllSame("aa1aa"))
	assert.Equal(this.T(), true, AllSame("好好好"))
	assert.Equal(this.T(), false, AllSame("好不好"))
}

func (this *SuiteString) TestCombine() {
	assert.ElementsMatch(this.T(), []string{"a", "b", "c", "1", "2", "3"}, Combine([]string{"a", "b", "c"}, []any{"1", "2", "3"}))
}

func (this *SuiteString) TestAt() {
	item := []string{"a", "b", "c"}
	assert.Equal(this.T(), "a", At(item, 0))
	assert.Equal(this.T(), "b", At(item, 1))
	assert.Equal(this.T(), "c", At(item, 2))
	assert.Equal(this.T(), "", At(item, 3))
}

func (this *SuiteString) TestUnique() {
	item := []string{"a", "b", "c", "a", "b", "c"}
	assert.ElementsMatch(this.T(), []string{"a", "b", "c"}, Unique(item))
}
