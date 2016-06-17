package cron

import (
	"github.com/gogits/cron"
	"github.com/duguying/simpleci/global"
)

func StartCron() {
	global.Cr := cron.New()
	global.Cr.AddFunc("0 30 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})
}
