package model

import (
	"fmt"
	"github.com/duguying/simpleci/global"
	"time"
)

type Project struct {
	Id          int64
	Name        string
	Description string
	CiMode      int
	Crontab     string
	GitUrl      string
	WorkDir     string
	CreatedAt   time.Time `xorm:"created"`
}

func GetProject(id int64) (*Project, error) {
	p := Project{Id: id}
	has, err := global.Eg.Get(&p)
	if err != nil {
		return nil, err
	} else {
		if has {
			return &p, nil
		} else {
			return nil, fmt.Errorf("%s", "build not exist")
		}
	}
}
