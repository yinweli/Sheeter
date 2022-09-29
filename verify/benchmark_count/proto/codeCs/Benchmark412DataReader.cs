// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class Benchmark412DataReader {
        public static string FileName() {
            return "benchmark412Data.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = Benchmark412DataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark412Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark412DataStorer Datas = null;
    }
}
