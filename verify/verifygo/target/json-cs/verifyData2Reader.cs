// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeter {
    public partial class VerifyData2Reader {
        public static readonly string Json = "json/verifyData2.json";

        public static Dictionary<long, VerifyData2> FromJsonFile(string path) {
            return FromJsonString(File.ReadAllText(path));
        }

        public static Dictionary<long, VerifyData2> FromJsonString(string data) {
            var temps = JsonConvert.DeserializeObject<Dictionary<string, VerifyData2>>(data);

            if (temps == null) {
                return null;
            }

            var datas = new Dictionary<long, VerifyData2>();

            foreach(var itor in temps) {
                datas[Convert.ToInt64(itor.Key)] = itor.Value;
            }

            return datas;
        }
    }
}
