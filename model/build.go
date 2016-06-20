package model

import (
	"fmt"
	"github.com/duguying/simpleci/global"
	"time"
)

type Build struct {
	Id        int64
	ProjectId int64
	CommitId  string
	Status    int
	Result    int
	UpdateLog string
	Log       string
	CreatedAt time.Time `xorm:"created"`
}

const (
	BUILD_RESULT_PENDING      = 0
	BUILD_RESULT_FINE         = 1
	BUILD_RESULT_ERROR        = 2
	BUILD_RESULT_UPDATE_FAILD = 3
	BUILD_RESULT_TIMEOUT      = 4
)

const (
	QUEUE_STATUS_WAITING  = 1
	QUEUE_STATUS_PENDING  = 2
	QUEUE_STATUS_FINISHED = 3
)

func AddBuild(result int, commitId, updateLog, log string) (int64, error) {
	b := Build{Result: result, CommitId: commitId, UpdateLog: updateLog, Log: log}
	return global.Eg.Insert(&b)
}

func GetBuildQueueTop() (*Build, error) {
	b := Build{}
	has, err := global.Eg.Where("status = ?", QUEUE_STATUS_WAITING).Asc("id").Get(&b)

	if err != nil {
		return nil, err
	} else {
		if !has {
			return nil, fmt.Errorf("%s", "not exist")
		} else {
			return &b, nil
		}
	}
}

func GetBuildPending() (*Build, error) {
	b := Build{}
	has, err := global.Eg.Where("status = ?", QUEUE_STATUS_PENDING).Asc("id").Get(&b)
	if err != nil {
		return nil, err
	} else {
		if !has {
			return nil, fmt.Errorf("%s", "not exist")
		} else {
			return &b, nil
		}
	}
}

func UpdateBuild(id int64, result int, updateLog, log string) (int64, error) {
	build := Build{Result: result, UpdateLog: updateLog, Log: log}
	return global.Eg.Id(id).Update(&build)
}

func GetBuild(id int64) (*Build, error) {
	b := Build{Id: id}
	has, err := global.Eg.Get(&b)
	if err != nil {
		return nil, err
	} else {
		if has {
			return &b, nil
		} else {
			return nil, fmt.Errorf("%s", "build not exist")
		}
	}

}
