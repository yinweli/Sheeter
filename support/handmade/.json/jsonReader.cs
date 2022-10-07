// 以下是模板驗證用程式碼
// using區段可能與實際給的不一致, 要注意

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson {
    public partial class RewardReader {
        public string DataName() {
            return "reward";
        }

        public string DataExt() {
            return "json";
        }

        public string DataFile() {
            return "reward.json";
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<RewardStorer>(data);
            return Datas != null;
        }

        public long[] MergeData(string data) {
            var repeats = new List<long>();
            var tmpl = JsonConvert.DeserializeObject<RewardStorer>(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new RewardStorer();

            foreach (var itor in tmpl.Datas) {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<long, Reward> Data {
            get {
                return Datas.Datas;
            }
        }

        private RewardStorer Datas = null;
    }
}

// 以下是為了通過編譯的程式碼, 不可使用