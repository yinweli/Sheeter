// 以下是模板驗證用程式碼
// using區段可能與實際給的不一致, 要注意

using Google.Protobuf; // 這為了通過編譯的程式碼, 不可使用
using Google.Protobuf.Reflection; // 這為了通過編譯的程式碼, 不可使用
using System.Collections.Generic;
using pb = global::Google.Protobuf; // 這為了通過編譯的程式碼, 不可使用

namespace SheeterProto {
    using Data_ = Reward;
    using PKey_ = System.Int64;
    using Storer_ = RewardStorer;

    public partial class RewardReader : Reader {
        public string DataName() {
            return "reward";
        }

        public string DataExt() {
            return "bytes";
        }

        public string DataFile() {
            return "reward.bytes";
        }

        public string FromData(byte[] data) {
            Storer_ result;

            try {
                result = Storer_.Parser.ParseFrom(data);
            } catch {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(byte[] data) {
            Storer_ result;

            try {
                result = Storer_.Parser.ParseFrom(data);
            } catch {
                return "merge data failed: deserialize failed";
            }

            if (result == null)
                return "merge data failed: result null";

            foreach (var itor in result.Datas) {
                if (storer.Datas.ContainsKey(itor.Key))
                    return "merge data failed: key repeat";

                storer.Datas[itor.Key] = itor.Value;
            }

            return string.Empty;
        }

        public bool TryGetValue(PKey_ key, out Data_ value) {
            return storer.Datas.TryGetValue(key, out value);
        }

        public bool ContainsKey(PKey_ key) {
            return storer.Datas.ContainsKey(key);
        }

        public IEnumerator<KeyValuePair<PKey_, Data_>> GetEnumerator() {
            return storer.Datas.GetEnumerator();
        }

        public Data_ this[PKey_ key] {
            get {
                return storer.Datas[key];
            }
        }

        public ICollection<PKey_> Keys {
            get {
                return storer.Datas.Keys;
            }
        }

        public ICollection<Data_> Values {
            get {
                return storer.Datas.Values;
            }
        }

        public int Count {
            get {
                return storer.Datas.Count;
            }
        }

        private Storer_ storer = new Storer_();
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