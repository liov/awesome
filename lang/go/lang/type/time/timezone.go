package main

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func main() {
	beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2024-06-20 00:00:00", time.UTC)
	fmt.Println(beginTime.UnixMilli())
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2024-06-21 00:00:00", time.UTC)
	fmt.Println(endTime.UnixMilli())
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		tz = time.Local
	}
	beginTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2024-06-20 00:00:00", tz)
	fmt.Println(beginTime.UnixMilli())
	endTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2024-06-21 00:00:00", tz)
	fmt.Println(endTime.UnixMilli())
}
