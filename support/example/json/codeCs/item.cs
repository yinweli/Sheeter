﻿// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson
{
    public partial class Item
    {
        // 物品數量
        [JsonProperty("Count")]
        public int Count { get; set; }

        // 物品編號
        [JsonProperty("ItemID")]
        public int ItemID { get; set; }
    }
}
