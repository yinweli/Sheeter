// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;

namespace Sheeter
{
    using Data_ = VerifyData;
    using Key_ = System.Int32;
    using Store_ = Dictionary<System.Int32, VerifyData>;

    /// <summary>
    /// VerifyData verify.xlsx#Data
    /// </summary>
    public partial class VerifyData
    {
        /// <summary>
        /// note2
        /// </summary>
        [JsonProperty("Name2")]
        public int Name2 { get; set; }

        /// <summary>
        /// note3
        /// </summary>
        [JsonProperty("Name3")]
        public int Name3 { get; set; }

        /// <summary>
        /// note4
        /// </summary>
        [JsonProperty("Name4")]
        public int Name4 { get; set; }

        /// <summary>
        /// pkey
        /// </summary>
        [JsonProperty("Pkey")]
        public System.Int32 Pkey { get; set; }

        /// <summary>
        /// note1
        /// </summary>
        [JsonProperty("Name1")]
        public int Name1 { get; set; }
    }

    /// <summary>
    /// VerifyData verify.xlsx#Data
    /// </summary>
    public partial class VerifyDataReader : Reader
    {
        /// <summary>
        /// 取得檔名物件
        /// </summary>
        public FileName FileName()
        {
            return new FileName("verifyData", ".json");
        }

        /// <summary>
        /// 讀取資料
        /// </summary>
        public string FromData(string data)
        {
            try
            {
                Data = JsonConvert.DeserializeObject<Store_>(data);
            } // try
            catch
            {
                return "from data: deserialize failed";
            } // catch

            if (Data == null)
                return "from data: deserialize failed";

            return string.Empty;
        }

        /// <summary>
        /// 合併資料
        /// </summary>
        public string MergeData(string data)
        {
            Store_ tmpl;

            try
            {
                tmpl = JsonConvert.DeserializeObject<Store_>(data);
            } // try
            catch
            {
                return "merge data: deserialize failed";
            } // catch

            if (tmpl == null)
                return "merge data: deserialize failed";

            foreach (var itor in tmpl)
            {
                if (Data.ContainsKey(itor.Key))
                    return "merge data: key duplicate";

                Data[itor.Key] = itor.Value;
            } // for

            return string.Empty;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            Data.Clear();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public bool TryGetValue(Key_ key, out Data_ value)
        {
            return Data.TryGetValue(key, out value);
        }

        /// <summary>
        /// 檢查索引是否存在
        /// </summary>
        public bool ContainsKey(Key_ key)
        {
            return Data.ContainsKey(key);
        }

        /// <summary>
        /// 取得迭代器
        /// </summary>
        public IEnumerator<KeyValuePair<Key_, Data_>> GetEnumerator()
        {
            return Data.GetEnumerator();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public Data_ this[Key_ key]
        {
            get { return Data[key]; }
        }

        /// <summary>
        /// 取得索引列表
        /// </summary>
        public ICollection<Key_> Keys
        {
            get { return Data.Keys; }
        }

        /// <summary>
        /// 取得資料列表
        /// </summary>
        public ICollection<Data_> Values
        {
            get { return Data.Values; }
        }

        /// <summary>
        /// 取得資料數量
        /// </summary>
        public int Count
        {
            get { return Data.Count; }
        }

        private Store_ Data = new Store_();
    }
}
