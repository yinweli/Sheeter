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
	assert.Nil(this.T(), target.Add("name1", &fields.Pkey{}, this.layer("{type1"), 0))
	assert.Nil(this.T(), target.Add("name2", &fields.Int{}, this.layer("{[]type2"), 0))
	assert.Nil(this.T(), target.Add("name3", &fields.Int{}, this.layer(""), 0))
	assert.Nil(this.T(), target.Add("name4", &fields.Int{}, this.layer("/"), 0))
	assert.Nil(this.T(), target.Add("name5", &fields.Int{}, this.layer(""), 2))
	assert.NotNil(this.T(), target.Add("", nil, nil, 0))
	assert.NotNil(this.T(), target.Add("name6", nil, nil, 0))
	assert.NotNil(this.T(), target.Add("name7", &fields.Int{}, this.layer("{[]type1"), 0))
	assert.NotNil(this.T(), target.Add("name7", &fields.Int{}, this.layer(""), -1))
}

func (this *SuiteLayoutJson) TestPack() {
	preset := map[string]interface{}{
		"data": map[string]interface{}{
			"name1": int64(0),
			"array": &[]map[string]interface{}{
				{"array1": int64(0), "array2": int64(0)},
				{"array1": int64(0), "array2": int64(0)},
			},
		},
	}
	actual := map[string]interface{}{
		"data": map[string]interface{}{
			"name1": int64(1),
			"array": &[]map[string]interface{}{
				{"array1": int64(2), "array2": int64(3)},
				{"array1": int64(4), "array2": int64(5)},
			},
		},
	}
	dataValid := []string{"0", "1", "2", "3", "4", "5"}
	dataInvalid := []string{"0", "a", "2", "3", "4", "5"}

	target := this.target()
	assert.Nil(this.T(), target.Add("name0", &fields.Empty{}, this.layer(""), 0))
	assert.Nil(this.T(), target.Add("name1", &fields.Pkey{}, this.layer("{data"), 0))
	assert.Nil(this.T(), target.Add("array1", &fields.Int{}, this.layer("{[]array"), 0))
	assert.Nil(this.T(), target.Add("array2", &fields.Int{}, this.layer(""), 0))
	assert.Nil(this.T(), target.Add("array1", &fields.Int{}, this.layer("/"), 0))
	assert.Nil(this.T(), target.Add("array2", &fields.Int{}, this.layer(""), 2))

	packs, pkey, err := target.Pack(dataValid, true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(0), pkey)
	assert.Equal(this.T(), preset, packs)

	packs, pkey, err = target.Pack(dataValid, false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(1), pkey)
	assert.Equal(this.T(), actual, packs)

	_, _, err = target.Pack(dataInvalid, false)
	assert.NotNil(this.T(), err)
}

func (this *SuiteLayoutJson) TestPkeyCount() {
	target := this.target()
	assert.Equal(this.T(), 0, target.PkeyCount())
	assert.Nil(this.T(), target.Add("name", &fields.Pkey{}, this.layer(""), 0))
	assert.Equal(this.T(), 1, target.PkeyCount())
}
