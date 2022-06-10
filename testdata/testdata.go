package testdata

import (
	"github.com/xuri/excelize/v2"
)

const RealConfig = "real.yaml"
const Defect1Config = "config.defect1.yaml"
const Defect2Config = "config.defect2.yaml"
const RealExcel = "real.xlsx"
const Defect1Excel = "excel.defect1.xlsx"
const Defect2Excel = "excel.defect2.xlsx"
const Defect3Excel = "excel.defect3.xlsx"
const Defect4Excel = "excel.defect4.xlsx"
const Defect5Excel = "excel.defect5.xlsx"
const Defect6Excel = "excel.defect6.xlsx"
const Defect7Excel = "excel.defect7.xlsx"
const Defect8Excel = "excel.defect8.xlsx"
const Defect9Excel = "excel.defect9.xlsx"
const Defect10Excel = "excel.defect10.xlsx"
const Defect11Excel = "excel.defect11.xlsx"
const SheetName = "Data"
const Text = "this a string"

// GetTestExcel 取得測試excel
func GetTestExcel(name string) *excelize.File {
	excel, err := excelize.OpenFile(Path(name))

	if err != nil {
		return nil
	} // if

	return excel
}
