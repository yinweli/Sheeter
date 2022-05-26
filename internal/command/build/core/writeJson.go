package core

import "fmt"

// WriteJson 寫入json
func WriteJson(task *Task) error {
	return nil
}

// buildJBoxes 建立json箱列表
func buildJBoxes(task *Task) (jboxes []jbox, err error) {
	for _, itor := range task.Columns {
		for row, data := range itor.Datas {
			result, err := itor.Field.Transform(data)

			if err != nil {
				return nil, fmt.Errorf("convert to json failed: %s [%s(%d) : %s]", task.Element.GetFullName(), itor.Name, row, err)
			} // if

			if len(jboxes) <= row {
				jboxes = append(jboxes, jbox{})
			} // if

			jboxes[row][itor.Name] = result
			_ = task.Progress.Add(1)
		} // for
	} // for

	return jboxes, nil
}

// jbox json箱
type jbox map[string]interface{}
