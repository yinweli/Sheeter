// 以下是模板驗證用程式碼

using Newtonsoft.Json;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Reward {
        //
        [JsonProperty("Item")]
        public Item[] Item { get; set; }
        // 天金
        [JsonProperty("Atium")]
        public long Atium { get; set; }
        // 魔晶
        [JsonProperty("Crystal")]
        public long Crystal { get; set; }
        // 獎勵說明
        [JsonProperty("Desc")]
        public string Desc { get; set; }
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

    public partial class RewardStorer {
        public Dictionary<long, Reward> Datas = new Dictionary<long, Reward>();
    }
}

// 以下是為了通過編譯的程式碼, 不可使用

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
}