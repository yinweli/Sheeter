// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Item {
        // 物品數量
        [JsonProperty("Count")]
        public long Count { get; set; }
        // 物品編號
        [JsonProperty("ItemID")]
        public long ItemID { get; set; }
        // 物品類型
        [JsonProperty("Type")]
        public long Type { get; set; }
    }

    public partial class ItemStorer {
        public Dictionary<long, Item> Datas = new Dictionary<long, Item>(); 
    }
}
