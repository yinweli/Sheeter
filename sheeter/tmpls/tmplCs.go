package tmpls

// ReaderCs cs語言讀取模板
var ReaderCs = Header + `
using System;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace {{$.Namespace | $.FirstUpper}}
{
    using Data_ = {{$.StructName}};
    using Key_ = {{$.PrimaryCs}};
    using Store_ = Dictionary<{{$.PrimaryCs}}, {{$.StructName}}>;

    /// <summary>
    /// {{$.StructNote}}
    /// </summary>
    public partial class {{$.StructName}}
    {
{{- range $.Field}}
        /// <summary>
        /// {{.FieldNote}}
        /// </summary>
        [JsonProperty("{{.FieldName}}")]
        public {{.FieldTypeCs}} {{.FieldName}} { get; set; }
{{- end}}
    }

    /// <summary>
    /// {{$.StructNote}}
    /// </summary>
    public partial class {{$.ReaderName}} : Reader
    {
        /// <summary>
        /// 取得檔名物件
        /// </summary>
        public FileName FileName()
        {
            return new FileName("{{$.JsonName}}", "{{$.JsonExt}}");
        }

        /// <summary>
        /// 讀取資料
        /// </summary>
        public string FromData(string data, bool clear, Progress progress)
        {
            Store_ tmpl;

            try
            {
                tmpl = JsonConvert.DeserializeObject<Store_>(data);
            } // try
            catch
            {
                return "from data: deserialize failed";
            } // catch

            if (tmpl == null)
                return "from data: deserialize failed";

            if (clear)
                this.data = new();

            var task = progress.Reg();
            var curr = 0;
            var total = tmpl.Count;

            foreach (var itor in tmpl)
            {
                if (this.data.ContainsKey(itor.Key))
                    return "from data: key duplicate [{{$.JsonName}} : " + itor.Key + "]";

                this.data[itor.Key] = itor.Value;
                curr++;
                progress.Set(task, curr, total);
            } // for

            return string.Empty;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            data.Clear();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public bool TryGetValue(Key_ key, out Data_ value)
        {
            return data.TryGetValue(key, out value);
        }

        /// <summary>
        /// 檢查索引是否存在
        /// </summary>
        public bool ContainsKey(Key_ key)
        {
            return data.ContainsKey(key);
        }

        /// <summary>
        /// 取得迭代器
        /// </summary>
        public IEnumerator<KeyValuePair<Key_, Data_>> GetEnumerator()
        {
            return data.GetEnumerator();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public Data_ this[Key_ key]
        {
            get { return data[key]; }
            set { data[key] = value; }
        }

        /// <summary>
        /// 取得索引列表
        /// </summary>
        public ICollection<Key_> Keys
        {
            get { return data.Keys; }
        }

        /// <summary>
        /// 取得資料列表
        /// </summary>
        public ICollection<Data_> Values
        {
            get { return data.Values; }
        }

        /// <summary>
        /// 取得資料數量
        /// </summary>
        public int Count
        {
            get { return data.Count; }
        }

        private Store_ data = new();
    }
}
`

// SheeterCs cs語言表格模板
var SheeterCs = Header + `
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace {{$.Namespace | $.FirstUpper}}
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

            foreach (var itor in new Reader[] {
{{- range $.Alone}}
                this.{{.StructName}},
{{- end}}
            })
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

{{- range $.Merge}}
            task.Add(
                Task.Run(() =>
                {
                    var first = true;
                    var result = true;

                    foreach (var itor in new Reader[]
                    {
{{- range $name := .MemberName}}
		                this.{{$name}},
{{- end}}
	                }) {
                        var filename = itor.FileName();
                        var data = loader.Load(filename);

                        if (data == null || data.Length == 0)
                            continue;

                        var error = {{.StructName}}.FromData(data, first, progress);

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
{{- end}}

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
{{- range $.Alone}}
            this.{{.StructName}}.Clear();
{{- end}}
{{- range $.Merge}}
            this.{{.StructName}}.Clear();
{{- end}}
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

{{- range $.Alone}}
        /// <summary>
        /// {{.StructNote}}
        /// </summary>
        public readonly {{.ReaderName}} {{.StructName}} = new();
{{- end}}
{{- range $.Merge}}
        /// <summary>
        /// {{.StructNote}}
        /// </summary>
        public readonly {{.ReaderName}} {{.StructName}} = new();
{{- end}}
    }
}
`

