// 以下是模板驗證用程式碼

using System.Collections.Generic;

namespace SheeterProto {
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
                var filename = itor.FileName();
                var data = Loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(filename.File, message);
                }
            }

            return result;
        }

        public bool MergeData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var filename = itor.FileName();
                var data = Loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(filename.File, message);
                }
            }

            return result;
        }

        public void Clear() {
            foreach (var itor in Readers) {
                itor.Clear();
            }
        }
    }

    public class FileName {
        public FileName(string name, string ext) {
            this.name = name;
            this.ext = ext;
        }

        public string Name {
            get {
                return name;
            }
        }

        public string Ext {
            get {
                return ext;
            }
        }

        public string File {
            get {
                return name + "." + ext;
            }
        }

        private readonly string name;
        private readonly string ext;
    }

    public interface Loader {
        public void Error(string name, string message);
        public byte[] Load(FileName filename);
    }

    public interface Reader {
        public FileName FileName();
        public string FromData(byte[] data);
        public string MergeData(byte[] data);
        public void Clear();
    }
}

// 以下是為了通過編譯的程式碼, 不可使用
