package main

import (
	_ "gf-generate/internal/logic"

	"gf-generate/internal/cmd"
	"gf-generate/internal/util"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	ctx := gctx.GetInitCtx()

	if path, _ := gfile.Search("hack"); path != "" {
		if adapter, ok := g.Cfg().GetAdapter().(*gcfg.AdapterFile); ok {
			if err := adapter.SetPath(path); err != nil {
				g.Log().Fatal(ctx, err)
			}
		}
	}

	g.View().BindFuncMap(g.MapStrAny{
		"CaseSnakeFirstUpper": util.CaseSnakeFirstUpper,
		"CaseCamel":           util.CaseCamel,
		"CaseCamelLower":      util.CaseCamelLower,
	})
	object, err := gcmd.NewFromObject(cmd.Gg)
	if err != nil {
		g.Log().Fatal(ctx, err.Error())
		return
	}
	err = object.AddObject(cmd.Controller, cmd.Service, cmd.Ts)
	if err != nil {
		g.Log().Fatal(ctx, err.Error())

		return
	}
	object.Run(ctx)
	//buffer := bytes.NewBuffer(nil)
	//buffer.WriteString("123")
	//fmt.Println(buffer.ReadString('\n'))
	//fmt.Println(buffer.ReadString('\n'))
}
