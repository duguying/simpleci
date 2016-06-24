package controller

import (
	"github.com/duguying/simpleci/model"
	"github.com/go-macaron/macaron"
	"log"
	"strconv"
	"strings"
)

type respObj struct {
	Success bool
	Msg     string
	Obj     interface{}
}

func HomeIndex(ctx *macaron.Context) {
	//我的项目
	//第一个项目的构建详情
	ctx.HTML(200, "home/index")
}

func HomeNewProject(ctx *macaron.Context) {
	ctx.HTML(200, "home/new_project")
}

func HomeCreateProject(ctx *macaron.Context, logger *log.Logger) {
	url := ctx.Query("url")
	lastSlash := strings.LastIndexByte(url, '/')
	urlParas := strings.Split(url, "/")
	name := ""
	if lastSlash == len(url)-1 {
		name = urlParas[len(urlParas)-2]
	} else {
		name = urlParas[len(urlParas)-1]
	}
	id, err := model.SaveProject(name, url)
	ctx.Data["success"] = true
	logger.Println(id)
	if err != nil {
		ctx.Data["success"] = false
		ctx.Data["errMsg"] = err
	} else {
		ctx.Data["projectID"] = id
		ctx.Data["host"] = ctx.Req.Host
	}
	ctx.HTML(200, "home/new_p_2")
}

func HomeCiMode(ctx *macaron.Context, logger *log.Logger) {
	cimode := ctx.Query("cimode")
	time := ctx.Query("time")
	projectID := ctx.Query("projectid")
	logger.Println(cimode)

	resp := new(respObj)
	resp.Success = true
	project := new(model.Project)
	var err error
	project.CiMode, err = strconv.Atoi(cimode)
	if err != nil {
		resp.Success = false
		resp.Msg = "请选择正确的ci模式"
	}
	project.Crontab = time
	projectidInt, err := strconv.Atoi(projectID)
	if err != nil {
		resp.Success = false
		resp.Msg = "请选择正确的项目"
	}
	err = model.UpdateProject(int64(projectidInt), project)
	if err != nil {
		logger.Println(err)
	}
	ctx.JSON(200, resp)
}
