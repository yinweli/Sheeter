// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class Benchmark05DataReader {
        public static string FileName() {
            return "benchmark05Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark05DataStorer>(data);
            return Datas != null;
        }

        public Dictionary<long, Benchmark05Data> Data {
            get {
                return Datas.Datas;
            }
        }

        private Benchmark05DataStorer Datas = null;
    }
}
