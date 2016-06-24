package controller

import (
	"github.com/duguying/simpleci/model"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func Login(c *macaron.Context) {
	c.HTML(200, "user/login")
}

func DoLogin(c *macaron.Context, sess session.Store) {
	name := c.Query("what1")
	pass := c.Query("what2")
	ok, user := model.UserLoginCheck(name, pass)
	resp := new(respObj)
	resp.Success = false
	if ok {
		sess.Set("username", user.Name)
		sess.Set("userid", user.Id)
		resp.Success = true
	}
	c.JSON(200, resp)
}

func UserRegisterPage(c *macaron.Context) {
	c.HTML(200, "user/register")
}

func DoReg(c *macaron.Context, sess session.Store) {
	name := c.Query("what1")
	pass := c.Query("what2")
	userid := model.AddOwner(name, pass)
	sess.Set("username", name)
	sess.Set("userid", userid)
	resp := new(respObj)
	resp.Success = true
	c.JSON(200, resp)
}

func Logout(c *macaron.Context, sess session.Store) {
	sess.Destory(c)
	c.Redirect("/user/login")
}
