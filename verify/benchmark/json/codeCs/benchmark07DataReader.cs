// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark07DataReader {
        public static string FileName() {
            return "benchmark07Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark07DataStorer>(data);
            return Datas != null;
        }

        public IDictionary<long, Benchmark07Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark07DataStorer Datas = null;
    }
}
