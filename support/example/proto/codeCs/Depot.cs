// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterProto {
    public partial class Depot {
        public Loader Loader { get; set; }
        public readonly ExampleDataReader ExampleData = new ExampleDataReader();
        private readonly List<Reader> Readers = new List<Reader>();

        public Depot() {
            Readers.Add(ExampleData);
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
        public byte[] Load(string name, string ext, string fullname);
    }

    public interface Reader {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(byte[] data);
        public string MergeData(byte[] data);
    }
}
