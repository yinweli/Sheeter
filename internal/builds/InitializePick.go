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
		if data, ok := itor.(*initializeElement); ok {
			if err := layoutType.Merge(data.layoutType); err != nil {
				return fmt.Errorf("initialize pick failed: %w", err)
			} // if

			if err := layoutDepend.Merge(data.layoutDepend); err != nil {
				return fmt.Errorf("initialize pick failed: %w", err)
			} // if

			named := &nameds.Named{ExcelName: data.ExcelName, SheetName: data.SheetName}
			json := &nameds.Json{ExcelName: data.ExcelName, SheetName: data.SheetName}
			proto := &nameds.Proto{ExcelName: data.ExcelName, SheetName: data.SheetName}

			if context.ExportJson {
				context.Encoding = append(context.Encoding, &encodingJson{
					Global:     context.Global,
					Named:      named,
					Json:       json,
					excel:      data.excel,
					layoutData: data.layoutData,
				})
			} // if

			if context.ExportProto {
				context.Encoding = append(context.Encoding, &encodingProto{
					Global:     context.Global,
					Named:      named,
					Proto:      proto,
					excel:      data.excel,
					layoutData: data.layoutData,
				})
			} // if
		} // if

		if data, ok := itor.(*initializeEnum); ok {
			named := &nameds.Named{ExcelName: data.ExcelName, SheetName: data.SheetName}
			enum := &nameds.Enum{ExcelName: data.ExcelName, SheetName: data.SheetName}

			if context.ExportEnum {
				context.Generate = append(context.Generate, &generateEnum{
					Global: context.Global,
					Named:  named,
					Enum:   enum,
					Enums:  data.layoutEnum.Enums(),
				})
				context.Poststep = append(context.Poststep, &poststepConvert{
					include:  enum.EnumSchemaPath(),
					outputCs: enum.EnumCsPath(),
					outputGo: enum.EnumGoPath(),
					source:   enum.EnumPath(),
				})
			} // if
		} // if
	} // for

	for _, itor := range layoutType.TypeNames() {
		types := layoutType.Types(itor)
		named := &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet}
		field := &nameds.Field{}
		json := &nameds.Json{ExcelName: types.Excel, SheetName: types.Sheet}
		proto := &nameds.Proto{ExcelName: types.Excel, SheetName: types.Sheet}
		depend := layoutDepend.Depends(itor)

		if context.ExportJson {
			context.Generate = append(context.Generate, &generateJson{
				Global: context.Global,
				Named:  named,
				Field:  field,
				Json:   json,
				Type:   types,
			})
			poststepJsonDepot.Struct = append(poststepJsonDepot.Struct, poststepJsonStruct{
				Named:  named,
				Reader: types.Reader,
			})
		} // if

		if context.ExportProto {
			context.Generate = append(context.Generate, &generateProto{
				Global: context.Global,
				Named:  named,
				Field:  field,
				Proto:  proto,
				Type:   types,
				Depend: depend,
			})
			context.Poststep = append(context.Poststep, &poststepConvert{
				include:  proto.ProtoSchemaPath(),
				outputCs: proto.ProtoCsPath(),
				outputGo: proto.ProtoGoPath(),
				source:   proto.ProtoPath(),
			})
			poststepProtoDepot.Struct = append(poststepProtoDepot.Struct, poststepProtoStruct{
				Named:  named,
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

	return nil
}
