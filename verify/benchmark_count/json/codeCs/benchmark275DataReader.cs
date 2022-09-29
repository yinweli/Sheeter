// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark275DataReader {
        public static string FileName() {
            return "benchmark275Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark275DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark275Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark275DataStorer Datas = null;
    }
}
