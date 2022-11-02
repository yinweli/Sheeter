// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterJson
{
    public partial class Depot
    {
        public Loader Loader { get; set; }
        public readonly File1DataReader File1Data = new File1DataReader();
        public readonly File2DataReader File2Data = new File2DataReader();
        public readonly Path1DataReader Path1Data = new Path1DataReader();
        public readonly Path2DataReader Path2Data = new Path2DataReader();
        public readonly Path3DataReader Path3Data = new Path3DataReader();
        private readonly List<Reader> Readers = new List<Reader>();

        public Depot()
        {
            Readers.Add(File1Data);
            Readers.Add(File2Data);
            Readers.Add(Path1Data);
            Readers.Add(Path2Data);
            Readers.Add(Path3Data);
        }

        public bool FromData()
        {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers)
            {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0)
                {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData()
        {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers)
            {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0)
                {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }

        public void Clear()
        {
            foreach (var itor in Readers)
            {
                itor.Clear();
            }
        }
    }

    public interface Loader
    {
        public void Error(string name, string message);
        public string Load(string name, string ext, string fullname);
    }

    public interface Reader
    {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(string data);
        public string MergeData(string data);
        public void Clear();
    }
}