// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

namespace Sheeter
{
    /// <summary>
    /// 表格資料
    /// </summary>
    public partial class Sheeter
    {
        public Sheeter(Loader loader)
        {
            this.loader = loader;
        }

        /// <summary>
        /// 讀取資料處理
        /// </summary>
        public bool FromData()
        {
            if (loader == null)
                return false;

            var result = true;

            foreach (var itor in new Reader[] { this.ExampleData })
            {
                var filename = itor.FileName();
                var data = loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var error = itor.FromData(data, true);

                if (error.Length != 0)
                {
                    result = false;
                    loader.Error(filename.File, error);
                } // if
            } // for

            return result;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            this.ExampleData.Clear();
        }

        /// <summary>
        /// 裝載器物件
        /// </summary>
        private readonly Loader loader;

        /// <summary>
        /// example.xlsx#Data
        /// </summary>
        public readonly ExampleDataReader ExampleData = new ExampleDataReader();
    }

    /// <summary>
    /// 裝載器介面
    /// </summary>
    public interface Loader
    {
        /// <summary>
        /// 讀取檔案
        /// </summary>
        public string Load(FileName filename);

        /// <summary>
        /// 錯誤處理
        /// </summary>
        public void Error(string name, string message);
    }

    /// <summary>
    /// 讀取器介面
    /// </summary>
    public interface Reader
    {
        /// <summary>
        /// 取得檔名物件
        /// </summary>
        public FileName FileName();

        /// <summary>
        /// 讀取資料
        /// </summary>
        public string FromData(string data, bool clear);

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear();
    }

    /// <summary>
    /// 檔名資料
    /// </summary>
    public class FileName
    {
        public FileName(string name, string ext)
        {
            this.name = name;
            this.ext = ext;
        }

        /// <summary>
        /// 取得名稱
        /// </summary>
        public string Name
        {
            get { return name; }
        }

        /// <summary>
        /// 取得副檔名
        /// </summary>
        public string Ext
        {
            get { return ext; }
        }

        /// <summary>
        /// 取得完整檔名
        /// </summary>
        public string File
        {
            get { return name + ext; }
        }

        /// <summary>
        /// 名稱
        /// </summary>
        private readonly string name;

        /// <summary>
        /// 副檔名
        /// </summary>
        private readonly string ext;
    }
}
