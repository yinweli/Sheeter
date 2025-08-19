package builds

import (
	"fmt"
	"sort"

	"github.com/yinweli/Sheeter/v3/sheeter/nameds"
	"github.com/yinweli/Sheeter/v3/sheeter/pipelines"
	"github.com/yinweli/Sheeter/v3/sheeter/tmpls"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Poststep 後製處理
func Poststep(config *Config, initializeData []*InitializeData) (result []any, err []error) {
	material := &PoststepData{
		Config: config,
		Named: &nameds.Named{ // 由於只需要AppName與Namespace, 所以不必填寫excel與sheet名稱
			Output: config.Output, // 但是需要表格器路徑, 所以要填寫輸出路徑
		},
	}

	for _, itor := range initializeData {
		material.Alone = append(material.Alone, &nameds.Named{
			ExcelName: itor.ExcelName,
			SheetName: itor.SheetName,
		})
	} // for

	sort.Slice(material.Alone, func(l, r int) bool { // 經過排序後讓產生程式碼時能夠更加一致
		lhs := material.Alone[l]
		rhs := material.Alone[r]
		return lhs.StructName() < rhs.StructName()
	})

	for _, itor := range config.Merged() {
		member := []*nameds.Named{}

		for _, term := range itor.Member() {
			excelName, sheetName := term.Name()
			member = append(member, &nameds.Named{
				ExcelName: excelName,
				SheetName: sheetName,
			})
		} // for

		material.Merge = append(material.Merge, &nameds.Merge{
			Name:   itor.Name(),
			Member: member,
		})
	} // for

	sort.Slice(material.Merge, func(l, r int) bool { // 經過排序後讓產生程式碼時能夠更加一致
		lhs := material.Merge[l]
		rhs := material.Merge[r]
		return lhs.Name < rhs.Name
	})

	result, err = pipelines.Pipeline[*PoststepData]("poststep", []*PoststepData{material}, []pipelines.Execute[*PoststepData]{
		generateSheeterCs,
		generateHelperCs,
		generateSheeterGo,
		generateHelperGo,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	return result, nil
}

// PoststepData 後製資料
type PoststepData struct {
	*Config                       // 設定資料
	*nameds.Named                 // 命名工具
	Alone         []*nameds.Named // 獨立結構列表
	Merge         []*nameds.Merge // 合併結構列表
}

// generateSheeterCs 產生cs語言表格程式碼
func generateSheeterCs(material *PoststepData) (result pipelines.Output) {
	path := material.SheeterPathCs()

	if err := utils.WriteTmpl(path, tmpls.SheeterCs, material); err != nil {
		result.Error = fmt.Errorf("generate sheeter cs: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}

// generateHelperCs 產生cs語言工具程式碼
func generateHelperCs(material *PoststepData) (result pipelines.Output) {
	path := material.HelperPathCs()

	if err := utils.WriteTmpl(path, tmpls.HelperCs, material); err != nil {
		result.Error = fmt.Errorf("generate helper cs: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}

// generateSheeterGo 產生go語言表格程式碼
func generateSheeterGo(material *PoststepData) (result pipelines.Output) {
	path := material.SheeterPathGo()

	if err := utils.WriteTmpl(path, tmpls.SheeterGo, material); err != nil {
		result.Error = fmt.Errorf("generate sheeter go: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}

// generateHelperGo 產生go語言工具程式碼
func generateHelperGo(material *PoststepData) (result pipelines.Output) {
	path := material.HelperPathGo()

	if err := utils.WriteTmpl(path, tmpls.HelperGo, material); err != nil {
		result.Error = fmt.Errorf("generate helper go: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}
