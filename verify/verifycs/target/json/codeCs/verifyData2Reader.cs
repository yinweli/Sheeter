// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class VerifyData2Reader {
        public static string FileName() {
            return "verifyData2.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<VerifyData2Storer>(data);
            return Datas != null;
        }

        public IDictionary<long, VerifyData2> Data {
            get {
                return Datas.Datas;
            }
        }

        private VerifyData2Storer Datas = null;
    }
}