// HelperCs cs語言工具模板
var HelperCs = Header + `
#nullable enable

using System;
using System.Collections.Generic;
using System.Globalization;
using System.Numerics;
using System.Text.RegularExpressions;
using System.Threading;

namespace {{$.Namespace | $.FirstUpper}}
{
    /// <summary>
    /// 裝載器介面
    /// </summary>
    public interface Loader
    {
        /// <summary>
        /// 讀取檔案, 實作時須注意必須維持執行緒安全
        /// </summary>
        public string Load(FileName filename);

        /// <summary>
        /// 錯誤處理, 實作時須注意必須維持執行緒安全
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
        public string FromData(string data, bool clear, Progress progress);

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear();
    }

    /// <summary>
    /// 字串解析組件
    /// </summary>
    public static class Parse
    {
        /// <summary>
        /// 解析函式類型
        /// </summary>
        public delegate (object? obj, bool ok) Parser(string value);

        /// <summary>
        /// 字串解析結果
        /// </summary>
        public class Result
        {
            public Result(object? value, string error = "")
            {
                this.value = value;
                this.error = error ?? string.Empty;
            }

            /// <summary>
            /// 取得結果物件
            /// </summary>
            public object? Value()
            {
                return value;
            }

            /// <summary>
            /// 取得結果物件, 會丟出例外
            /// </summary>
            public T ValueAs<T>()
            {
                if (string.IsNullOrEmpty(error) == false)
                    throw new InvalidOperationException(error);

                if (value is not T result)
                {
                    var actual = value is null ? "null" : value.GetType().ToString();
                    var expect = typeof(T).ToString();
                    throw new InvalidCastException($"value is {actual} not {expect}");
                } // if

                return result;
            }

            /// <summary>
            /// 取得結果物件, 不會丟出例外
            /// </summary>
            public bool TryValue<T>(out T value)
            {
                if (string.IsNullOrEmpty(error) == false || this.value is not T result)
                {
                    value = default!;
                    return false;
                } // if

                value = result;
                return true;
            }

            /// <summary>
            /// 取得錯誤訊息
            /// </summary>
            public string Error()
            {
                return error;
            }

            public static implicit operator bool(Result obj)
            {
                return obj is not null && string.IsNullOrEmpty(obj.error);
            }

            /// <summary>
            /// 結果物件
            /// </summary>
            private readonly object? value;

            /// <summary>
            /// 錯誤訊息
            /// </summary>
            private readonly string error = string.Empty;
        }

        static Parse()
        {
            AddParse<int>(value =>
            {
                value = (value ?? string.Empty).Trim();

                if (int.TryParse(value, out int result))
                    return (result, true);

                return (null, false);
            });
            AddParse<long>(value =>
            {
                value = (value ?? string.Empty).Trim();

                if (long.TryParse(value, out long result))
                    return (result, true);

                return (null, false);
            });
            AddParse<float>(value =>
            {
                value = (value ?? string.Empty).Trim();

                if (float.TryParse(value, out float result))
                    return (result, true);

                return (null, false);
            });
            AddParse<double>(value =>
            {
                value = (value ?? string.Empty).Trim();

                if (double.TryParse(value, out double result))
                    return (result, true);

                return (null, false);
            });
            AddParse<string>(value =>
            {
                return (value, true);
            });
            AddParse<Ratio>(value =>
            {
                value = (value ?? string.Empty).Trim();

                if (value.Length == 0)
                    return (new Ratio("0"), true);

                if (double.TryParse(value, out _) == false)
                    return (null, false);

                return (new Ratio(value), true);
            });
            AddParse<Duration>(value =>
            {
                value = (value ?? string.Empty).Trim();

                if (value.Length == 0)
                    return (new Duration("0s"), true);

                if (Duration.parse(value, out _) == false)
                    return (null, false);

                return (new Duration(value), true);
            });
        }

        /// <summary>
        /// 新增解析
        /// </summary>
        public static void AddParse<T>(Parser parser)
        {
            @lock.EnterWriteLock();

            try
            {
                parse[typeof(T)] = parser;
            } // try
            finally
            {
                @lock.ExitWriteLock();
            } // finally
        }

        /// <summary>
        /// 執行解析
        /// </summary>
        public static Result RunParse<T>(string value)
        {
            @lock.EnterReadLock();

            try
            {
                var type = typeof(T);

                if (parse.TryGetValue(type, out var parser) == false || parser == null)
                    return new Result(null, $"{type} not exist");

                var result = parser(value);

                if (result.ok == false)
                    return new Result(null, $"{type} parse failed");

                return new Result(result.obj);
            } // try
            finally
            {
                @lock.ExitReadLock();
            } // finally
        }

        /// <summary>
        /// 字串擴充: 執行解析
        /// </summary>
        public static Result Parses<T>(this string value)
        {
            return RunParse<T>(value);
        }

        /// <summary>
        /// 字串擴充: 執行解析, 會丟出例外
        /// </summary>
        public static T ParseAs<T>(this string value)
        {
            return RunParse<T>(value).ValueAs<T>();
        }

        /// <summary>
        /// 字串擴充: 執行解析, 不會丟出例外
        /// </summary>
        public static bool TryParse<T>(this string value, out T result)
        {
            return RunParse<T>(value).TryValue(out result);
        }

        /// <summary>
        /// 解析列表
        /// </summary>
        private static readonly Dictionary<Type, Parser> parse = new();

        /// <summary>
        /// 解析執行緒鎖
        /// </summary>
        private static readonly ReaderWriterLockSlim @lock = new();
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

    /// <summary>
    /// 進度資料
    /// </summary>
    public class Progress
    {
        /// <summary>
        /// 重置進度
        /// </summary>
        public void Reset()
        {
            @lock.EnterWriteLock();

            try
            {
                done = false;
                task = 0;
                data.Clear();
            } // try
            finally
            {
                @lock.ExitWriteLock();
            } // finally
        }

        /// <summary>
        /// 完成進度
        /// </summary>
        public void Complete()
        {
            @lock.EnterWriteLock();

            try
            {
                done = true;
            } // try
            finally
            {
                @lock.ExitWriteLock();
            } // finally
        }

        /// <summary>
        /// 註冊進度
        /// </summary>
        public int Reg()
        {
            @lock.EnterWriteLock();

            try
            {
                task++;
                data[task] = 0.0f;
                return task;
            } // try
            finally
            {
                @lock.ExitWriteLock();
            } // finally
        }

        /// <summary>
        /// 設定進度
        /// </summary>
        public void Set(int task, int curr, int total)
        {
            @lock.EnterWriteLock();

            try
            {
                var value = 0.0f;

                if (curr <= 0 || total <= 0)
                    value = 0.0f;
                else if (curr >= total)
                    value = 1.0f;
                else
                    value = (float)curr / total;

                data[task] = value;
            } // try
            finally
            {
                @lock.ExitWriteLock();
            } // finally
        }

        /// <summary>
        /// 取得進度, 進度值為 0 ~ 1 的浮點數
        /// </summary>
        public float Get()
        {
            @lock.EnterReadLock();

            try
            {
                if (done)
                    return 1.0f;

                if (data.Count == 0)
                    return 0.0f;

                var curr = 0.0f;

                foreach (var itor in data.Values)
                    curr += itor;

                return curr / data.Count;
            } // try
            finally
            {
                @lock.ExitReadLock();
            } // finally
        }

        /// <summary>
        /// 完成旗標
        /// </summary>
        private bool done = false;

        /// <summary>
        /// 進度編號
        /// </summary>
        private int task = 0;

        /// <summary>
        /// 進度列表
        /// </summary>
        private readonly Dictionary<int, float> data = new();

        /// <summary>
        /// 執行緒鎖
        /// </summary>
        private readonly ReaderWriterLockSlim @lock = new();
    }

    /// <summary>
    /// 比例類型
    /// </summary>
    public readonly struct Ratio
    {
        public Ratio(string value)
        {
            this.value = (value ?? string.Empty).Trim();
        }

        /// <summary>
        /// 取得浮點數
        /// </summary>
        public float Float32()
        {
            return float.TryParse(value, out var result) ? result : 0;
        }

        /// <summary>
        /// 取得浮點數
        /// </summary>
        public double Float64()
        {
            return double.TryParse(value, out var result) ? result : 0;
        }

        /// <summary>
        /// 取得字串
        /// </summary>
        public override string ToString()
        {
            return value;
        }

        /// <summary>
        /// 以基準值計算比例值
        /// </summary>
        public int RatioInt32(int @base)
        {
            return (int)ratio(@base);
        }

        /// <summary>
        /// 以基準值計算比例值
        /// </summary>
        public long RatioInt64(int @base)
        {
            return ratio(@base);
        }

        /// <summary>
        /// 以基準值計算比例值
        /// </summary>
        private long ratio(int @base)
        {
            if (@base <= 0)
                return 0;

            if (decimal.TryParse(value, out var result) == false)
                return 0;

            try
            {
                return (long)decimal.Truncate(result * @base);
            } // try
            catch (OverflowException)
            {
                return 0;
            } // catch
        }

        /// <summary>
        /// 比例字串
        /// </summary>
        private readonly string value;
    }

    /// <summary>
    /// 時長類型
    /// </summary>
    public readonly struct Duration
    {
        public Duration(string value)
        {
            this.value = (value ?? string.Empty).Trim();
        }

        /// <summary>
        /// 取得時長
        /// </summary>
        public TimeSpan Interval()
        {
            return parse(value, out var result) ? result : TimeSpan.Zero;
        }

        /// <summary>
        /// 取得字串
        /// </summary>
        public override string ToString()
        {
            return value;
        }

        /// <summary>
        /// 時長字串
        /// </summary>
        private readonly string value;

        /// <summary>
        /// 解析時長
        /// </summary>
        internal static bool parse(string input, out TimeSpan result)
        {
            result = TimeSpan.Zero;

            var s = (input ?? string.Empty).Trim();
            var sign = 1;

            if (s.StartsWith("-", StringComparison.Ordinal))
            {
                sign = -1;
                s = s.Substring(1).TrimStart();
            } // if

            while (s.Length > 0)
            {
                var match = regex.Match(s);

                if (match.Success == false)
                {
                    if (string.IsNullOrWhiteSpace(s)) // 如果剩下的都是空白, 表示解析完成
                        break;

                    return false;
                } // if

                var value = match.Groups[1].Value;

                if (long.TryParse(value, out var number) == false)
                    return false;

                var unit = match.Groups[2].Value.ToLowerInvariant();

                switch (unit)
                {
                    case "d":
                        result += TimeSpan.FromDays(number);
                        break;

                    case "h":
                        result += TimeSpan.FromHours(number);
                        break;

                    case "m":
                        result += TimeSpan.FromMinutes(number);
                        break;

                    case "s":
                        result += TimeSpan.FromSeconds(number);
                        break;

                    case "ms":
                        result += TimeSpan.FromMilliseconds(number);
                        break;

                    default:
                        return false;
                } // switch

                s = s[match.Length..];
            } // while

            if (sign < 0)
                result = -result;

            return true;
        }

        /// <summary>
        /// 時長正則表達式
        /// </summary>
        internal static readonly Regex regex = new Regex(
            @"^\s*(\d+)\s*(ms|s|m|h|d)",
            RegexOptions.IgnoreCase | RegexOptions.Compiled
        );
    }
}
`
