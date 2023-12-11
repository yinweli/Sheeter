package tmpls

// ReaderCs json讀取器cs語言模板
var ReaderCs = Header + `
using Newtonsoft.Json;

namespace {{$.Namespace | $.FirstUpper}}
{
    using Data_ = {{$.StructName}};
    using Key_ = {{$.PkeyCs}};
    using Store_ = Dictionary<{{$.PkeyCs}}, {{$.StructName}}>;

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
        public string FromData(string data, bool clear)
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
                Data = new Store_();

            foreach (var itor in tmpl)
            {
                if (Data.ContainsKey(itor.Key))
                    return "from data: key duplicate";

                Data[itor.Key] = itor.Value;
            } // for

            return string.Empty;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            Data.Clear();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public bool TryGetValue(Key_ key, out Data_ value)
        {
            return Data.TryGetValue(key, out value);
        }

        /// <summary>
        /// 檢查索引是否存在
        /// </summary>
        public bool ContainsKey(Key_ key)
        {
            return Data.ContainsKey(key);
        }

        /// <summary>
        /// 取得迭代器
        /// </summary>
        public IEnumerator<KeyValuePair<Key_, Data_>> GetEnumerator()
        {
            return Data.GetEnumerator();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public Data_ this[Key_ key]
        {
            get { return Data[key]; }
        }

        /// <summary>
        /// 取得索引列表
        /// </summary>
        public ICollection<Key_> Keys
        {
            get { return Data.Keys; }
        }

        /// <summary>
        /// 取得資料列表
        /// </summary>
        public ICollection<Data_> Values
        {
            get { return Data.Values; }
        }

        /// <summary>
        /// 取得資料數量
        /// </summary>
        public int Count
        {
            get { return Data.Count; }
        }

        private Store_ Data = new Store_();
    }
}
`

// SheeterCs json表格器go語言模板
var SheeterCs = Header + `
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
            this.reader = new List<Reader>() {
{{- range $.Struct}}
                this.{{.StructName}},
{{- end}}
            };
        }

        /// <summary>
        /// 讀取資料處理
        /// </summary>
        public bool FromData()
        {
            if (loader == null)
                return false;

            var result = true;

            foreach (var itor in reader)
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
            foreach (var itor in reader)
                itor.Clear();
        }

        /// <summary>
        /// 裝載器物件
        /// </summary>
        private readonly Loader loader;

        /// <summary>
        /// 讀取器列表
        /// </summary>
        private readonly List<Reader> reader;

{{- range $.Struct}}
        /// <summary>
        /// {{.StructNote}}
        /// </summary>
        public readonly {{.ReaderName}} {{.StructName}} = new {{.ReaderName}}();
{{- end}}
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
`
