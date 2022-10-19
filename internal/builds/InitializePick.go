package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
)

// InitializePick 初始化揀選
func InitializePick(context *Context) error {
	layoutType := layouts.NewLayoutType()
	layoutDepend := layouts.NewLayoutDepend()
	poststepJsonDepot := &poststepJsonDepot{
		Global: context.Global,
		Named:  &nameds.Named{},
		Json:   &nameds.Json{},
	}
	poststepProtoDepot := &poststepProtoDepot{
		Global: context.Global,
		Named:  &nameds.Named{},
		Proto:  &nameds.Proto{},
	}

	for _, itor := range context.Element {
		data, ok := itor.(*initializeElement)

		if ok == false {
			return fmt.Errorf("initialize pick failed: cast failed")
		} // if

		if err := layoutType.Merge(data.layoutType); err != nil {
			return fmt.Errorf("initialize pick failed: %w", err)
		} // if

		if err := layoutDepend.Merge(data.layoutDepend); err != nil {
			return fmt.Errorf("initialize pick failed: %w", err)
		} // if

		if context.ExportJson {
			context.Encoding = append(context.Encoding, &encodingJson{
				Global:     context.Global,
				Named:      &nameds.Named{ExcelName: data.ExcelName, SheetName: data.SheetName},
				Json:       &nameds.Json{ExcelName: data.ExcelName, SheetName: data.SheetName},
				excel:      data.excel,
				layoutJson: data.layoutJson,
			})
		} // if

		if context.ExportProto {
			context.Encoding = append(context.Encoding, &encodingProto{
				Global:     context.Global,
				Named:      &nameds.Named{ExcelName: data.ExcelName, SheetName: data.SheetName},
				Proto:      &nameds.Proto{ExcelName: data.ExcelName, SheetName: data.SheetName},
				excel:      data.excel,
				layoutJson: data.layoutJson,
			})
		} // if
	} // for

	for _, itor := range layoutType.TypeNames() {
		// 由於這些組件都會在多執行緒環境下執行, 所以要注意不能有改變自身資料的操作出現
		types := layoutType.Types(itor)
		depend := layoutDepend.Depends(itor)

		if context.ExportJson {
			context.Generate = append(context.Generate, &generateJson{
				Global: context.Global,
				Named:  &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet},
				Field:  &nameds.Field{},
				Json:   &nameds.Json{ExcelName: types.Excel, SheetName: types.Sheet},
				Type:   types,
			})
			poststepJsonDepot.Struct = append(poststepJsonDepot.Struct, poststepJsonStruct{
				Named:  &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet},
				Reader: types.Reader,
			})
		} // if

		if context.ExportProto {
			context.Generate = append(context.Generate, &generateProto{
				Global: context.Global,
				Named:  &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet},
				Field:  &nameds.Field{},
				Proto:  &nameds.Proto{ExcelName: types.Excel, SheetName: types.Sheet},
				Type:   types,
				Depend: depend,
			})
			proto := &nameds.Proto{ExcelName: types.Excel, SheetName: types.Sheet}
			context.Poststep = append(context.Poststep, &poststepConvert{
				include:  proto.ProtoSchemaPath(),
				outputCs: proto.ProtoCsPath(),
				outputGo: proto.ProtoGoPath(),
				source:   proto.ProtoPath(),
			})
			poststepProtoDepot.Struct = append(poststepProtoDepot.Struct, poststepProtoStruct{
				Named:  &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet},
				Reader: types.Reader,
			})
		} // if
	} // for

	if context.ExportJson {
		context.Poststep = append(context.Poststep, poststepJsonDepot)
	} // if

	if context.ExportProto {
		context.Poststep = append(context.Poststep, poststepProtoDepot)
	} // if

	if context.Format {
		context.Poststep = append(context.Poststep, &poststepFormat{})
	} // if

	return nil
}
