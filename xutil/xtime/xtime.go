package xtime

import (
	"time"
)

const (
	DefaultDate     = "2006-01-02"
	DefaultTime     = "15:04:05"
	DefaultDateTime = "2006-01-02 15:04:05"
)

type XTime time.Time

func Now() int64 {
	return time.Now().UnixNano() / 1e6
}

func TodayDateStr() string {
	return time.Now().Format(DefaultDate)
}

func TodayTimeStr() string {
	return time.Now().Format(DefaultTime)
}

func TodayDateTimeStr() string {
	return time.Now().Format(DefaultDateTime)
}

func TodayStart() int64 {
	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	return date.UnixNano() / 1e6
}

func Time(days int) XTime {
	date := time.Now()
	return XTime(time.Date(date.Year(), date.Month(), date.Day()+days, 0, 0, 0, 0, date.Location()))
}

func (xtime XTime) DayStart() int64 {
	t := time.Time(xtime)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixNano() / 1e6
}

func (xtime XTime) DateStr() string {
	return time.Time(xtime).Format(DefaultDate)
}

func (xtime XTime) TimeStr() string {
	return time.Time(xtime).Format(DefaultTime)
}

func (xtime XTime) DateTimeStr() string {
	return time.Time(xtime).Format(DefaultDateTime)
}
