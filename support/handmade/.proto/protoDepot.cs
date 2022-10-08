﻿// 以下是模板驗證用程式碼

using System.Collections.Generic;

namespace SheeterProto {
    public partial class Depot {
        public readonly RewardReader Reward = new RewardReader();
        private readonly List<ReaderInterface> Readers = new List<ReaderInterface>();

        public Depot() {
            Readers.Add(Reward);
        }

        public bool FromData(DelegateLoad load, DelegateError error) {
            var result = true;

            foreach (var itor in Readers) {
                var data = load(itor.DataName(), itor.DataExt());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData(DelegateLoad load, DelegateError error) {
            var result = true;

            foreach (var itor in Readers) {
                var data = load(itor.DataName(), itor.DataExt());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    error(itor.DataName(), message);
                }
            }

            return result;
        }

        public delegate void DelegateError(string name, string message);
        public delegate byte[] DelegateLoad(string name, string ext);
    }

    public interface ReaderInterface {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(byte[] data);
        public string MergeData(byte[] data);
    }
}

// 以下是為了通過編譯的程式碼, 不可使用