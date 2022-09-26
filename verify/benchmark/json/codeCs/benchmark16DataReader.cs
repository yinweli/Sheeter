// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeter {
    using Benchmark16DataStorer = Dictionary<long, Benchmark16Data>;

    public partial class Benchmark16DataReader {
        public static string FileName() {
            return "benchmark16Data.json";
        }

        public bool FromFullPath(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromHalfPath(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<Benchmark16DataStorer>(data);
            return Datas != null;
        }

        public Benchmark16DataStorer Datas = null;
    }
}
