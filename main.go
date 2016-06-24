package main

import (
	"github.com/duguying/simpleci/global"
	"github.com/duguying/simpleci/model"
	"github.com/duguying/simpleci/queue"
	"github.com/duguying/simpleci/router"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func main() {
	global.Ma = macaron.Classic()
	global.Ma.Use(session.Sessioner())
	global.Ma.Use(macaron.Renderer())

	model.InitModel()
	router.InitRouter()
	queue.StartQueue()
	global.Ma.Run(4004)
}
