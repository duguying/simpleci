package main

import (
	"github.com/duguying/simpleci/global"
	"github.com/duguying/simpleci/router"
	"github.com/go-macaron/macaron"
)

func main() {
	global.Ma = macaron.Classic()
	router.InitRouter()
	global.Ma.Run()
}
