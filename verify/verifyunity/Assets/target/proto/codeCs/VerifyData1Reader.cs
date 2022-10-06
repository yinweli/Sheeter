// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterProto {
    public partial class VerifyData1Reader {
        public string DataName() {
            return "verifyData1";
        }

        public string DataExt() {
            return "bytes";
        }

        public string DataFile() {
            return "verifyData1.bytes";
        }

        public bool FromData(byte[] data) {
            Datas = VerifyData1Storer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public long[] MergeData(byte[] data) {
            var repeats = new List<long>();
            var tmpl = VerifyData1Storer.Parser.ParseFrom(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new VerifyData1Storer();

            foreach (var itor in tmpl.Datas) {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<long, VerifyData1> Data {
            get {
                return Datas.Datas;
            }
        }

        private VerifyData1Storer Datas = null;
    }
}
