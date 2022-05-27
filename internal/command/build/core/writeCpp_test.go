package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"Sheeter/internal/util"
	"Sheeter/testdata"
)

func TestWrite(t *testing.T) {
	tmpl, err := template.New("test").Parse(templateCpp)

	if err != nil {
		fmt.Println(err)
		return
	} // if

	cargo := &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Global: &Global{
			CppLibraryPath: "testme.h",
			CppNamespace:   "sheeter",
			CsNamespace:    "sheeter",
			GoPackage:      "sheeter",
			ExcelPath:      testdata.RootPath,
			LineOfNote:     1,
			LineOfField:    2,
			LineOfData:     3,
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

	err = tmpl.Execute(os.Stdout, cargo)

	if err != nil {
		fmt.Println(err)
		return
	} // if

}
