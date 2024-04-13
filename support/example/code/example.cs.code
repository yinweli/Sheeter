using Newtonsoft.Json;

namespace example
{
    public static class Example
    {
        // 請先安裝Newtonsoft.Json, sheeter中使用此函式庫進行json操作

        public static void Main()
        {
            Usecase();
        }

        // 在多執行緒環境下執行時, 各個表格物件的取用資料內容操作都是執行緒安全的
        // 但是FromData與MergeData操作則是非執行緒安全的, 請注意此點

        /// <summary>
        /// 使用範例
        /// </summary>
        private static void Usecase()
        {
            // 要使用sheeter, 首先建立繼承自Sheeter.Loader介面的裝載器
            // 裝載器負責從磁碟(或是其他的資料來源)讀取資料, 這部分由使用者自行處理
            // 範例中的裝載器只是簡單的從磁碟讀取檔案而已
            var loader = new FileLoader();
            // 接著建立Sheeter.Sheeter物件, 這是存取表格資料最主要的物件
            // 要記得把剛剛建立的裝載器設定進去
            var sheet = new Sheeter.Sheeter(loader);

            // 然後執行FromData(或是MergeData)函式來讀取表格資料
            if (sheet.FromData() == false)
            {
                Console.WriteLine("load failed: from data failed");
                return;
            }

            // 之後就可以用Sheeter.Sheeter底下的各個表格物件來取用資料內容
            if (sheet.VerifyData.TryGetValue(1, out var data1))
            {
                Console.WriteLine(JsonConvert.SerializeObject(data1));
                Console.WriteLine("get data success");
            }
            else
            {
                Console.WriteLine("get data failed");
            }

            if (sheet.VerifyData.TryGetValue(2, out var data2))
            {
                Console.WriteLine(JsonConvert.SerializeObject(data2));
                Console.WriteLine("get data success");
            }
            else
            {
                Console.WriteLine("get data failed");
            }
        }

        /// <summary>
        /// 裝載器
        /// </summary>
        class FileLoader : Sheeter.Loader
        {
            /// <summary>
            /// 用於讀取資料檔案, sheeter提供給你FileName物件
            /// 使用者依靠FileName的功能取得檔名來讀取資料, 並回傳檔案內容給sheeter
            /// </summary>
            /// <param name="filename">檔名物件</param>
            /// <returns>檔案內容</returns>
            public string Load(Sheeter.FileName filename)
            {
                return File.ReadAllText(Path.Combine("json", filename.File)); // 資料檔案放在json目錄下
            }

            /// <summary>
            /// 用於處理錯誤, 範例中只是單純印出錯誤訊息
            /// </summary>
            /// <param name="name">檔案名稱</param>
            /// <param name="message">錯誤訊息</param>
            public void Error(string name, string message)
            {
                Console.WriteLine("file load failed: " + name + ": " + message);
            }
        }
    }
}
