// 以下是模板驗證用程式碼

using System.Collections.Generic;

namespace SheeterJson {
    public partial class Depot {
        public Loader Loader { get; set; }
        public readonly RewardReader Reward = new RewardReader();
        private readonly List<Reader> Readers = new List<Reader>();

        public Depot() {
            Readers.Add(Reward);
        }

        public bool FromData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }
    }

    public interface Loader {
        public void Error(string name, string message);
        public string Load(string name, string ext, string fullname);
    }

    public interface Reader {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(string data);
        public string MergeData(string data);
    }
}

// 以下是為了通過編譯的程式碼, 不可使用