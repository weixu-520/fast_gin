package cron_ser

import (
	"fmt"
	"time"
)

func Func1() {
	fmt.Println(time.Now().Format(("2005-01-02 15:04:05")))
}
func cronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))

	// crontab.AddFunc("* * * * * *", Func1)
	crontab.Start()
}
