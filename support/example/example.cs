using Newtonsoft.Json;
using System;
using System.IO;

namespace Example
{
    public static class Example
    {
        public static void Main()
        {
            exampleJson();
            exampleProto();
            exampleEnum();
        }

        // 在多執行緒環境下執行時, 各個表格物件的取用資料內容操作都是執行緒安全的
        // 但是倉庫物件的FromData與MergeData操作則是非執行緒安全的, 請注意此點

        /// <summary>
        /// json範例
        /// </summary>
        private static void exampleJson()
        {
            // 要使用sheeter, 首先建立繼承自sheeterJson.Loader介面的讀取器
            // 讀取器負責從磁碟(或是其他的資料來源)取得資料的流程, 這部分由使用者自行處理
            // 範例中的讀取器只是簡單的從磁碟讀取檔案而已
            var loader = new JsonFileLoader();
            // 接著建立sheeterJson.Depot物件, 這是存取表格資料最主要的物件
            // 要記得把剛剛建立的讀取器設定進去
            var depot = new SheeterJson.Depot() { Loader = loader };

            // 然後執行FromData(或是MergeData)函式來讀取表格資料
            if (depot.FromData() == false)
            {
                Console.WriteLine("json failed: from data failed");
                return;
            }

            // 之後就可以用Depot底下的各個表格物件來取用資料內容
            if (depot.ExampleData.TryGetValue(1, out var data))
            {
                Console.WriteLine(JsonConvert.SerializeObject(data));
                Console.WriteLine("json success");
            }
            else
            {
                Console.WriteLine("json failed: get data failed");
            }
        }

        /// <summary>
        /// proto範例
        /// </summary>
        private static void exampleProto()
        {
            // 要使用sheeter, 首先建立繼承自sheeterProto.Loader介面的讀取器
            // 讀取器負責從磁碟(或是其他的資料來源)取得資料的流程, 這部分由使用者自行處理
            // 範例中的讀取器只是簡單的從磁碟讀取檔案而已
            var loader = new ProtoFileLoader();
            // 接著建立sheeterProto.Depot物件, 這是存取表格資料最主要的物件
            // 要記得把剛剛建立的讀取器設定進去
            var depot = new SheeterProto.Depot() { Loader = loader };

            // 然後執行FromData(或是MergeData)函式來讀取表格資料
            if (depot.FromData() == false)
            {
                Console.WriteLine("proto failed: from data failed");
                return;
            }

            // 之後就可以用Depot底下的各個表格物件來取用資料內容
            if (depot.ExampleData.TryGetValue(1, out var data))
            {
                Console.WriteLine(data);
                Console.WriteLine("proto success");
            }
            else
            {
                Console.WriteLine("proto failed: get data failed");
            }
        }

        /// <summary>
        /// enum範例
        /// </summary>
        private static void exampleEnum()
        {
            // 列舉就直接使用就好
            Console.WriteLine(SheeterEnum.ExampleEnum.Name0);
            Console.WriteLine(SheeterEnum.ExampleEnum.Name1);
            Console.WriteLine(SheeterEnum.ExampleEnum.Name2);
            Console.WriteLine("enum success");
        }
    }

    /// <summary>
    /// json檔案讀取器
    /// </summary>
    class JsonFileLoader : SheeterJson.Loader
    {
        /// <summary>
        /// 用於處理讀取資料錯誤, 範例中只是單純印出錯誤訊息
        /// </summary>
        /// <param name="name">檔案名稱</param>
        /// <param name="message">錯誤訊息</param>
        public void Error(string name, string message)
        {
            Console.WriteLine(name + ": json file load failed: " + message);
        }

        /// <summary>
        /// 用於讀取資料檔案, Depot會提供給你檔案名稱, 副檔名, 完整名稱
        /// 使用者需要依靠以上資訊來讀取資料檔案, 並回傳資料給Depot
        /// </summary>
        /// <param name="filename">檔名物件</param>
        /// <returns>檔案內容</returns>
        public string Load(SheeterJson.FileName filename)
        {
            // 因為工作路徑在 bin/Debug/net5.0/ 底下,所以只好加3個".."了
            return File.ReadAllText(Path.Combine("..", "..", "..", "json", "data", filename.File));
        }
    }

    /// <summary>
    /// proto檔案讀取器
    /// </summary>
    class ProtoFileLoader : SheeterProto.Loader
    {
        /// <summary>
        /// 用於處理讀取資料錯誤, 範例中只是單純印出錯誤訊息
        /// </summary>
        /// <param name="name">檔案名稱</param>
        /// <param name="message">錯誤訊息</param>
        public void Error(string name, string message)
        {
            Console.WriteLine(name + ": proto file load failed: " + message);
        }

        /// <summary>
        /// 用於讀取資料檔案, Depot會提供給你檔案名稱, 副檔名, 完整名稱
        /// 使用者需要依靠以上資訊來讀取資料檔案, 並回傳資料給Depot
        /// </summary>
        /// <param name="filename">檔名物件</param>
        /// <returns>檔案內容</returns>
        public byte[] Load(SheeterProto.FileName filename)
        {
            // 因為工作路徑在 bin/Debug/net5.0/ 底下,所以只好加3個".."了
            return File.ReadAllBytes(
                Path.Combine("..", "..", "..", "proto", "data", filename.File)
            );
        }
    }
}
