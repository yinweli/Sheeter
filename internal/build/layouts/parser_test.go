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
	item1      Layout
	item2      Layout
	item3      Layout
	item4      Layout
	itemName   Layout
	itemField  Layout
	itemPkey   Layout
	itemLayer1 Layout
	itemLayer2 Layout
	itemBack   Layout
}

func (this *SuiteParser) SetupSuite() {
	this.item1 = Layout{
		Name:   "name1",
		Note:   "note1",
		Field:  &fields.Pkey{},
		Layers: []layers.Layer{{Name: "array", Type: layers.LayerArray}, {Name: "struct", Type: layers.LayerStruct}},
		Back:   0,
	}
	this.item2 = Layout{
		Name:   "name2",
		Note:   "note2",
		Field:  &fields.Int{},
		Layers: []layers.Layer{},
		Back:   1,
	}
	this.item3 = Layout{
		Name:   "name3",
		Note:   "note3",
		Field:  &fields.Int{},
		Layers: []layers.Layer{{Name: "", Type: layers.LayerDivider}, {Name: "struct", Type: layers.LayerStruct}},
		Back:   0,
	}
	this.item4 = Layout{
		Name:   "name4",
		Note:   "note4",
		Field:  &fields.Int{},
		Layers: []layers.Layer{},
		Back:   2,
	}
	this.itemName = Layout{
		Name:   "",
		Note:   "",
		Field:  nil,
		Layers: nil,
		Back:   0,
	}
	this.itemField = Layout{
		Name:   "nameField",
		Note:   "",
		Field:  nil,
		Layers: nil,
		Back:   0,
	}
	this.itemPkey = Layout{
		Name:   "namePkey",
		Note:   "",
		Field:  &fields.Pkey{},
		Layers: []layers.Layer{},
		Back:   0,
	}
	this.itemLayer1 = Layout{
		Name:   "nameLayer1",
		Note:   "",
		Field:  &fields.Int{},
		Layers: []layers.Layer{{Name: "data", Type: layers.LayerDivider}},
		Back:   0,
	}
	this.itemLayer2 = Layout{
		Name:   "nameLayer2",
		Note:   "",
		Field:  &fields.Int{},
		Layers: []layers.Layer{{Name: "data", Type: layers.LayerStruct}},
		Back:   0,
	}
	this.itemBack = Layout{
		Name:   "nameBack",
		Note:   "",
		Field:  &fields.Int{},
		Layers: []layers.Layer{},
		Back:   -1,
	}
}

func (this *SuiteParser) target() *Parser {
	return NewParser()
}

func (this *SuiteParser) add(parser *Parser, layout Layout) error {
	return parser.Add(layout.Name, layout.Note, layout.Field, layout.Layers, layout.Back)
}

func (this *SuiteParser) TestAdd() {
	target := this.target()

	assert.Nil(this.T(), this.add(target, this.item1))
	assert.Nil(this.T(), this.add(target, this.item2))
	assert.Nil(this.T(), this.add(target, this.item3))
	assert.Nil(this.T(), this.add(target, this.item4))
	assert.NotNil(this.T(), this.add(target, this.itemName))
	assert.NotNil(this.T(), this.add(target, this.item1))
	assert.NotNil(this.T(), this.add(target, this.itemField))
	assert.NotNil(this.T(), this.add(target, this.itemPkey))
	assert.Nil(this.T(), this.add(target, this.itemLayer1))
	assert.NotNil(this.T(), this.add(target, this.itemLayer2))
	assert.NotNil(this.T(), this.add(target, this.itemBack))
}

func (this *SuiteParser) TestPack() {
	// TODO: 做到這裡!
}

func (this *SuiteParser) TestLayouts() {

}

func (this *SuiteParser) TestNewParser() {
	assert.NotNil(this.T(), NewParser())
}
