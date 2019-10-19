package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Date(format string, timestamp ...int64) string {
	var ts = time.Now().Unix()
	if len(timestamp) > 0 {
		ts = timestamp[0]
	}
	var t = time.Unix(ts, 0)
	Y := strconv.Itoa(t.Year())
	m := fmt.Sprintf("%02d", t.Month())
	d := fmt.Sprintf("%02d", t.Day())
	H := fmt.Sprintf("%02d", t.Hour())
	i := fmt.Sprintf("%02d", t.Minute())
	s := fmt.Sprintf("%02d", t.Second())

	format = strings.Replace(format, "Y", Y, -1)
	format = strings.Replace(format, "m", m, -1)
	format = strings.Replace(format, "d", d, -1)
	format = strings.Replace(format, "H", H, -1)
	format = strings.Replace(format, "i", i, -1)
	format = strings.Replace(format, "s", s, -1)
	return format
}

func GetNowTimeStamp() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func GetNowDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func Get7DayBeforeDate() (start string, end string) {
	nowTime := time.Now()
	year := nowTime.Year()
	month := nowTime.Month()
	day := nowTime.Day()
	loc := nowTime.Location()
	tmpEndTime := time.Date(year, month, day, 23, 59, 59, 0, loc)
	endTime := tmpEndTime.AddDate(0, 0, -1)

	tmpStartTime := time.Date(year, month, day, 0, 0, 0, 0, loc)
	startTime := tmpStartTime.AddDate(0, 0, -7)

	return startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")

}

func FormatTimestamp(timeStr string) int64 {
	tm, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return tm.Unix()
}
