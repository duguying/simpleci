package controller

import (
	"github.com/duguying/simpleci/model"
	"github.com/go-macaron/macaron"
	"strings"
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
	url := ctx.Query("url")
	lastSlash := strings.LastIndexByte(url, '/')
	urlParas := strings.Split(url, "/")
	name := ""
	if lastSlash == len(url)-1 {
		name = urlParas[len(urlParas)-2]
	} else {
		name = urlParas[len(urlParas)-1]
	}
	_, err := model.SaveProject(name, url)
	ctx.Data["success"] = true
	if err != nil {
		ctx.Data["success"] = false
		ctx.Data["errMsg"] = err
	}
	ctx.HTML(200, "home/new_p_2")
}
