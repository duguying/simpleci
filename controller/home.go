package controller

import (
	"github.com/go-macaron/macaron"
)

func HomeIndex(ctx *macaron.Context) {
	//我的项目
	//第一个项目的构建详情
	ctx.HTML(200, "home/index")
}
