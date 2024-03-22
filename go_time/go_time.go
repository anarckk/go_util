package go_time

import (
	"strconv"
	"time"
)

// CurrentDatetimeStr 输出标准的当前时间戳字符串 "yyyy-MM-dd HH:mm:ss"
func CurrentDatetimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurrentDateStr 输出标准的当前日期字符串 "yyyy-MM-dd"
func CurrentDateStr() string {
	return time.Now().Format("2006-01-02")
}

// CurrentTimeMillis 获取当前时间戳秒数 1703555599
func CurrentTimeSecond() int64 {
	return time.Now().Unix()
}

// CurrentTimeMillis 获取当前时间戳毫秒数 1686273835040
func CurrentTimeMillis() int64 {
	return time.Now().UnixMilli()
}

// CurrentTimeMillisStr 获取当前时间戳毫秒数 1686273835040
func CurrentTimeMillisStr() string {
	return strconv.Itoa(int(CurrentTimeMillis()))
}

// FormatTime 时间格式化字符串
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FromMilliSecond 从时间戳毫秒数转成对象
func FromMilliSecond(ms int64) time.Time {
	return time.Unix(ms/1000, (ms%1000)*int64(time.Millisecond))
}

// GetNowUTC 返回当前时间的UTC时间，会比东八区少8小时，也就是说这是零区的时间
func GetNowUTC() time.Time {
	return time.Now().UTC()
}

// ConvertUTCToLocal 把UTC时间转为当地时间
// 这个api可以返回东八区时间
func ConvertUTCToLocal(utc time.Time) time.Time {
	return utc.Local()
}
