package controller

import (
	"github.com/duguying/simpleci/runner"
	"gopkg.in/macaron.v1"
)

func IndexPage(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}

func ProjectsPage(ctx *macaron.Context) {
	projects := []map[string]string{
		map[string]string{
			"id":          "1",
			"name":        "name",
			"description": "hello test",
			"mode":        "1",
			"create_at":   "time",
		},
	}
	ctx.Data["projects"] = projects
	ctx.HTML(200, "project_list")
}

func BuildsPage(ctx *macaron.Context) {
	projectId := ctx.Params("project_id")
	ctx.Data["project_id"] = projectId
	ctx.HTML(200, "build_list")
}

func BuildDetailPage(ctx *macaron.Context) {
	ctx.HTML(200, "build_detail")
}

func Runci(ctx *macaron.Context) {
	projectId := ctx.Params("project_id")
	info := map[string]interface{}{
		"commit_id": "1234",
	}
	result := runner.AddRunTask(projectId, info)
	ctx.JSON(200, result)
}
