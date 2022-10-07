// 以下是模板驗證用程式碼
// using區段可能與實際給的不一致, 要注意

using Google.Protobuf; // 這為了通過編譯的程式碼, 不可使用
using Google.Protobuf.Reflection; // 這為了通過編譯的程式碼, 不可使用
using System.Collections.Generic;
using pb = global::Google.Protobuf; // 這為了通過編譯的程式碼, 不可使用

namespace SheeterProto {
    public partial class RewardReader {
        public string DataName() {
            return "reward";
        }

        public string DataExt() {
            return "bytes";
        }

        public string DataFile() {
            return "reward.bytes";
        }

        public bool FromData(byte[] data) {
            Datas = RewardStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public long[] MergeData(byte[] data) {
            var repeats = new List<long>();
            var tmpl = RewardStorer.Parser.ParseFrom(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new RewardStorer();

            foreach (var itor in tmpl.Datas) {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<long, Reward> Data {
            get {
                return Datas.Datas;
            }
        }

        private RewardStorer Datas = null;
    }
}

// 以下是為了通過編譯的程式碼, 不可使用

namespace SheeterProto {
    public sealed partial class Reward : pb::IMessage<Reward> {
        public MessageDescriptor Descriptor => throw new System.NotImplementedException();

        public int CalculateSize() {
            throw new System.NotImplementedException();
        }

        public Reward Clone() {
            throw new System.NotImplementedException();
        }

        public bool Equals(Reward? other) {
            throw new System.NotImplementedException();
        }

        public void MergeFrom(Reward message) {
            throw new System.NotImplementedException();
        }

        public void MergeFrom(CodedInputStream input) {
            throw new System.NotImplementedException();
        }

        public void WriteTo(CodedOutputStream output) {
            throw new System.NotImplementedException();
        }
    }

    public sealed partial class RewardStorer : pb::IMessage<RewardStorer> {
        private static readonly pb::MessageParser<RewardStorer> _parser = new pb::MessageParser<RewardStorer>(() => new RewardStorer());
        public static pb::MessageParser<RewardStorer> Parser { get { return _parser; } }
        public Dictionary<long, Reward> Datas = new Dictionary<long, Reward>();
        public MessageDescriptor Descriptor => throw new System.NotImplementedException();

        public int CalculateSize() {
            throw new System.NotImplementedException();
        }

        public RewardStorer Clone() {
            throw new System.NotImplementedException();
        }

        public bool Equals(RewardStorer? other) {
            throw new System.NotImplementedException();
        }

        public void MergeFrom(RewardStorer message) {
            throw new System.NotImplementedException();
        }

        public void MergeFrom(CodedInputStream input) {
            throw new System.NotImplementedException();
        }

        public void WriteTo(CodedOutputStream output) {
            throw new System.NotImplementedException();
        }
    }
}