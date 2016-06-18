package global

import (
	"github.com/go-macaron/macaron"
	"github.com/go-xorm/xorm"
	"github.com/gogits/cron"
)

var (
	Ma *macaron.Macaron
	Cr *cron.Cron
	En *xorm.Engine
)
