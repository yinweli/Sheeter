package core

// TODO: d n p i j 這些都是結構(節點)名稱, column其實只要記錄他屬於哪個結構(節點)就好
//       當然也有不屬於結構(節點)的column, 那就是屬於root節點的
//       結構(節點)不可以同名卻不同類型(陣列/結構)
//       還是要有地方紀錄結構(節點)的上下屬資料, 才有辦法做layout
//       最後layout要有能力產出可以裝資料的結構(或是投入資料, 他幫你組裝並丟出物件)

// TODO: layer stack
// TODO: layout stack

/*
layout stack: 後進先出堆疊
    先加一個 root map 到堆疊中
    然後依照布局增加陣列/map進去堆疊去
    back的時候會從堆疊中移除最後元素

+------+------+------+------+------+------+------+------+------+------+------+------+
|col1  |col2  |col3  |col4  |col5  |col6  |col7  |col8  |col9  |col10 |col11 |col12 |
+------+------+------+------+------+------+------+------+------+------+------+------+
|       {d     {[]n          /          }} {j          } {p                       }}|
+------+------+------+------+------+------+------+------+------+------+------+------+
|       {d     {[]n   {i          } {j          }      } {p                       }}|
+------+------+------+------+------+------+------+------+------+------+------+------+
|{[]d   {x          }        {y         }} /      {x          }        {y         }}|
+------+------+------+------+------+------+------+------+------+------+------+------+

layer: before, self, after

開始之前要加入rootMap                # map { }
col1: 最後元素新增 col1:值           # map { col1 }
col2: 最後元素新增 d:map             # map { col1, d map { } }
      最後元素新增 col2:值           # map { col1, d map { col2 } }
col3: 最後元素新增 n:陣列            # map { col1, d map { col2, n[] { } } }
      最後元素新增 陣列元素          # map { col1, d map { col2, n[] { map { } } } }
      最後元素新增 col3:值           # map { col1, d map { col2, n[] { map { col3 } } } }
col4: 最後元素新增 col4:值           # map { col1, d map { col2, n[] { map { col3, col4 } } } }
col5: 最後元素新增 陣列元素          # map { col1, d map { col2, n[] { map { col3, col4 }, map { } } } }
      最後元素新增 col5:值           # map { col1, d map { col2, n[] { map { col3, col4 }, map { col5 } } } }
col6: 最後元素新增 col6:值           # map { col1, d map { col2, n[] { map { col3, col4 }, map { col5, col6 } } } }
      彈出最後元素                   # map { col1, d map { col2, n[] { } } }
      因為最後元素是陣列, 多彈出一次 # map { col1, d map { col2 } }
      彈出最後元素                   # map { col1 }
*/

// Builder 布局建立器
type Builder struct {
	duplField duplField // 欄位重複檢查器
	duplLayer duplLayer // 階層重複檢查器
	layouts   []layout  // 布局列表
}

// Add 新增布局
func (this *Builder) Add() {

}

// Pack 打包資料
func (this *Builder) Pack() map[string]interface{} {
	return nil
}

// layout 布局資料
type layout struct {
	name   string
	action int
}

// duplField 欄位重複檢查器
type duplField struct {
	datas map[string]bool // 資料列表
}

// Check 重複檢查
func (this *duplField) Check(field string) bool {
	if _, ok := this.datas[field]; ok {
		return false
	} // if

	this.datas[field] = true
	return true
}

// duplLayer 階層重複檢查器
type duplLayer struct {
	datas map[string]int // 資料列表
}

// Check 重複檢查
func (this *duplLayer) Check(layers ...Layer) bool {
	for _, itor := range layers {
		if type_, ok := this.datas[itor.Name]; ok {
			return type_ == itor.Type
		} // if

		this.datas[itor.Name] = itor.Type
	} // for

	return true
}
