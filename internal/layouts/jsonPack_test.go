package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xuri/excelize/v2"

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
	excelName   string
	sheetName   string
	lineOfField int
	lineOfLayer int
	lineOfData  int
	excel       *excelize.File
}

func (this *SuiteJsonPack) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excelName = testdata.ExcelNameJsonPack
	this.sheetName = testdata.SheetName
	this.lineOfField = 1
	this.lineOfLayer = 2
	this.lineOfData = 4

	excel, err := excelize.OpenFile(this.excelName)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), excel)
	this.excel = excel
}

func (this *SuiteJsonPack) TearDownSuite() {
	_ = this.excel.Close()
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJsonPack) getRows(line int) *excelize.Rows {
	rows, err := this.excel.Rows(this.sheetName)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)

	for l := 0; l < line; l++ {
		assert.True(this.T(), rows.Next())
	} // for

	return rows
}

func (this *SuiteJsonPack) getCols(line int) []string {
	rows, err := this.excel.Rows(this.sheetName)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	defer func() { _ = rows.Close() }()

	for l := 0; l < line; l++ {
		assert.True(this.T(), rows.Next())
	} // for

	cols, err := rows.Columns()
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), cols)

	return cols
}

func (this *SuiteJsonPack) TestJsonPack() {
	data1, err := utils.JsonMarshal(testdata.GetExcelContentJsonPack(true))
	assert.Nil(this.T(), err)
	data2, err := utils.JsonMarshal(testdata.GetExcelContentJsonPack(false))
	assert.Nil(this.T(), err)

	fieldCol := this.getCols(this.lineOfField)
	layerCol := this.getCols(this.lineOfLayer)
	layoutJson := NewLayoutJson()

	for col, itor := range fieldCol {
		name, field, tag, err := fields.Parser(itor)
		assert.Nil(this.T(), err)
		layer, back, err := layers.Parser(utils.GetItem(layerCol, col))
		assert.Nil(this.T(), err)
		assert.Nil(this.T(), layoutJson.Add(name, field, tag, layer, back))
	} // for

	rows := this.getRows(this.lineOfData)
	json, err := JsonPack(rows, layoutJson, []string{"tag"})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(data1), string(json))
	_ = rows.Close()

	rows = this.getRows(this.lineOfData)
	json, err = JsonPack(rows, layoutJson, []string{})
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(data2), string(json))
	_ = rows.Close()
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
