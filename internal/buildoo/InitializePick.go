package buildoo

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

	for _, itor := range context.element {
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

		// 由於這些組件都會在多執行緒環境下執行, 所以要注意不能有改變自身資料的操作出現
		named := &nameds.Named{ExcelName: data.Excel, SheetName: data.Sheet}
		json := &nameds.Json{ExcelName: data.Excel, SheetName: data.Sheet}
		proto := &nameds.Proto{ExcelName: data.Excel, SheetName: data.Sheet}

		if context.ExportJson {
			context.encoding = append(context.encoding, &encodingJson{
				Global:     context.Global,
				Named:      named,
				Json:       json,
				excel:      data.excel,
				layoutJson: data.layoutJson,
			})
		} // if

		if context.ExportProto {
			context.encoding = append(context.encoding, &encodingProto{
				Global:     context.Global,
				Named:      named,
				Proto:      proto,
				excel:      data.excel,
				layoutJson: data.layoutJson,
			})
		} // if
	} // for

	for _, itor := range layoutType.TypeNames() {
		// 由於這些組件都會在多執行緒環境下執行, 所以要注意不能有改變自身資料的操作出現
		types := layoutType.Types(itor)
		depend := layoutDepend.Depends(itor)
		named := &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet}
		field := &nameds.Field{}
		json := &nameds.Json{ExcelName: types.Excel, SheetName: types.Sheet}
		proto := &nameds.Proto{ExcelName: types.Excel, SheetName: types.Sheet}

		if context.ExportJson {
			context.generate = append(context.generate, &generateJson{
				Global: context.Global,
				Named:  named,
				Field:  field,
				Json:   json,
				Type:   types,
			})
			poststepJsonDepot.Struct = append(poststepJsonDepot.Struct, poststepJsonStruct{
				Named:  &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet},
				Reader: types.Reader,
			})
		} // if

		if context.ExportProto {
			context.generate = append(context.generate, &generateProto{
				Global: context.Global,
				Named:  named,
				Field:  field,
				Proto:  proto,
				Type:   types,
				Depend: depend,
			})
			context.poststep = append(context.poststep, &poststepConvert{
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
		context.poststep = append(context.poststep, poststepJsonDepot)
	} // if

	if context.ExportProto {
		context.poststep = append(context.poststep, poststepProtoDepot)
	} // if

	if context.Format {
		context.poststep = append(context.poststep, &poststepFormat{})
	} // if

	return nil
}
