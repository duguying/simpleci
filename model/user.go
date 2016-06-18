package model

import (
	"github.com/duguying/simpleci/global"
	"github.com/gogather/com"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Password  string
	Salt      string
	Role      int
	CreatedAt time.Time `xorm:"created"`
}

// add a user
func AddUser(name, password string, role int) {
	salt := com.RandString(7)
	password = com.Md5(password + salt)
	u := User{Name: name, Password: password, Salt: salt, Role: role}
	global.Eg.Insert(&u)
}

// add an admin
func AddAdmin(name, password string) {
	AddUser(name, password, 1)
}

// add an project owner
func AddOwner(name, password string) {
	AddUser(name, password, 0)
}
