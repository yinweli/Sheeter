package core

import "Sheeter/internal/util"

// Task 執行工作
func Task(global *Global, element *Element) error {
	ctx := &Context{
		Global:  global,
		Element: element,
	}
	err := TaskExcel(ctx)
	defer util.SilentClose(ctx.Excel)

	if err != nil {
		return err
	} // if

	err = TaskFields(ctx)

	if err != nil {
		return err
	} // if

	err = TaskNotes(ctx)

	if err != nil {
		return err
	} // if

	err = TaskJson(ctx)

	if err != nil {
		return err
	} // if

	err = TaskJsonCpp(ctx)

	if err != nil {
		return err
	} // if

	err = TaskJsonCs(ctx)

	if err != nil {
		return err
	} // if

	err = TaskJsonGo(ctx)

	if err != nil {
		return err
	} // if

	return nil
}
