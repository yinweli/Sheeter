package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/testdata"
)

func TestLayoutJson(t *testing.T) {
	suite.Run(t, new(SuiteLayoutJson))
}

type SuiteLayoutJson struct {
	suite.Suite
	workDir string
}

func (this *SuiteLayoutJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteLayoutJson) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteLayoutJson) target() *LayoutJson {
	return NewLayoutJson()
}

func (this *SuiteLayoutJson) layer(input string) []layers.Layer {
	layer, _, _ := layers.Parser(input)
	return layer
}

func (this *SuiteLayoutJson) TestNewLayoutJson() {
	assert.NotNil(this.T(), NewLayoutJson())
}

func (this *SuiteLayoutJson) TestAdd() {
	target := this.target()
	assert.Nil(this.T(), target.Add("name1", &fields.Pkey{}, "", this.layer("{type1"), 0))
	assert.Nil(this.T(), target.Add("name2", &fields.Int{}, "", this.layer("{[]type2"), 0))
	assert.Nil(this.T(), target.Add("name3", &fields.Int{}, "", this.layer(""), 0))
	assert.Nil(this.T(), target.Add("name4", &fields.Int{}, "", this.layer("/"), 0))
	assert.Nil(this.T(), target.Add("name5", &fields.Int{}, "", this.layer(""), 2))
	assert.NotNil(this.T(), target.Add("", nil, "", nil, 0))
	assert.NotNil(this.T(), target.Add("name6", nil, "", nil, 0))
	assert.NotNil(this.T(), target.Add("name7", &fields.Int{}, "", this.layer("{[]type1"), 0))
	assert.NotNil(this.T(), target.Add("name7", &fields.Int{}, "", this.layer(""), -1))
}

func (this *SuiteLayoutJson) TestPack() {
	actual1 := map[string]interface{}{
		"data": map[string]interface{}{
			"name1": int64(1),
			"name2": int64(2),
			"array": &[]map[string]interface{}{
				{"array1": int64(3), "array2": int64(4)},
				{"array1": int64(5), "array2": int64(6)},
			},
		},
	}
	actual2 := map[string]interface{}{
		"data": map[string]interface{}{
			"name1": int64(1),
			"array": &[]map[string]interface{}{
				{"array1": int64(3), "array2": int64(4)},
				{"array1": int64(5), "array2": int64(6)},
			},
		},
	}
	dataValid := []string{"0", "1", "2", "3", "4", "5", "6"}
	dataInvalid := []string{"0", "1", "a", "3", "4", "5", "6"}

	target := this.target()
	assert.Nil(this.T(), target.Add("name0", &fields.Empty{}, "", this.layer(""), 0))
	assert.Nil(this.T(), target.Add("name1", &fields.Pkey{}, "", this.layer("{data"), 0))
	assert.Nil(this.T(), target.Add("name2", &fields.Int{}, "tag", this.layer(""), 0))
	assert.Nil(this.T(), target.Add("array1", &fields.Int{}, "", this.layer("{[]array"), 0))
	assert.Nil(this.T(), target.Add("array2", &fields.Int{}, "", this.layer(""), 0))
	assert.Nil(this.T(), target.Add("array1", &fields.Int{}, "", this.layer("/"), 0))
	assert.Nil(this.T(), target.Add("array2", &fields.Int{}, "", this.layer(""), 2))

	packs, pkey, err := target.Pack(dataValid, []string{})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(1), pkey)
	assert.Equal(this.T(), actual1, packs)

	packs, pkey, err = target.Pack(dataValid, []string{""})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(1), pkey)
	assert.Equal(this.T(), actual1, packs)

	packs, pkey, err = target.Pack(dataValid, []string{"tag"})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(1), pkey)
	assert.Equal(this.T(), actual2, packs)

	_, _, err = target.Pack(dataInvalid, []string{})
	assert.NotNil(this.T(), err)
}

func (this *SuiteLayoutJson) TestPkeyCount() {
	target := this.target()
	assert.Equal(this.T(), 0, target.PkeyCount())
	assert.Nil(this.T(), target.Add("name", &fields.Pkey{}, "", this.layer(""), 0))
	assert.Equal(this.T(), 1, target.PkeyCount())
}

func (this *SuiteLayoutJson) TestIsExclude() {
	assert.True(this.T(), isExclude("tag1", []string{"tag1"}))
	assert.False(this.T(), isExclude("tag2", []string{}))
}
