// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark493DataReader {
        public static string FileName() {
            return "benchmark493Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark493DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark493Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark493DataStorer Datas = null;
    }
}
