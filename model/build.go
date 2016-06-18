package model

import (
	"time"
)

type Build struct {
	Id        int64
	Result    int
	Log       int64
	CreatedAt time.Time `xorm:"created"`
}
