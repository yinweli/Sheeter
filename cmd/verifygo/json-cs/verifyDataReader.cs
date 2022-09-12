// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;

namespace verifydata
{
    public partial class Reader
    {
        public static readonly string Json = "json/verifyData.json";

        public static Dictionary<long, Struct> FromJsonFile(string path)
        {
            return FromJsonString(File.ReadAllText(path));
        }

        public static Dictionary<long, Struct> FromJsonString(string data)
        {
            var temps = JsonConvert.DeserializeObject<Dictionary<string, Struct>>(data);

            if (temps == null)
            {
                return null;
            }

            var datas = new Dictionary<long, Struct>();

            foreach (var itor in temps)
            {
                datas[Convert.ToInt64(itor.Key)] = itor.Value;
            }

            return datas;
        }
    }
}
