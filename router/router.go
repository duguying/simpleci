package router

import (
	"github.com/duguying/simpleci/controller"
	"github.com/duguying/simpleci/global"
)

func InitRouter() {
	global.Ma.Get("/", controller.IndexPage)
	global.Ma.Get("/projects", controller.ProjectsPage)
	global.Ma.Get("/project/:project_id", controller.BuildsPage)
	global.Ma.Get("/hook/:project_id", controller.Runci)
	global.Ma.Get("/home", controller.HomeIndex)
	global.Ma.Get("/home/new_project", controller.HomeNewProject)
	global.Ma.Post("/home/create_project", controller.HomeCreateProject)
	global.Ma.Post("/home/ci_mode", controller.HomeCiMode)

	global.Ma.Get("/user/login", controller.Login)
	global.Ma.Get("/user/do_login", controller.DoLogin)
}
