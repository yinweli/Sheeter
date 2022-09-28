// generated by sheeter, DO NOT EDIT.

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class ExampleDataReader {
        public static string FileName() {
            return "exampleData.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = ExampleDataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public IDictionary<long, ExampleData> Data {
            get {
                return Datas.Datas;
            }
        }

        private ExampleDataStorer Datas = null;
    }
}
