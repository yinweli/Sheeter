// 以下是模板驗證用程式碼
// using區段可能與實際給的不一致, 要注意

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson {
    using Data_ = Reward;
    using PKey_ = System.Int64;
    using Storer_ = RewardStorer;

    public partial class RewardReader : Reader {
        public string DataName() {
            return "reward";
        }

        public string DataExt() {
            return "json";
        }

        public string DataFile() {
            return "reward.json";
        }

        public string FromData(string data) {
            Storer_ result;

            try {
                result = JsonConvert.DeserializeObject<Storer_>(data);
            } catch {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(string data) {
            Storer_ result;

            try {
                result = JsonConvert.DeserializeObject<Storer_>(data);
            } catch {
                return "merge data failed: deserialize failed";
            }

            if (result == null)
                return "merge data failed: result null";

            foreach (var itor in result.Datas) {
                if (storer.Datas.ContainsKey(itor.Key))
                    return "merge data failed: key repeat";

                storer.Datas[itor.Key] = itor.Value;
            }

            return string.Empty;
        }

        public void Clear() {
            storer.Datas.Clear();
        }

        public bool TryGetValue(PKey_ key, out Data_ value) {
            return storer.Datas.TryGetValue(key, out value);
        }

        public bool ContainsKey(PKey_ key) {
            return storer.Datas.ContainsKey(key);
        }

        public IEnumerator<KeyValuePair<PKey_, Data_>> GetEnumerator() {
            return storer.Datas.GetEnumerator();
        }

        public Data_ this[PKey_ key] {
            get {
                return storer.Datas[key];
            }
        }

        public ICollection<PKey_> Keys {
            get {
                return storer.Datas.Keys;
            }
        }

        public ICollection<Data_> Values {
            get {
                return storer.Datas.Values;
            }
        }

        public int Count {
            get {
                return storer.Datas.Count;
            }
        }

        private Storer_ storer = new Storer_();
    }
}

// 以下是為了通過編譯的程式碼, 不可使用