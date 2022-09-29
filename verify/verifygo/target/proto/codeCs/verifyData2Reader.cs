// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class VerifyData2Reader {
        public static string FileName() {
            return "verifyData2.pbd";
        }

        public bool FromPath(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = VerifyData2Storer.Parser.ParseFrom(data);
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
