// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System.Collections.Generic;

namespace sheeter {
    public partial class Benchmark04Data {
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

    public partial class Benchmark04DataStorer {
        public Dictionary<long, Benchmark04Data> Datas = new Dictionary<long, Benchmark04Data>(); 
    }
}
