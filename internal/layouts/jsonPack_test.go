package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJsonPack(t *testing.T) {
	suite.Run(t, new(SuiteJsonPack))
}

type SuiteJsonPack struct {
	suite.Suite
	workDir     string
	lineOfName  int
	lineOfField int
	lineOfLayer int
	lineOfData  int
	excel       excels.Excel
}

func (this *SuiteJsonPack) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.lineOfName = 1
	this.lineOfField = 3
	this.lineOfLayer = 4
	this.lineOfData = 5
	assert.Nil(this.T(), this.excel.Open(testdata.ExcelJsonPack))
}

func (this *SuiteJsonPack) TearDownSuite() {
	this.excel.Close()
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJsonPack) TestJsonPack() {
	data1, err := utils.JsonMarshal(testdata.GetExcelContentJsonPack(true))
	assert.Nil(this.T(), err)
	data2, err := utils.JsonMarshal(testdata.GetExcelContentJsonPack(false))
	assert.Nil(this.T(), err)

	line, err := this.excel.GetLine(testdata.SheetData, this.lineOfName, this.lineOfField, this.lineOfLayer)
	assert.Nil(this.T(), err)
	nameLine := line[this.lineOfName]
	fieldLine := line[this.lineOfField]
	layerLine := line[this.lineOfLayer]
	layoutJson := NewLayoutJson()

	for col, itor := range nameLine {
		name := itor
		field, tag, err := fields.Parser(utils.GetItem(fieldLine, col))
		assert.Nil(this.T(), err)
		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))
		assert.Nil(this.T(), err)
		assert.Nil(this.T(), layoutJson.Add(name, field, tag, layer, back))
	} // for

	sheet, err := this.excel.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Nextn(this.lineOfData))
	json, err := JsonPack(sheet, layoutJson, []string{"tag"})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(data1), string(json))
	sheet.Close()

	sheet, err = this.excel.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Nextn(this.lineOfData))
	json, err = JsonPack(sheet, layoutJson, []string{})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(data2), string(json))
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
