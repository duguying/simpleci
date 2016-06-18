package controller

import (
	"github.com/go-macaron/macaron"
)

func IndexPage(ctx *macaron.Context) {
	ctx.Data["Name"] = "jeremy"
	ctx.HTML(200, "hello")
}
