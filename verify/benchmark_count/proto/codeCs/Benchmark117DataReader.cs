// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class Benchmark117DataReader {
        public static string FileName() {
            return "benchmark117Data.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = Benchmark117DataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark117Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark117DataStorer Datas = null;
    }
}
