// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark03DataJ {
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

    public partial class Benchmark03DataJStorer {
        public Dictionary<long, Benchmark03DataJ> Datas = new Dictionary<long, Benchmark03DataJ>(); 
    }
}
