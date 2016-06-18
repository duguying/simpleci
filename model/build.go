package model

import (
	"github.com/duguying/simpleci/global"
	"time"
)

type Build struct {
	Id        int64
	Result    int
	UpdateLog string
	Log       string
	CreatedAt time.Time `xorm:"created"`
}

const (
	BUILD_RESULT_FINE         = 0
	BUILD_RESULT_ERROR        = 1
	BUILD_RESULT_UPDATE_FAILD = 2
	BUILD_RESULT_TIMEOUT      = 3
)

func AddBuild(result int, updateLog, log string) (int64, error) {
	b := Build{Result: result, UpdateLog: updateLog, Log: log}
	return global.Eg.Insert(&b)
}
