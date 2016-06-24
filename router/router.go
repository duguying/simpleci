package router

import (
	"github.com/duguying/simpleci/controller"
	"github.com/duguying/simpleci/global"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func addSession(ctx *macaron.Context, sess session.Store) {
	ctx.Data["username"] = sess.Get("username")
	ctx.Data["userid"] = sess.Get("userid")
}

func InitRouter() {
	global.Ma.Get("/", controller.IndexPage)
	global.Ma.Get("/projects", controller.ProjectsPage)
	global.Ma.Get("/project/:project_id", controller.BuildsPage)
	global.Ma.Get("/hook/:project_id", controller.Runci)

	global.Ma.Get("/home", addSession, controller.HomeIndex)
	global.Ma.Get("/home/new_project", controller.HomeNewProject)
	global.Ma.Post("/home/create_project", controller.HomeCreateProject)
	global.Ma.Post("/home/ci_mode", controller.HomeCiMode)

	global.Ma.Get("/user/login", controller.Login)
	global.Ma.Post("/user/do_login", controller.DoLogin)
	global.Ma.Get("/user/register", controller.UserRegisterPage)
	global.Ma.Post("/user/do_reg", controller.DoReg)
	global.Ma.Get("/user/logout", controller.Logout)
}
