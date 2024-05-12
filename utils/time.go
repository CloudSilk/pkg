package utils

import (
	"database/sql"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

// ParseTime 时间字符串转time.Time
func ParseTime(str string) time.Time {
	return ParseTimeWithLayout(str, TimeLayout)
}

// ParseTime2 时间字符串(结尾加上23:59:59)转time.Time
func ParseTime2(str string) time.Time {
	if len(str) == 10 {
		str += " 23:59:59"
	}
	return ParseTimeWithLayout(str, TimeLayout)
}

// FormatTime time.Time转时间字符串 格式2006-01-02 15:04:05
func FormatTime(date time.Time) string {
	return date.Format(TimeLayout)
}

// ParseDate 日期字符串转time.Time
func ParseDate(str string) time.Time {
	return ParseTimeWithLayout(str, DateLayout)
}

// FormatDate time.Time转日期字符串 格式2006-01-02
func FormatDate(date time.Time) string {
	return date.Format(DateLayout)
}

// ParseSqlNullTime 时间字符串转sql.NullTime
func ParseSqlNullTime(str string) sql.NullTime {
	return parseSqlNullTime(str, TimeLayout)
}

// FormatSqlNullTime sql.NullTime转时间字符串 格式2006-01-02 15:04:05
func FormatSqlNullTime(t sql.NullTime) string {
	return formatSqlNullTime(t, TimeLayout)
}

// ParseSqlNullDate 日期字符串转sql.NullTime
func ParseSqlNullDate(str string) sql.NullTime {
	return parseSqlNullTime(str, DateLayout)
}

// FormatSqlNullDate sql.NullTime转日期字符串 格式2006-01-02
func FormatSqlNullDate(t sql.NullTime) string {
	return formatSqlNullTime(t, DateLayout)
}

func parseSqlNullTime(str string, layout string) sql.NullTime {
	if str == "" {
		return sql.NullTime{}
	}
	date := ParseTimeWithLayout(str, layout)
	return sql.NullTime{
		Time:  date,
		Valid: true,
	}
}

func formatSqlNullTime(t sql.NullTime, layout string) string {
	if !t.Valid {
		return ""
	}
	return t.Time.Format(layout)
}

func ParseTimeWithLayout(str, layout string) time.Time {
	date, _ := time.Parse(layout, str)
	return date
}

// TimeStringToUnix 时间字符串转时间戳
func TimeStringToUnix(str string) int64 {
	if str == "" {
		return 0
	}
	return ParseTime(str).Unix()
}

// DateStringToUnix 日期字符串转时间戳
func DateStringToUnix(str string) int64 {
	if str == "" {
		return 0
	}
	return ParseDate(str).Unix()
}

// UnixToTimeString 时间戳转时间字符串 格式2006-01-02 15:04:05
func UnixToTimeString(n int64) string {
	if n <= 0 {
		return ""
	}
	return FormatTime(time.Unix(n, 0))
}

// UnixToDateString 时间戳转日期字符串 格式2006-01-02
func UnixToDateString(n int64) string {
	if n <= 0 {
		return ""
	}
	return FormatDate(time.Unix(n, 0))
}
