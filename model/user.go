package model

import (
	"fmt"
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
func AddUser(name, password string, role int) int64 {
	salt := com.RandString(7)
	password = com.Md5(password + salt)
	u := User{Name: name, Password: password, Salt: salt, Role: role}
	global.Eg.Insert(&u)
	return u.Id
}

// add an admin
func AddAdmin(name, password string) int64 {
	return AddUser(name, password, 1)
}

// add an project owner
func AddOwner(name, password string) int64 {
	return AddUser(name, password, 0)
}

func UserLoginCheck(name, pass string) (bool, *User) {
	user := new(User)
	_, err := global.Eg.Where("name = ?", name).Get(user)
	logPass := com.Md5(pass + user.Salt)
	if err != nil {
		fmt.Println(err)
	}
	return logPass == user.Password, user
}
