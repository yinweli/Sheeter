package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadContent(t *testing.T) {
	cargo := mockReadContentCargo()
	err := ReadContent(cargo)
	assert.Nil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets[cargo.Global.GetLineOfField()] = []string{"name0#????", "name1#bool", "name2#int"}
	err = ReadContent(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets = Sheets{{"name0#pkey", "name1#int", "name2#int"}}
	err = ReadContent(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets[4][0] = "2"
	err = ReadContent(cargo)
	assert.NotNil(t, err)
}

func TestBuildColumns(t *testing.T) {
	cargo := mockReadContentCargo()
	pkey, err := buildColumns(cargo)
	assert.Nil(t, err)
	assert.NotNil(t, pkey)
	assert.Equal(t, "name0", pkey.Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), pkey.Field.TypeExcel())
	assert.Equal(t, 3, len(cargo.Columns))
	assert.Equal(t, "name0", cargo.Columns[0].Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), cargo.Columns[0].Field.TypeExcel())
	assert.Equal(t, "name1", cargo.Columns[1].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), cargo.Columns[1].Field.TypeExcel())
	assert.Equal(t, "name2", cargo.Columns[2].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), cargo.Columns[2].Field.TypeExcel())

	cargo = mockReadContentCargo()
	cargo.Sheets[cargo.Global.GetLineOfField()] = []string{"name0#????", "name1#bool", "name2#int"}
	pkey, err = buildColumns(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets[cargo.Global.GetLineOfField()] = []string{"name0#pkey", "name1#pkey", "name2#int"}
	pkey, err = buildColumns(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets[cargo.Global.GetLineOfField()] = []string{"name0#int", "name1#int", "name2#int"}
	pkey, err = buildColumns(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets[cargo.Global.GetLineOfField()] = []string{"", "", ""}
	pkey, err = buildColumns(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets[cargo.Global.GetLineOfField()] = []string{}
	pkey, err = buildColumns(cargo)
	assert.NotNil(t, err)

	cargo = mockReadContentCargo()
	cargo.Sheets = Sheets{}
	pkey, err = buildColumns(cargo)
	assert.NotNil(t, err)
}

func TestBuildNotes(t *testing.T) {
	cargo := mockReadContentCargo()
	err := buildNotes(cargo)
	assert.Nil(t, err)
	assert.Equal(t, "note0", cargo.Columns[0].Note)
	assert.Equal(t, "note1", cargo.Columns[1].Note)
	assert.Equal(t, "note2", cargo.Columns[2].Note)

	cargo = mockReadContentCargo()
	cargo.Sheets = Sheets{}
	err = buildNotes(cargo)
	assert.NotNil(t, err)
}

func TestBuildDatas(t *testing.T) {
	cargo := mockReadContentCargo()
	err := buildDatas(cargo)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, cargo.Columns[0].Datas)
	assert.Equal(t, []string{"1", "2", "3"}, cargo.Columns[1].Datas)
	assert.Equal(t, []string{"1", "2", "3"}, cargo.Columns[2].Datas)

	cargo = mockReadContentCargo()
	cargo.Sheets = Sheets{}
	err = buildDatas(cargo)
	assert.Nil(t, err)
}

func TestPkeyCheck(t *testing.T) {
	cargo := mockReadContentCargo()
	pkey := &Column{Datas: []string{"1", "2", "3"}}
	err := pkeyCheck(cargo, pkey)
	assert.Nil(t, err)

	pkey.Datas = append(pkey.Datas, "3")
	err = pkeyCheck(cargo, pkey)
	assert.NotNil(t, err)
}

func mockReadContentCargo() *Cargo {
	return &Cargo{
		Progress: NewProgress(0, "test", ioutil.Discard),
		Global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfField: 1,
			LineOfNote:  2,
			LineOfData:  3,
		},
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Sheets: Sheets{
			{"name0#pkey", "name1#int", "name2#int"},
			{"note0", "note1", "note2"},
			{"1", "1", "1"},
			{"2", "2", "2"},
			{"3", "3", "3"},
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldPkey{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}},
		},
	}
}
