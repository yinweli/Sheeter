// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeter {
    using VerifyData2Storer = Dictionary<long, VerifyData2>;

    public partial class VerifyData2Reader {
        public static string FileName() {
            return "verifyData2.json";
        }

        public bool FromFullPath(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromHalfPath(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<VerifyData2Storer>(data);
            return Datas != null;
        }

        public VerifyData2Storer Datas = null;
    }
}
