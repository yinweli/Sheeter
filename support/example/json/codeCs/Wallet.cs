﻿// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson
{
    public partial class Wallet
    {
        // 魔晶
        [JsonProperty("Crystal")]
        public long Crystal { get; set; }

        // 鑽石
        [JsonProperty("Diamond")]
        public long Diamond { get; set; }

        // 金幣
        [JsonProperty("Gold")]
        public long Gold { get; set; }
    }
}
