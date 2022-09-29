// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark129DataReader {
        public static string FileName() {
            return "benchmark129Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark129DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark129Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark129DataStorer Datas = null;
    }
}
