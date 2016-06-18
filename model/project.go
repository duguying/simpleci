package model

import (
	"time"
)

type Project struct {
	Id          int64
	Name        string
	Description string
	CiMode      int
	Crontab     string
	CreatedAt   time.Time `xorm:"created"`
}