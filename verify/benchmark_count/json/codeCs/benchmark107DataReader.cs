// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark107DataReader {
        public static string FileName() {
            return "benchmark107Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark107DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark107Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark107DataStorer Datas = null;
    }
}
