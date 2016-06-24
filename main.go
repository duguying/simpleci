package main

import (
	"github.com/duguying/simpleci/global"
	"github.com/duguying/simpleci/model"
	"github.com/duguying/simpleci/queue"
	"github.com/duguying/simpleci/router"
	"github.com/go-macaron/macaron"
	"github.com/go-macaron/session"
)

func main() {
	global.Ma = macaron.Classic()
	global.Ma.Use(macaron.Renderer())
	global.Ma.Use(session.Sessioner())
	model.InitModel()
	router.InitRouter()
	queue.StartQueue()
	global.Ma.Run(4004)
}
