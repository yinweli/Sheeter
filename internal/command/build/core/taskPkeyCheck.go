package core

import "fmt"

// TaskPkeyCheck 主要索引檢查
func TaskPkeyCheck(ctx *Context) error {
	datas := make(map[string]bool)

	for _, itor := range ctx.Pkey.Datas {
		if _, exist := datas[itor]; exist {
			return fmt.Errorf("pkey duplicate: %s", ctx.LogName())
		} // if

		datas[itor] = true
	} // for

	_ = ctx.Progress.Add(1)
	return nil
}
