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
}
