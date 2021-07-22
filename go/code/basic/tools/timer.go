package tools

import (
	"fmt"
	"math/rand"
	"time"
)

const Format20060102 = "2006-01-02"

func CasualTimeCount(size int) {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(size)
	fmt.Println(count)
	time.Sleep(time.Duration(count) * time.Second)
}

// 给定一个时间，获取这个时间所在周的第一天（周一)
func GetISOWeekBegin(date time.Time) string {
	w := 0
	week := date.Weekday()
	if week == time.Sunday {
		w = 7
	} else {
		w = int(week)
	}
	d := date.AddDate(0, 0, -w+1)
	return d.Format("2006-1-2")
}

// 取给定时间的星期第一天（周一）的时间对象
func GetISOWeekBeginTime(date time.Time) time.Time {
	w := 0
	week := date.Weekday()
	if week == time.Sunday {
		w = 7
	} else {
		w = int(week)
	}
	return date.AddDate(0, 0, -w+1)
}

// 获取当前日期的星期一
func GetCurrentISOWeekBeginTime() time.Time {
	date := time.Now()
	w := 0
	week := date.Weekday()
	if week == time.Sunday {
		w = 7
	} else {
		w = int(week)
	}
	return date.AddDate(0, 0, -w+1)
}

/**
获取当天零点时间戳
*/
func GetDayZeroTimestamp() (int64, error) {
	timeStr := time.Now().Format("2006-01-02")
	d, err := time.Parse(Format20060102, timeStr)
	if err != nil {
		return 0, err
	}
	return d.Unix(), nil
}

// 返回0点日期时分
func GetDayZeroTime(t time.Time) (time.Time, error) {
	timeStr := t.Format(Format20060102)
	val, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	if err != nil {
		return t, err
	}
	return val, err
}

// 判断两个时间间隔之间
func IsCurrentWorkWay(currentWeekDay, nexWeekDay time.Time) bool {
	now := time.Now()
	return !now.Before(currentWeekDay) && now.Before(nexWeekDay)
}

// 判断时间是否一致
func IsSameTime(t1 time.Time, t2 time.Time) bool {
	return ! t1.Before(t2) && ! t1.After(t2)
}
