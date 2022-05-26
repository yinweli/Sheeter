package core

import "fmt"

// WriteJson 寫入json
func WriteJson(cargo *Cargo) error {
	return nil
}

// buildJBoxes 建立json箱列表
func buildJBoxes(cargo *Cargo) (jboxes []jbox, err error) {
	for _, itor := range cargo.Columns {
		for row, data := range itor.Datas {
			result, err := itor.Field.Transform(data)

			if err != nil {
				return nil, fmt.Errorf("convert to json failed: %s [%s(%d) : %s]", cargo.Element.GetFullName(), itor.Name, row, err)
			} // if

			if len(jboxes) <= row {
				jboxes = append(jboxes, jbox{})
			} // if

			jboxes[row][itor.Name] = result
			_ = cargo.Progress.Add(1)
		} // for
	} // for

	return jboxes, nil
}

// jbox json箱
type jbox map[string]interface{}
