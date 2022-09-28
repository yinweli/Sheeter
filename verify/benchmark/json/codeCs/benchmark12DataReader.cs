// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeter {
    public partial class Benchmark12DataReader {
        public static string FileName() {
            return "benchmark12Data.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark12DataStorer>(data);
            return Datas != null;
        }

        public Benchmark12DataStorer Datas = null;
    }
}
