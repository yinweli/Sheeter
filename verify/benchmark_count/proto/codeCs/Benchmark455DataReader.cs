// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class Benchmark455DataReader {
        public static string FileName() {
            return "benchmark455Data.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = Benchmark455DataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark455Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark455DataStorer Datas = null;
    }
}
