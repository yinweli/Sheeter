// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class Benchmark417DataReader {
        public static string FileName() {
            return "benchmark417Data.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = Benchmark417DataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark417Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark417DataStorer Datas = null;
    }
}
