package router

import (
	"github.com/duguying/simpleci/controller"
	"github.com/duguying/simpleci/global"
)

func InitRouter() {
	global.Ma.Get("/", controller.IndexPage)
}
