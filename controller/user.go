package controller

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func Login(c *macaron.Context) {
	c.HTML(200, "user/login")
}

func DoLogin(c *macaron.Context, sess session.Store) {
	name := c.Query("what1")
	//pass := c.Query("what2")
	sess.Set("user", name)
	c.Redirect("/home")
}