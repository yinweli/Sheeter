package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/sheeter/fields"
	"github.com/yinweli/Sheeter/sheeter/layers"
	"github.com/yinweli/Sheeter/sheeter/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJsonPack(t *testing.T) {
	suite.Run(t, new(SuiteJsonPack))
}

type SuiteJsonPack struct {
	suite.Suite
	testdata.TestEnv
	lineOfName  int
	lineOfField int
	lineOfLayer int
	lineOfData  int
	excel       excels.Excel
}

func (this *SuiteJsonPack) SetupSuite() {
	this.Change("test-jsonPack")
	this.lineOfName = 1
	this.lineOfField = 3
	this.lineOfLayer = 4
	this.lineOfData = 5
	assert.Nil(this.T(), this.excel.Open(testdata.ExcelJsonPack))
}

func (this *SuiteJsonPack) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteJsonPack) TestJsonPack() {
	completeBytes, _ := utils.JsonMarshal(map[string]interface{}{
		"Datas": map[sheeter.PkeyType]interface{}{
			1: map[string]interface{}{
				"Name0": 1,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
					},
					"Name1": true,
				},
			},
			2: map[string]interface{}{
				"Name0": 2,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
					},
					"Name1": false,
				},
			},
			3: map[string]interface{}{
				"Name0": 3,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
					},
					"Name1": true,
				},
			},
		},
	})
	excludeBytes, _ := utils.JsonMarshal(map[string]interface{}{
		"Datas": map[sheeter.PkeyType]interface{}{
			1: map[string]interface{}{
				"Name0": 1,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
					},
					"Name1": true,
				},
				"Name4": 1,
			},
			2: map[string]interface{}{
				"Name0": 2,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
					},
					"Name1": false,
				},
				"Name4": 2,
			},
			3: map[string]interface{}{
				"Name0": 3,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
					},
					"Name1": true,
				},
				"Name4": 3,
			},
		},
	})

	line, err := this.excel.GetLine(testdata.SheetData, this.lineOfName, this.lineOfField, this.lineOfLayer)
	assert.Nil(this.T(), err)
	nameLine := line[this.lineOfName]
	fieldLine := line[this.lineOfField]
	layerLine := line[this.lineOfLayer]
	layoutData := NewLayoutData()

	for col, itor := range nameLine {
		name := itor
		field, tag, err := fields.Parser(utils.GetItem(fieldLine, col))
		assert.Nil(this.T(), err)
		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))
		assert.Nil(this.T(), err)
		assert.Nil(this.T(), layoutData.Add(name, field, tag, layer, back))
	} // for

	sheet, err := this.excel.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Nextn(this.lineOfData))
	json, err := JsonPack(sheet, layoutData, []string{"tag"})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(completeBytes), string(json))
	sheet.Close()

	sheet, err = this.excel.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Nextn(this.lineOfData))
	json, err = JsonPack(sheet, layoutData, []string{})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(excludeBytes), string(json))
	sheet.Close()
}

func (this *SuiteJsonPack) TestJsonFirstUpper() {
	input := map[string]interface{}{
		"name1": 1,
		"name2": []int{1, 2, 3},
		"name3": map[string]interface{}{
			"name1": 1,
			"name2": []int{1, 2, 3},
			"name3": map[string]interface{}{
				"name1": "a",
				"name2": "b",
			},
		},
		"name4": &[]map[string]interface{}{
			{
				"name1": 1,
				"name2": []int{1, 2, 3},
				"name3": map[string]interface{}{
					"name1": "a",
					"name2": "b",
				},
			},
			{
				"name1": 1,
				"name2": []int{1, 2, 3},
				"name3": map[string]interface{}{
					"name1": "a",
					"name2": "b",
				},
			},
		},
		"name5": &[]map[string]interface{}{
			{
				"name1": &[]map[string]interface{}{
					{
						"name1": 1,
						"name2": []int{1, 2, 3},
					},
					{
						"name1": 1,
						"name2": []int{1, 2, 3},
					},
				},
			},
		},
	}
	expected := map[string]interface{}{
		"Name1": 1,
		"Name2": []int{1, 2, 3},
		"Name3": map[string]interface{}{
			"Name1": 1,
			"Name2": []int{1, 2, 3},
			"Name3": map[string]interface{}{
				"Name1": "a",
				"Name2": "b",
			},
		},
		"Name4": &[]map[string]interface{}{
			{
				"Name1": 1,
				"Name2": []int{1, 2, 3},
				"Name3": map[string]interface{}{
					"Name1": "a",
					"Name2": "b",
				},
			},
			{
				"Name1": 1,
				"Name2": []int{1, 2, 3},
				"Name3": map[string]interface{}{
					"Name1": "a",
					"Name2": "b",
				},
			},
		},
		"Name5": &[]map[string]interface{}{
			{
				"Name1": &[]map[string]interface{}{
					{
						"Name1": 1,
						"Name2": []int{1, 2, 3},
					},
					{
						"Name1": 1,
						"Name2": []int{1, 2, 3},
					},
				},
			},
		},
	}

	assert.Equal(this.T(), expected, jsonFirstUpper(input))
}
