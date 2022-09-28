// generated by sheeter, DO NOT EDIT.

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class Benchmark14DataReader {
        public static string FileName() {
            return "benchmark14Data.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = Benchmark14DataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark14Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark14DataStorer Datas = null;
    }
}
