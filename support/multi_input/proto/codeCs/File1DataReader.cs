// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterProto
{
    using Data_ = File1Data;
    using PKey_ = System.Int64;
    using Storer_ = File1DataStorer;

    public partial class File1DataReader : Reader
    {
        public FileName FileName()
        {
            return new FileName("file1Data", ".bytes");
        }

        public string FromData(byte[] data)
        {
            Storer_ result;

            try
            {
                result = Storer_.Parser.ParseFrom(data);
            }
            catch
            {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(byte[] data)
        {
            Storer_ result;

            try
            {
                result = Storer_.Parser.ParseFrom(data);
            }
            catch
            {
                return "merge data failed: deserialize failed";
            }

            if (result == null)
                return "merge data failed: result null";

            foreach (var itor in result.Datas)
            {
                if (storer.Datas.ContainsKey(itor.Key))
                    return "merge data failed: key repeat";

                storer.Datas[itor.Key] = itor.Value;
            }

            return string.Empty;
        }

        public void Clear()
        {
            storer.Datas.Clear();
        }

        public bool TryGetValue(PKey_ key, out Data_ value)
        {
            return storer.Datas.TryGetValue(key, out value);
        }

        public bool ContainsKey(PKey_ key)
        {
            return storer.Datas.ContainsKey(key);
        }

        public IEnumerator<KeyValuePair<PKey_, Data_>> GetEnumerator()
        {
            return storer.Datas.GetEnumerator();
        }

        public Data_ this[PKey_ key]
        {
            get { return storer.Datas[key]; }
        }

        public ICollection<PKey_> Keys
        {
            get { return storer.Datas.Keys; }
        }

        public ICollection<Data_> Values
        {
            get { return storer.Datas.Values; }
        }

        public int Count
        {
            get { return storer.Datas.Count; }
        }

        private Storer_ storer = new Storer_();
    }
}
