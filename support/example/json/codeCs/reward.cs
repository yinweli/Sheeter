﻿// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson
{
    public partial class Reward
    {
        //
        [JsonProperty("Item")]
        public Item[] Item { get; set; }

        // 天金
        [JsonProperty("Atium")]
        public long Atium { get; set; }

        // 魔晶
        [JsonProperty("Crystal")]
        public long Crystal { get; set; }

        // 鑽石
        [JsonProperty("Diamond")]
        public long Diamond { get; set; }

        // 精鐵
        [JsonProperty("FelIron")]
        public long FelIron { get; set; }

        // 金幣
        [JsonProperty("Gold")]
        public long Gold { get; set; }
    }
}
