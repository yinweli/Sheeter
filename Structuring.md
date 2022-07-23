layerParse => layout

{name1
{[]name1
{name1 {name2
{[]name1 {name2
{name1 {[]name2
{[]name1 {[]name2
{name1 }
{[]name1 }
{name1 } }
{[]name1 } }
}
} }

+------+------+------+------+------+------+------+------+------+------+------+------+
|col1  |col2  |col3  |col4  |col5  |col6  |col7  |col8  |col9  |col10 |col11 |col12 |
+------+------+------+------+------+------+------+------+------+------+------+------+
|       {d     {[]n        } {[]n      } } {j          } {p                      } }|
+------+------+------+------+------+------+------+------+------+------+------+------+
|       {d     {[]n {i     }        {j          }      } {p                      } }|
+------+------+------+------+------+------+------+------+------+------+------+------+

d n p i j 這些都是結構(節點)名稱, column其實只要記錄他屬於哪個結構(節點)就好
當然也有不屬於結構(節點)的column, 那就是屬於root節點的
結構(節點)不能重複
還是要有地方紀錄結構(節點)的上下屬資料, 才有辦法做layout
最後layout要有能力產出可以裝資料的結構(或是投入資料, 他幫你組裝並丟出物件)

{
    "col1": 0,
    "d": {
        "col2": 0,
        "n": [
            {
                "col3": 0,
                "col4": 0
            },
            {
                "col5": 0,
                "col6": 0
            }
        ]
    },
    "j": {
        "col7": 0,
        "col8": 0
    },
    "p": {
        "col9": 0,
        "col10": 0,
        "col11": 0,
        "col12": 0
    }
}

{
    "col1": 0,
    "d": {
        "col2": 0,
        "n": [
            {
                "i": {
                    "col3": 0,
                    "col4": 0
                },
                "col5": 0,
                "j": {
                    "col6": 0,
                    "col7": 0
                },
                "col8": 0
            }
        ]
    },
    "p": {
        "col9": 0,
        "col10": 0,
        "col11": 0,
        "col12": 0
    }
}