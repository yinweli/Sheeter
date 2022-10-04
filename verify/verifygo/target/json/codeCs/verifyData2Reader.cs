// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.IO;
using System.Collections.Generic;

namespace sheeterJson {

    public partial class VerifyData2Reader {

        public static string FileName()
        {
            return "verifyData2.json";
        }

        public bool FromPath(string path)
        {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data)
        {
            Datas = JsonConvert.DeserializeObject<VerifyData2Storer>(data);
            return Datas != null;
        }

        public long[] MergePath(params string[] path)
        {
            var repeats = new List<long>();

            foreach (var itor in path)
            {
                try
                {
                    repeats.AddRange(MergeData(File.ReadAllText(Path.Combine(itor, FileName()))));
                }
                catch
                {
                    // do nothing
                }
            }

            return repeats.ToArray();
        }

        public long[] MergeData(string data)
        {
            var repeats = new List<long>();
            var tmpl = JsonConvert.DeserializeObject<VerifyData2Storer>(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new VerifyData2Storer();

            foreach (var itor in tmpl.Datas)
            {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<long, VerifyData2> Data
        {
            get
            {
                return Datas.Datas;
            }
        }

        private VerifyData2Storer Datas = null;
    }
}
