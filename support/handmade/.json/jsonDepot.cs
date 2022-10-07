// 以下是模板驗證用程式碼

using System.Collections.Generic;

namespace SheeterJson {
    public partial class Depot {
        public readonly RewardReader Reward = new RewardReader();

        public string[] FromData(DelegateFromData func) {
            var errors = new List<string>();

            foreach (var itor in readers) {
                var data = func(itor.DataName(), itor.DataExt());

                if (data != null && itor.FromData(data) == false) {
                    errors.Add(itor.DataFile());
                }
            }



            return errors.ToArray();
        }

        public delegate string DelegateFromData(string name, string ext);

        public string[] MergeData(DelegateMergeData func) {
            var errors = new List<string>();

            return errors.ToArray();
        }

        public delegate byte[] DelegateMergeData(string name, string ext);

        public Depot() {
            readers.Add(Reward);
        }

        private readonly List<ReaderInterface> readers = new List<ReaderInterface>();
    }

    public interface ReaderInterface {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public bool FromData(string data);
        public long[] MergeData(string data);
    }
}

// 以下是為了通過編譯的程式碼, 不可使用