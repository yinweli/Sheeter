package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/internal/build/layers"
)

func TestParser(t *testing.T) {
	suite.Run(t, new(SuiteParser))
}

type SuiteParser struct {
	suite.Suite
	fieldPkey    fields.Field
	fieldInt     fields.Field
	layerStruct  []layers.Layer
	layerArray   []layers.Layer
	layerDivider []layers.Layer
	layerFailed  []layers.Layer
	layerEmpty   []layers.Layer
	packs        map[string]interface{}
	data         []string
	dataJson     []string
}

func (this *SuiteParser) SetupSuite() {
	this.fieldPkey = &fields.Pkey{}
	this.fieldInt = &fields.Int{}
	this.layerStruct, _, _ = layers.Parser("{S")
	this.layerArray, _, _ = layers.Parser("{[]N")
	this.layerDivider, _, _ = layers.Parser("/")
	this.layerFailed, _, _ = layers.Parser("{[]S")
	this.layerEmpty = []layers.Layer{}
	this.packs = map[string]interface{}{
		"S": map[string]interface{}{
			"n1": int64(1),
			"N": &[]map[string]interface{}{
				{"a1": int64(2), "a2": int64(3)},
				{"a1": int64(4), "a2": int64(5)},
			},
		},
	}
	this.data = []string{"0", "1", "2", "3", "4", "5"}
	this.dataJson = []string{"0", "a", "2", "3", "4", "5"}
}

func (this *SuiteParser) target() *Parser {
	return NewParser()
}

func (this *SuiteParser) TestAdd() {
	target := this.target()

	assert.Nil(this.T(), target.Add("n1", "", this.fieldPkey, this.layerStruct, 0))
	assert.Nil(this.T(), target.Add("n2", "", this.fieldInt, this.layerArray, 0))
	assert.Nil(this.T(), target.Add("n3", "", this.fieldInt, this.layerEmpty, 0))
	assert.Nil(this.T(), target.Add("n4", "", this.fieldInt, this.layerDivider, 0))
	assert.Nil(this.T(), target.Add("n5", "", this.fieldInt, this.layerEmpty, 2))
	assert.NotNil(this.T(), target.Add("", "", nil, nil, 0))
	assert.NotNil(this.T(), target.Add("n6", "", nil, nil, 0))
	assert.NotNil(this.T(), target.Add("n7", "", this.fieldPkey, nil, 0))
	assert.NotNil(this.T(), target.Add("n8", "", this.fieldInt, this.layerFailed, 0))
	assert.NotNil(this.T(), target.Add("n8", "", this.fieldInt, this.layerEmpty, -1))
}

func (this *SuiteParser) TestPack() {
	target := this.target()

	assert.Nil(this.T(), target.Add("n0", "", &fields.Empty{}, this.layerEmpty, 0))
	assert.Nil(this.T(), target.Add("n1", "", this.fieldPkey, this.layerStruct, 0))
	assert.Nil(this.T(), target.Add("a1", "", this.fieldInt, this.layerArray, 0))
	assert.Nil(this.T(), target.Add("a2", "", this.fieldInt, this.layerEmpty, 0))
	assert.Nil(this.T(), target.Add("a1", "", this.fieldInt, this.layerDivider, 0))
	assert.Nil(this.T(), target.Add("a2", "", this.fieldInt, this.layerEmpty, 2))

	packs, pkey, err := target.Pack(this.data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "1", pkey)
	assert.Equal(this.T(), this.packs, packs)

	_, _, err = target.Pack(this.dataJson)
	assert.NotNil(this.T(), err)
}

func (this *SuiteParser) TestLayouts() {
	target := this.target()

	assert.Nil(this.T(), target.Add("n0", "", &fields.Empty{}, this.layerEmpty, 0))
	assert.Nil(this.T(), target.Add("n1", "", this.fieldPkey, this.layerStruct, 0))
	assert.Nil(this.T(), target.Add("a1", "", this.fieldInt, this.layerArray, 0))
	assert.Nil(this.T(), target.Add("a2", "", this.fieldInt, this.layerEmpty, 0))
	assert.Nil(this.T(), target.Add("a1", "", this.fieldInt, this.layerDivider, 0))
	assert.Nil(this.T(), target.Add("a2", "", this.fieldInt, this.layerEmpty, 2))
	assert.NotNil(this.T(), target.Layouts())
}

func (this *SuiteParser) TestNewParser() {
	assert.NotNil(this.T(), NewParser())
}
