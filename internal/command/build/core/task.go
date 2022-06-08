package core

import "io"

// Task 執行工作
func Task(global *Global, element *Element, writer io.Writer) error {
	ctx := &Context{
		Global:  global,
		Element: element,
	}
	err := TaskExcel(ctx, writer)

	if err != nil {
		return err
	} // if

	err = TaskColumns(ctx)

	if err != nil {
		return err
	} // if

	err = TaskNotes(ctx)

	if err != nil {
		return err
	} // if

	err = TaskDatas(ctx)

	if err != nil {
		return err
	} // if

	err = TaskPkeyCheck(ctx)

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

	_ = ctx.Progress.Finish()
	return nil
}
