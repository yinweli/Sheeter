package sheeter

import (
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Loader 裝載器介面
type Loader interface {
	// Load 讀取檔案, 實作時須注意必須維持執行緒安全
	Load(filename FileName) []byte

	// Error 錯誤處理, 實作時須注意必須維持執行緒安全
	Error(name string, err error)
}

// Reader 讀取器介面
type Reader interface {
	// FileName 取得檔名物件
	FileName() FileName

	// FromData 讀取資料
	FromData(data []byte, clear bool, progress *Progress) error

	// Clear 清除資料
	Clear()
}

// Parser 解析函式類型
type Parser[T any] func(value string) (result T, err error)

// AddParse 新增解析
func AddParse[T any](parser Parser[T]) {
	parseLock.Lock()
	defer parseLock.Unlock()

	parse[reflect.TypeOf((*T)(nil)).Elem()] = func(value string) (result any, err error) {
		return parser(value)
	}
}

// RunParse 執行解析
func RunParse[T any](value string) (result T, err error) {
	parseLock.RLock()
	parser := parse[reflect.TypeOf((*T)(nil)).Elem()]
	parseLock.RUnlock()

	if parser == nil {
		return result, fmt.Errorf("run parse: type not exist")
	} // if

	v, err := parser(value)

	if err != nil {
		return result, fmt.Errorf("run parse: %w", err)
	} // if

	result, ok := v.(T)

	if ok == false {
		return result, fmt.Errorf("run parse: type invalid")
	} // if

	return result, nil
}

var parse = map[reflect.Type]func(string) (any, error){} // 解析列表
var parseLock = sync.RWMutex{}                           // 解析執行緒鎖

// NewFileName 建立檔名資料
func NewFileName(name, ext string) FileName {
	return FileName{
		name: name,
		ext:  ext,
	}
}

// FileName 檔名資料
type FileName struct {
	name string // 名稱
	ext  string // 副檔名
}

// Name 取得名稱
func (this FileName) Name() string {
	return this.name
}

// Ext 取得副檔名
func (this FileName) Ext() string {
	return this.ext
}

// File 取得完整檔名
func (this FileName) File() string {
	return this.name + this.ext
}

func NewProgress() *Progress {
	return &Progress{
		data: map[int]float32{},
	}
}

// Progress 進度資料
type Progress struct {
	done bool            // 完成旗標
	task int             // 進度編號
	data map[int]float32 // 進度列表
	lock sync.RWMutex    // 執行緒鎖
}

// Reset 重置進度
func (this *Progress) Reset() {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.done = false
	this.task = 0
	this.data = map[int]float32{}
}

// Complete 完成進度
func (this *Progress) Complete() {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.done = true
}

// Reg 註冊進度
func (this *Progress) Reg() int {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.task++
	this.data[this.task] = 0
	return this.task
}

// Set 設定進度
func (this *Progress) Set(task, curr, total int) {
	this.lock.Lock()
	defer this.lock.Unlock()

	value := float32(0)

	switch {
	case curr <= 0, total <= 0:
		value = 0

	case curr >= total:
		value = 1

	default:
		value = float32(curr) / float32(total)
	} // switch

	this.data[task] = value
}

// Get 取得進度, 進度值為 0 ~ 1 的浮點數
func (this *Progress) Get() float32 {
	this.lock.RLock()
	defer this.lock.RUnlock()

	if this.done {
		return 1
	} // if

	total := float32(len(this.data))

	if total == 0 {
		return 0
	} // if

	curr := float32(0)

	for _, itor := range this.data {
		curr += itor
	} // for

	return curr / total
}

// Ratio 比例類型
type Ratio string

// Float32 取得浮點數
func (this Ratio) Float32() float32 {
	s := strings.TrimSpace(string(this))
	result, _ := strconv.ParseFloat(s, 32)
	return float32(result)
}

// Float64 取得浮點數
func (this Ratio) Float64() float64 {
	s := strings.TrimSpace(string(this))
	result, _ := strconv.ParseFloat(s, 64)
	return result
}

// String 取得字串
func (this Ratio) String() string {
	return string(this)
}

// Ratio 以基準值計算比例值
func (this Ratio) Ratio(base int) int {
	return int(this.ratio(base))
}

// Ratio32 以基準值計算比例值
func (this Ratio) Ratio32(base int) int32 {
	return int32(this.ratio(base))
}

// Ratio64 以基準值計算比例值
func (this Ratio) Ratio64(base int) int64 {
	return this.ratio(base)
}

// ratio 以基準值計算比例值
func (this Ratio) ratio(base int) int64 {
	if base <= 0 {
		return 0
	} // if

	s := strings.TrimSpace(string(this))
	q, ok := (&big.Rat{}).SetString(s)

	if ok == false {
		return 0
	} // if

	t := new(big.Rat).Mul(q, big.NewRat(int64(base), 1))
	return (&big.Int{}).Div(t.Num(), t.Denom()).Int64()
}

// Duration 時長類型
type Duration string

// Interval 取得時長
func (this Duration) Interval() time.Duration {
	result, _ := this.parse()
	return result
}

// String 取得字串
func (this Duration) String() string {
	return string(this)
}

// parse 解析時長
func (this Duration) parse() (result time.Duration, err error) {
	s := strings.TrimSpace(string(this))
	sign := 1

	if len(s) > 0 && s[0] == '-' {
		s = strings.TrimSpace(s[1:])
		sign = -1
	} // if

	for i := 0; i < len(s); {
		l := durationCompile.FindStringSubmatchIndex(s[i:])

		if l == nil {
			remain := strings.TrimSpace(s[i:])

			if remain == "" { // 如果剩下的都是空白, 表示解析完成
				break
			} // if

			if len(remain) > 16 { // 改一下剩下字串, 讓錯誤訊息不要太長
				remain = remain[:16] + "..."
			} // if

			return 0, fmt.Errorf("duration: invalid token %q", remain)
		} // if

		value := s[i+l[2] : i+l[3]]
		number, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return 0, fmt.Errorf("duration: invalid number %q: %w", value, err)
		} // if

		unit := strings.ToLower(s[i+l[4] : i+l[5]])

		switch unit {
		case "d":
			result += time.Duration(number) * 24 * time.Hour

		case "h":
			result += time.Duration(number) * time.Hour

		case "m":
			result += time.Duration(number) * time.Minute

		case "s":
			result += time.Duration(number) * time.Second

		case "ms":
			result += time.Duration(number) * time.Millisecond

		default:
			return 0, fmt.Errorf("duration: invalid unit %q", unit)
		} // switch

		i += l[1]
	} // for

	result = time.Duration(sign) * result
	return result, nil
}

