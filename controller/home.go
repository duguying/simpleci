package controller

import (
	"fmt"
	"github.com/go-macaron/macaron"
)

func HomeIndex(ctx *macaron.Context) {
	//我的项目
	//第一个项目的构建详情
	ctx.HTML(200, "home/index")
}

func HomeNewProject(ctx *macaron.Context) {
	ctx.HTML(200, "home/new_project")
}

func HomeCreateProject(ctx *macaron.Context) {
	url := ctx.Params(":url") //拿不到url
	fmt.Println(url)
	url, _ = ctx.Req.Body().String() //可以拿到
	fmt.Println(url)
	ctx.PlainText(200, []byte("success"))
}
