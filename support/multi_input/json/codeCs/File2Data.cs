﻿// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson
{
    public partial class File2Data
    {
        //
        [JsonProperty("Reward")]
        public Reward Reward { get; set; }

        // 是否啟用
        [JsonProperty("Enable")]
        public bool Enable { get; set; }

        // 忽略
        [JsonProperty("Ignore")]
        public int Ignore { get; set; }

        // 索引
        [JsonProperty("Key")]
        public System.Int32 Key { get; set; }

        // 名稱
        [JsonProperty("Name")]
        public string Name { get; set; }
    }

    public partial class File2DataStorer
    {
        public Dictionary<System.Int32, File2Data> Datas =
            new Dictionary<System.Int32, File2Data>();
    }
}