var durationCompile = regexp.MustCompile("(?i)^\\s*(\\d+)\\s*(ms|s|m|h|d)") // 時長正則表達式

func init() {
	AddParse[int](func(value string) (result int, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return 0, nil
		} // if

		return strconv.Atoi(value)
	})
	AddParse[int32](func(value string) (result int32, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return 0, nil
		} // if

		v, err := strconv.ParseInt(value, 10, 32)

		if err != nil {
			return 0, err
		} // if

		return int32(v), nil
	})
	AddParse[int64](func(value string) (result int64, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return 0, nil
		} // if

		v, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return 0, err
		} // if

		return v, nil
	})
	AddParse[float32](func(value string) (result float32, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return 0, nil
		} // if

		v, err := strconv.ParseFloat(value, 32)

		if err != nil {
			return 0, err
		} // if

		return float32(v), nil
	})
	AddParse[float64](func(value string) (result float64, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return 0, nil
		} // if

		v, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return 0, err
		} // if

		return v, nil
	})
	AddParse[string](func(value string) (result string, err error) {
		return value, nil
	})
	AddParse[Ratio](func(value string) (result Ratio, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return "0", nil
		} // if

		if _, err = strconv.ParseFloat(value, 64); err != nil {
			return "0", err
		} // if

		return Ratio(value), nil
	})
	AddParse[Duration](func(value string) (result Duration, err error) {
		value = strings.TrimSpace(value)

		if value == "" {
			return "0s", nil
		} // if

		if _, err = Duration(value).parse(); err != nil {
			return "0s", err
		} // if

		return Duration(value), nil
	})
}
