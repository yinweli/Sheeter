// 以下是模板驗證用程式碼

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class RewardReader {
        public static string FileName() {
            return "reward.json";
        }

        public bool FromPath(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<RewardStorer>(data);
            return Datas != null;
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