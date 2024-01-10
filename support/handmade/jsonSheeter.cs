// 以下是驗證程式碼, 不可使用
// using區段可能與實際使用的不一致, 要注意

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
            var index = 0;

            foreach (var itor in new Reader[] { this.Alone0, this.Alone1, })
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

            index = 0;

            foreach (var itor in new Reader[] { this.Alone0, this.Alone1, })
            {
                var filename = itor.FileName();
                var data = loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var error = Merge0.FromData(data, index == 0);

                if (error.Length != 0)
                {
                    result = false;
                    loader.Error(filename.File, error);
                } // if

                index++;
            } // for

            index = 0;

            foreach (var itor in new Reader[] { this.Alone0, this.Alone1, })
            {
                var filename = itor.FileName();
                var data = loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var error = Merge1.FromData(data, index == 0);

                if (error.Length != 0)
                {
                    result = false;
                    loader.Error(filename.File, error);
                } // if

                index++;
            } // for

            return result;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            Alone0.Clear();
            Alone1.Clear();
            Merge0.Clear();
            Merge1.Clear();
        }

        /// <summary>
        /// 裝載器物件
        /// </summary>
        private readonly Loader loader;

        /// <summary>
        /// 獨立表格說明
        /// </summary>
        public readonly HandmadeReader Alone0 = new HandmadeReader();

        /// <summary>
        /// 獨立表格說明
        /// </summary>
        public readonly HandmadeReader Alone1 = new HandmadeReader();

        /// <summary>
        /// 合併表格說明
        /// </summary>
        public readonly HandmadeReader Merge0 = new HandmadeReader();

        /// <summary>
        /// 合併表格說明
        /// </summary>
        public readonly HandmadeReader Merge1 = new HandmadeReader();
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
