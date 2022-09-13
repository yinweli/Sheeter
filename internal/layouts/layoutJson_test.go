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
	workDir        string
	fieldEmpty     fields.Field
	fieldPkey      fields.Field
	fieldInt       fields.Field
	layerStruct    []layers.Layer
	layerArray     []layers.Layer
	layerDivider   []layers.Layer
	layerDuplicate []layers.Layer
	layerNone      []layers.Layer
	packPreset     map[string]interface{}
	packValue      map[string]interface{}
	dataValid      []string
	dataInvalid    []string
}

func (this *SuiteLayoutJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.fieldEmpty = &fields.Empty{}
	this.fieldPkey = &fields.Pkey{}
	this.fieldInt = &fields.Int{}
	this.layerStruct, _, _ = layers.Parser("{S")
	this.layerArray, _, _ = layers.Parser("{[]N")
	this.layerDivider, _, _ = layers.Parser("/")
	this.layerDuplicate, _, _ = layers.Parser("{[]S")
	this.layerNone = []layers.Layer{}
	this.packPreset = map[string]interface{}{
		"S": map[string]interface{}{
			"n1": int64(0),
			"N": &[]map[string]interface{}{
				{"a1": int64(0), "a2": int64(0)},
				{"a1": int64(0), "a2": int64(0)},
			},
		},
	}
	this.packValue = map[string]interface{}{
		"S": map[string]interface{}{
			"n1": int64(1),
			"N": &[]map[string]interface{}{
				{"a1": int64(2), "a2": int64(3)},
				{"a1": int64(4), "a2": int64(5)},
			},
		},
	}
	this.dataValid = []string{"0", "1", "2", "3", "4", "5"}
	this.dataInvalid = []string{"0", "a", "2", "3", "4", "5"}
}

func (this *SuiteLayoutJson) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteLayoutJson) target() *LayoutJson {
	return NewLayoutJson()
}

func (this *SuiteLayoutJson) TestNewLayoutJson() {
	assert.NotNil(this.T(), NewLayoutJson())
}

func (this *SuiteLayoutJson) TestAdd() {
	target := this.target()

	assert.Nil(this.T(), target.Add("n1", this.fieldPkey, this.layerStruct, 0))
	assert.Nil(this.T(), target.Add("n2", this.fieldInt, this.layerArray, 0))
	assert.Nil(this.T(), target.Add("n3", this.fieldInt, this.layerNone, 0))
	assert.Nil(this.T(), target.Add("n4", this.fieldInt, this.layerDivider, 0))
	assert.Nil(this.T(), target.Add("n5", this.fieldInt, this.layerNone, 2))
	assert.NotNil(this.T(), target.Add("", nil, nil, 0))
	assert.NotNil(this.T(), target.Add("n6", nil, nil, 0))
	assert.NotNil(this.T(), target.Add("n7", this.fieldInt, this.layerDuplicate, 0))
	assert.NotNil(this.T(), target.Add("n7", this.fieldInt, this.layerNone, -1))
}

func (this *SuiteLayoutJson) TestPack() {
	target := this.target()

	assert.Nil(this.T(), target.Add("n0", this.fieldEmpty, this.layerNone, 0))
	assert.Nil(this.T(), target.Add("n1", this.fieldPkey, this.layerStruct, 0))
	assert.Nil(this.T(), target.Add("a1", this.fieldInt, this.layerArray, 0))
	assert.Nil(this.T(), target.Add("a2", this.fieldInt, this.layerNone, 0))
	assert.Nil(this.T(), target.Add("a1", this.fieldInt, this.layerDivider, 0))
	assert.Nil(this.T(), target.Add("a2", this.fieldInt, this.layerNone, 2))

	packs, pkey, err := target.Pack(this.dataValid, true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "0", pkey)
	assert.Equal(this.T(), this.packPreset, packs)

	packs, pkey, err = target.Pack(this.dataValid, false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "1", pkey)
	assert.Equal(this.T(), this.packValue, packs)

	_, _, err = target.Pack(this.dataInvalid, false)
	assert.NotNil(this.T(), err)
}

func (this *SuiteLayoutJson) TestPkeyCount() {
	target := this.target()

	assert.Equal(this.T(), 0, target.PkeyCount())
	assert.Nil(this.T(), target.Add("n0", this.fieldPkey, this.layerNone, 0))
	assert.Equal(this.T(), 1, target.PkeyCount())
}
