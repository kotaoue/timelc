package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type timeTable struct {
	start   time.Time
	end     time.Time
	meeting time.Time
}

var tt timeTable

func init() {
	s := flag.String("s", "09:30", "Start time.")
	e := flag.String("e", "18:30", "End time.")
	m := flag.String("m", "0", "Meeting minute. e.g. 30*3")
	h := flag.String("h", "0", "Meeting hour. e.g. 1*2")
	flag.Parse()

	tt.start, _ = time.Parse("15:04", *s)
	tt.end, _ = time.Parse("15:04", *e)

	tt.meeting = calcNonWorkTime(tt.meeting, *m, time.Minute)
	tt.meeting = calcNonWorkTime(tt.meeting, *h, time.Hour)
}

func main() {
	fmt.Println("今日の作業")
	fmt.Printf("予定 %s-%s\n", tt.start.Format("15:04"), tt.end.Format("15:04"))
	total := tt.end.Sub(tt.start)
	workStart := tt.start.Add((time.Duration(tt.meeting.Hour()) * time.Hour) + (time.Duration(tt.meeting.Minute()) * time.Minute))
	work := tt.end.Sub(workStart)
	fmt.Printf("業務時間 %d時間%02d分\n", (int(total.Hours()) % 24), (int(total.Minutes()) % 60))
	fmt.Printf("作業時間 %d時間%02d分\n", (int(work.Hours()) % 24), (int(work.Minutes()) % 60))
	fmt.Printf("会議時間 %d時間%02d分\n", tt.meeting.Hour(), tt.meeting.Minute())
}

func calcNonWorkTime(t time.Time, s string, d time.Duration) time.Time {
	x := 0
	y := 1
	ss := strings.Split(s, "*")

	if len(ss) > 0 {
		x, _ = strconv.Atoi(ss[0])
	}
	if len(ss) > 1 {
		y, _ = strconv.Atoi(ss[1])
	}

	return t.Add(time.Duration(x) * time.Duration(y) * d)
}
