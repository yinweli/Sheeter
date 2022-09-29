// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {
    public partial class VerifyData1Reader {
        public static string FileName() {
            return "verifyData1.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<VerifyData1Storer>(data);
            return Datas != null;
        }

        public IDictionary<long, VerifyData1> Data {
            get {
                return Datas.Datas;
            }
        }

        private VerifyData1Storer Datas = null;
    }
}
