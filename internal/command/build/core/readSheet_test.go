package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadSheet(t *testing.T) {
	cargo := mockCargo()

	err := ReadSheet(cargo, 0)
	assert.Nil(t, err)

	cargo.Element.Excel = testdata.Error1Excel
	cargo.Element.Sheet = testdata.Error1Sheet
	err = ReadSheet(cargo, 0)
	assert.NotNil(t, err)

	cargo.Element.Excel = testdata.Error2Excel
	cargo.Element.Sheet = testdata.Error2Sheet
	err = ReadSheet(cargo, 0)
	assert.NotNil(t, err)

	cargo.Element.Excel = testdata.Error3Excel
	cargo.Element.Sheet = testdata.Error3Sheet
	err = ReadSheet(cargo, 0)
	assert.NotNil(t, err)
}

func TestBuildSheet(t *testing.T) {
	cargo := mockCargo()

	sheet, err := buildSheet(cargo, 0)
	assert.Nil(t, err)
	assert.Equal(t, 12, len(sheet))
	assert.Equal(t, 16, len(sheet[0]))
	assert.Equal(t, "checkpoint", sheet[0][15])
	assert.Equal(t, "checkpoint", sheet[11][15])

	cargo.Element.Excel = testdata.FakeExcel
	cargo.Element.Sheet = testdata.FakeSheet
	sheet, err = buildSheet(cargo, 0)
	assert.NotNil(t, err)

	cargo.Element.Excel = testdata.FakeExcel
	cargo.Element.Sheet = testdata.UnknownSheet
	sheet, err = buildSheet(cargo, 0)
	assert.NotNil(t, err)

	cargo.Element.Excel = testdata.UnknownExcel
	sheet, err = buildSheet(cargo, 0)
	assert.NotNil(t, err)
}

func TestBuildColumns(t *testing.T) {
	cargo := mockCargo()
	fields := []string{"field0#pkey", "field1#bool", "field2#int", "", "field3#text"}

	pkey, err := buildColumns(cargo, [][]string{{}, fields})
	assert.Nil(t, err)
	assert.NotNil(t, pkey)
	assert.Equal(t, "field0", pkey.Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), pkey.Field.TypeExcel())
	assert.Equal(t, 3, len(cargo.Columns))
	assert.Equal(t, "field0", cargo.Columns[0].Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), cargo.Columns[0].Field.TypeExcel())
	assert.Equal(t, "field1", cargo.Columns[1].Name)
	assert.Equal(t, (&FieldBool{}).TypeExcel(), cargo.Columns[1].Field.TypeExcel())
	assert.Equal(t, "field2", cargo.Columns[2].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), cargo.Columns[2].Field.TypeExcel())

	fields = []string{"field0#????", "field1#bool", "field2#int"}
	pkey, err = buildColumns(cargo, [][]string{{}, fields})
	assert.NotNil(t, err)

	fields = []string{"field0#pkey", "field1#pkey", "field2#int"}
	pkey, err = buildColumns(cargo, [][]string{{}, fields})
	assert.NotNil(t, err)

	fields = []string{"field0#int", "field1#int", "field2#int"}
	pkey, err = buildColumns(cargo, [][]string{{}, fields})
	assert.NotNil(t, err)

	fields = []string{}
	pkey, err = buildColumns(cargo, [][]string{{}, fields})
	assert.NotNil(t, err)
}

func TestBuildNotes(t *testing.T) {
	cargo := mockCargo()
	notes := []string{"note0", "note1", "note2"}

	err := buildNotes(cargo, [][]string{notes})
	assert.Nil(t, err)
	assert.Equal(t, "note0", cargo.Columns[0].Note)
	assert.Equal(t, "note1", cargo.Columns[1].Note)
	assert.Equal(t, "note2", cargo.Columns[2].Note)
}

func TestBuildDatas(t *testing.T) {
	cargo := mockCargo()
	data0 := []string{"data0", "data1", "data2"}
	data1 := []string{"data4", "data5", "data6"}
	data2 := []string{"data7", "data8", "data9"}

	err := buildDatas(cargo, [][]string{{}, {}, data0, data1, data2})
	assert.Nil(t, err)
	assert.Equal(t, []string{"data0", "data4", "data7"}, cargo.Columns[0].Datas)
	assert.Equal(t, []string{"data1", "data5", "data8"}, cargo.Columns[1].Datas)
	assert.Equal(t, []string{"data2", "data6", "data9"}, cargo.Columns[2].Datas)
}

func TestPkeyCheck(t *testing.T) {
	cargo := mockCargo()
	pkey := &Column{Datas: []string{"1", "2", "3", "4", "5"}}

	err := pkeyCheck(cargo, pkey)
	assert.Nil(t, err)

	pkey.Datas = append(pkey.Datas, "5")
	err = pkeyCheck(cargo, pkey)
	assert.NotNil(t, err)
}

func mockCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfNote:  1,
			LineOfField: 2,
			LineOfData:  3,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		},
	}
}
