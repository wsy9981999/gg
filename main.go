package main

import (
	"gf-generate/internal/cmd"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.GetInitCtx()
	object, err := gcmd.NewFromObject(cmd.Gg)
	if err != nil {
		g.Log().Fatal(ctx, err.Error())
		return
	}
	err = object.AddObject(cmd.Controller, cmd.Service)
	if err != nil {
		g.Log().Fatal(ctx, err.Error())

		return
	}
	object.Run(ctx)

}
