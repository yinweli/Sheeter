namespace testdata
{
    using System;
    using System.Collections.Generic;

    using System.Globalization;
    using Newtonsoft.Json;
    using Newtonsoft.Json.Converters;

    public partial class Struct
    {
        [JsonProperty("Reward")]
        public Reward Reward { get; set; }

        [JsonProperty("enable")]
        public bool Enable { get; set; }

        [JsonProperty("key")]
        public long Key { get; set; }

        [JsonProperty("name")]
        public string Name { get; set; }
    }

    public partial class Reward
    {
        [JsonProperty("Item")]
        public Item[] Item { get; set; }

        [JsonProperty("atium")]
        public long Atium { get; set; }

        [JsonProperty("crystal")]
        public long Crystal { get; set; }

        [JsonProperty("diamond")]
        public long Diamond { get; set; }

        [JsonProperty("felIron")]
        public long FelIron { get; set; }

        [JsonProperty("gold")]
        public long Gold { get; set; }
    }

    public partial class Item
    {
        [JsonProperty("count")]
        public long Count { get; set; }

        [JsonProperty("itemID")]
        public long ItemId { get; set; }

        [JsonProperty("type")]
        public long Type { get; set; }
    }
}
