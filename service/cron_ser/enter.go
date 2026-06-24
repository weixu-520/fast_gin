package cron_ser

import (
	"fmt"
	"time"
)

func Func1() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func CronInit() {
	// TODO: 后续集成 cron 定时任务功能
	fmt.Println("cron 定时任务模块待实现")
}
