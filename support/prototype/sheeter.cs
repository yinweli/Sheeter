using System.Collections.Generic;
using System.Threading.Tasks;

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
        public async Task<bool> FromData()
        {
            progress.Reset();

            if (loader == null)
                return false;

            var task = new List<Task<bool>>();

            foreach (var itor in new Reader[] { this.Alone0, this.Alone1 })
            {
                var tmpl = itor;

                task.Add(
                    Task.Run(() =>
                    {
                        var filename = tmpl.FileName();
                        var data = loader.Load(filename);

                        if (data == null || data.Length == 0)
                            return true;

                        var error = tmpl.FromData(data, true, progress);

                        if (error.Length == 0)
                            return true;

                        loader.Error(filename.File, error);
                        return false;
                    })
                );
            } // for

            task.Add(
                Task.Run(() =>
                {
                    var first = true;
                    var result = true;

                    foreach (var itor in new Reader[] { this.Alone0, this.Alone1 })
                    {
                        var filename = itor.FileName();
                        var data = loader.Load(filename);

                        if (data == null || data.Length == 0)
                            continue;

                        var error = Merge0.FromData(data, first, progress);

                        if (error.Length != 0)
                        {
                            loader.Error(filename.File, error);
                            result = false;
                        } // if

                        first = false;
                    } // for

                    return result;
                })
            );
            task.Add(
                Task.Run(() =>
                {
                    var first = true;
                    var result = true;

                    foreach (var itor in new Reader[] { this.Alone0, this.Alone1 })
                    {
                        var filename = itor.FileName();
                        var data = loader.Load(filename);

                        if (data == null || data.Length == 0)
                            continue;

                        var error = Merge1.FromData(data, first, progress);

                        if (error.Length != 0)
                        {
                            loader.Error(filename.File, error);
                            result = false;
                        } // if

                        first = false;
                    } // for

                    return result;
                })
            );

            var result = await Task.WhenAll(task);

            progress.Complete();
            return result.All(itor => itor);
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            progress.Reset();
            Alone0.Clear();
            Alone1.Clear();
            Merge0.Clear();
            Merge1.Clear();
        }

        /// <summary>
        /// 取得進度值
        /// </summary>
        public float Progress()
        {
            return progress.Get();
        }

        /// <summary>
        /// 裝載器物件
        /// </summary>
        private readonly Loader loader;

        /// <summary>
        /// 進度物件
        /// </summary>
        private readonly Progress progress = new();

        /// <summary>
        /// 獨立表格說明
        /// </summary>
        public readonly HandmadeReader Alone0 = new();

        /// <summary>
        /// 獨立表格說明
        /// </summary>
        public readonly HandmadeReader Alone1 = new();

        /// <summary>
        /// 合併表格說明
        /// </summary>
        public readonly HandmadeReader Merge0 = new();

        /// <summary>
        /// 合併表格說明
        /// </summary>
        public readonly HandmadeReader Merge1 = new();
    }
}
