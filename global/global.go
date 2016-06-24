package global

import (
	"github.com/go-xorm/xorm"
	"github.com/gogits/cron"
	"gopkg.in/macaron.v1"
)

var (
	Ma *macaron.Macaron
	Cr *cron.Cron
	Eg *xorm.Engine
	// Lc bool = false
)
