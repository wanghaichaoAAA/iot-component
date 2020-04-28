package utils

import (
	"strings"
	"time"
)

func TimeSubDay(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}
func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func NowStringNoSeparator() string {
	return time.Now().Format("20060102150405")
}

func GetDataTime(timeStr string) string {
	tmp1 := strings.ReplaceAll(timeStr, "-", "")
	tmp2 := strings.ReplaceAll(tmp1, ":", "")
	return strings.ReplaceAll(tmp2, " ", "")
}

func GetTimeByQN(qnStr string) (time.Time, error) {
	str := "20200318090013089"
	runes := []rune(str)
	timeLayout := "2006-01-02-15-04-05"
	y := string(runes[0:4])
	m := string(runes[4:6])
	d := string(runes[6:8])
	h := string(runes[8:10])
	mi := string(runes[10:12])
	s := string(runes[12:14])
	mark := "-"
	timeStr := y + mark + m + mark + d + mark + h + mark + mi + mark + s
	loc, _ := time.LoadLocation("Local")
	parse, err := time.ParseInLocation(timeLayout, timeStr, loc)
	if err != nil {
		utilLogger.Error("格式化成时间错误", err, qnStr)
		return time.Now(), err
	}
	return parse, nil
}

func GetPreNowTime(secondInt int) string {
	time := time.Now().Add(-time.Second * time.Duration(secondInt))
	return time.Format("2006-01-02 15:04:05")
}

//将过去的秒数转成格式化时间字符串
func PastString(seconds int) string {

	secondsDuration := time.Duration(seconds)
	startTime := time.Now().Add(time.Second * (-secondsDuration)).Format("2006-01-02 15:04:05")
	return startTime
}

func IsTime(str string) bool {
	_, err := time.Parse("20060102150405", str)
	return err == nil
}

func FormatTimeString(str string) (string, bool) {
	if str == "" {
		return "", false
	}
	times, err := time.Parse("20060102150405", str)
	if err != nil {
		return "", false
	}
	return times.Format("2006-01-02 15:04:05"), true
}

func GetTimeByString(str string) time.Time {
	time, _ := time.Parse("20060102150405", str)
	return time
}
