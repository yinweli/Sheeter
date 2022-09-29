// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark353DataReader {
        public static string FileName() {
            return "benchmark353Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark353DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark353Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark353DataStorer Datas = null;
    }
}