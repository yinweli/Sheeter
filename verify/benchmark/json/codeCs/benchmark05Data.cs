// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark05Data {
        // 
        [JsonProperty("Reward")]
        public Reward Reward { get; set; }
        // 是否啟用
        [JsonProperty("Enable")]
        public bool Enable { get; set; }
        // 索引
        [JsonProperty("Key")]
        public long Key { get; set; }
        // 名稱
        [JsonProperty("Name")]
        public string Name { get; set; }
    }

    public partial class Benchmark05DataStorer {
        public Dictionary<long, Benchmark05Data> Datas = new Dictionary<long, Benchmark05Data>(); 
    }
}
