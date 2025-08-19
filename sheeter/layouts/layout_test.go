package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestLayout(t *testing.T) {
	suite.Run(t, new(SuiteLayout))
}

type SuiteLayout struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteLayout) SetupSuite() {
	this.Env = testdata.EnvSetup("test-layouts-layout")
}

func (this *SuiteLayout) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteLayout) TestLayout() {
	target, failed := NewLayout(
		[]string{"", "", "ignore", "12", "12", "13", "13", ""},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5", ""},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5", ""},
		[]string{"", "int", "", "int", "[]int", "string", "[]string", ""},
	)
	assert.Empty(this.T(), failed)
	assert.NotNil(this.T(), target)

	_, failed = NewLayout(
		[]string{"", ""},
		[]string{"", "name"},
		[]string{"", "note"},
		[]string{"", testdata.Unknown},
	)
	assert.NotEmpty(this.T(), failed)

	_, failed = NewLayout(
		[]string{"", ""},
		[]string{"", ""},
		[]string{"", "note"},
		[]string{"", "int"},
	)
	assert.NotEmpty(this.T(), failed)

	_, failed = NewLayout(
		[]string{"", "", "1"},
		[]string{"", "name", "name"},
		[]string{"", "note", "note"},
		[]string{"", "int", "int"},
	)
	assert.NotEmpty(this.T(), failed)

	_, failed = NewLayout(
		[]string{""},
		[]string{""},
		[]string{""},
		[]string{""},
	)
	assert.NotEmpty(this.T(), failed)
}

func (this *SuiteLayout) TestPrimary() {
	target, _ := NewLayout(
		[]string{"", "", "ignore", "12", "12", "13", "13", ""},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5", ""},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5", ""},
		[]string{"", "int", "", "int", "[]int", "string", "[]string", ""},
	)
	assert.NotNil(this.T(), target.Primary())
}

func (this *SuiteLayout) TestSelect() {
	target, _ := NewLayout(
		[]string{"", "", "ignore", "12", "12", "13", "13", ""},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5", ""},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5", ""},
		[]string{"", "int", "", "int", "[]int", "string", "[]string", ""},
	)
	assert.Len(this.T(), target.Select("1"), 5)
	assert.Len(this.T(), target.Select("2"), 3)
	assert.Len(this.T(), target.Select("3"), 3)
	assert.Len(this.T(), target.Select("4"), 1)
}

func (this *SuiteLayout) TestPack() {
	data := []string{"", "1", "", "2", "1,2,3", "a", "a,b,c", ""}
	dataInvalid := []string{"", "1", "", testdata.Unknown, "1,2,3", "a", "a,b,c", ""}
	actual1 := map[string]interface{}{
		"name1": int32(1),
		"name2": int32(2),
		"name3": []int32{1, 2, 3},
		"name4": "a",
		"name5": []string{"a", "b", "c"},
	}
	actual2 := map[string]interface{}{
		"name1": int32(1),
		"name2": int32(2),
		"name3": []int32{1, 2, 3},
	}
	actual3 := map[string]interface{}{
		"name1": int32(1),
		"name4": "a",
		"name5": []string{"a", "b", "c"},
	}
	target, _ := NewLayout(
		[]string{"", "", "ignore", "12", "12", "13", "13", ""},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5", ""},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5", ""},
		[]string{"", "int", "", "int", "[]int", "string", "[]string", ""},
	)

	primary, pack, err := target.Pack("1", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(1), primary)
	assert.Equal(this.T(), actual1, pack)

	primary, pack, err = target.Pack("2", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(1), primary)
	assert.Equal(this.T(), actual2, pack)

	primary, pack, err = target.Pack("3", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(1), primary)
	assert.Equal(this.T(), actual3, pack)

	_, _, err = target.Pack("1", dataInvalid)
	assert.NotNil(this.T(), err)
}
