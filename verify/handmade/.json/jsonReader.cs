// 以下是模板驗證用程式碼

using Newtonsoft.Json;

namespace sheeterJson {

    public partial class RewardReader {

        public static string FileName()
        {
            return "reward.json";
        }

        public bool FromPath(string path)
        {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data)
        {
            Datas = JsonConvert.DeserializeObject<RewardStorer>(data);
            return Datas != null;
        }

        public long[] MergePath(params string[] path)
        {
            var repeats = new List<long>();

            foreach (var itor in path)
            {
                try
                {
                    repeats.AddRange(MergeData(File.ReadAllText(Path.Combine(itor, FileName()))));
                }
                catch
                {
                    // do nothing
                }
            }

            return repeats.ToArray();
        }

        public long[] MergeData(string data)
        {
            var repeats = new List<long>();
            var tmpl = JsonConvert.DeserializeObject<RewardStorer>(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new RewardStorer();

            foreach (var itor in tmpl.Datas)
            {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<long, Reward> Data
        {
            get
            {
                return Datas.Datas;
            }
        }

        private RewardStorer Datas = null;
    }
}

// 以下是為了通過編譯的程式碼, 不可使用