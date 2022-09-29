// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark112DataReader {
        public static string FileName() {
            return "benchmark112Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark112DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark112Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark112DataStorer Datas = null;
    }
}