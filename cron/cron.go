package cron

import (
	"github.com/duguying/simpleci/global"
	"github.com/gogits/cron"
)

func StartCron() {
	global.Cr = cron.New()
	global.Cr.AddFunc("0 30 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})
}
