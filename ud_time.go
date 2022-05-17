package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cast"
)

const OneDaySec = 24 * 60 * 60
const HourSec = 60 * 60
const MinuteSec = 60
const OneDayMills = OneDaySec * 1000

var GMTimeOffset int64
var ServerEnv string

// GetUnixNow 获取unix UTC时间
func GetUnixNow() (now int64) {
	// now = NowUTCTime().Unix()
	return NowUnixMs()
}

// NowUTCTime 获取UTC时间
func NowUTCTime() (now time.Time) {
	now = time.Now().UTC()
	if GMTimeOffset != 0 {
		now = now.Add(time.Duration(GMTimeOffset) * time.Second)
	}
	return
}

func NowUnixMs() int64 {
	return NowUTCTime().Unix()
}

func NowUnixMills() int64 {
	return NowUnixNano() / (1000 * 1000)
}

func NowUnixNano() int64 {
	return NowUTCTime().UnixNano()
}

// YYMMDDhhmmssToInt64 将{2021,9,7,4,0,0}时间格式转为UTC时间。
func YYMMDDhhmmssToInt64(strDate string) int64 {
	strTime := strings.Trim(strDate, "{}")
	tl := strings.Split(strTime, ",")

	l := len(tl)
	var h int
	if l > 3 {
		h = cast.ToInt(tl[3])
	}
	var m int
	if l > 4 {
		m = cast.ToInt(tl[4])
	}
	var s int
	if l > 5 {
		s = cast.ToInt(tl[5])
	}

	var theTime = time.Date(cast.ToInt(tl[0]),
		time.Month(cast.ToInt(tl[1])),
		cast.ToInt(tl[2]),
		h, m, s, 0, time.UTC)

	return theTime.Unix()
}

func GetTimeInteger(aTime time.Time) int {
	return aTime.Year()*10000 + int(aTime.Month())*100 + aTime.Day()
}

func GetCurrentMills(aTime time.Time) int64 {
	return aTime.UnixNano() / (1000 * 1000)
}

// CalSecondsThisDay 计算当日过去了多少秒
func CalSecondsThisDay() (seconds int64) {
	hour, minute, second := NowUTCTime().Clock()
	return int64(hour*3600 + minute*60 + second)

}

// ParseReaderTimeFormat 将24|00|00这一时间格式转为秒数,86400s 好作比较
func ParseReaderTimeFormat(timeArray []int) (seconds int) {
	return cast.ToInt(timeArray[0])*HourSec + cast.ToInt(timeArray[1])*MinuteSec + cast.ToInt(timeArray[2])
}

// WeekByDate 判断时间是当年的第几周
func WeekByDate(t time.Time) (week int) {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())
	// 今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}

	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	// 如果是周日，则算作上一周，而非新一周
	if t.Weekday() == 0 {
		week--
	}
	return

}

// GetWeekDay 返回时间的周几
func GetWeekDay(time time.Time) int {
	weekday := int(time.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return weekday
}

func GetCurrentDayByMills(curTime int64) int64 {
	return curTime / OneDayMills
}

func GetWeekStart(timeNow time.Time) int64 {
	timeNowWeekZero := timeNow.Unix() - int64(OneDaySec*timeNow.Weekday()) - int64(HourSec*timeNow.Hour()) - int64(MinuteSec*timeNow.Minute())
	return timeNowWeekZero
}

// 获取某一天的0点时间
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

//
func GetWeekdayZeroTime(t time.Time, weekday int) int64 {
	dayObj := GetZeroTime(t)
	if int(t.Weekday()) == weekday {
		return dayObj.UTC().Unix()
	}
	offset := int(int(t.Weekday()) - weekday)
	//
	offset = offset * -1

	day := dayObj.AddDate(0, 0, offset)

	return day.UTC().Unix()
}

func String2TimeInt64(strDateTime string) int64 {

	// layout := "2006-01-02 15:04:05"
	// str := "2016-07-25 11:45:26"
	t, err := time.ParseInLocation("2006-01-02 15:04:05", strDateTime, time.Local)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return t.Unix()
}
