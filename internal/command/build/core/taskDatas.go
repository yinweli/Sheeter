package core

// TaskDatas 建立資料
func TaskDatas(ctx *Context) error {
	for _, itor := range ctx.Columns {
		itor.Datas = []string{} // 把資料列表清空, 避免不必要的問題
	} // for

	rowSize := len(ctx.Sheets)

	// TODO: 結果花時間的地方在這邊

	for row := ctx.Global.GetLineOfData(); row < rowSize; row++ {
		colSize := len(ctx.Sheets[row])

		for col, itor := range ctx.Columns {
			var data string

			if col >= 0 && col < colSize { // 資料行的數量可能因為空白格的關係會短缺, 所以要檢查一下
				data = ctx.Sheets[row][col]
			} // if

			_ = ctx.Progress.Add(1)
			itor.Datas = append(itor.Datas, data)
		} // for
	} // for

	return nil
}
