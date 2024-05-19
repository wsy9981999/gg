package cmd

import (
	"context"
	"fmt"
	"gf-generate/internal/consts"
	"gf-generate/internal/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"path/filepath"
	"sync"
)

var Controller = cController{}

type cController struct {
	g.Meta `name:"controller"`
}
type cControllerInput struct {
	g.Meta    `name:"controller" config:"gfcli.gen.ctrl"`
	Module    string `name:"module" arg:"true" v:"required"`
	Name      string `name:"name" arg:"true" v:"required"`
	Path      string `name:"path" short:"p" d:"/" `
	Version   string `name:"version" short:"v" d:"v1"`
	Method    string `name:"method" short:"m" d:"get"`
	SrcFolder string `name:"srcFolder" short:"s" d:"api"`
	Rest      bool   `name:"rest" short:"r" orphan:"true" d:"false"`
}
type cControllerOutput struct {
}

func (receiver cController) Controller(ctx context.Context, in cControllerInput) (out *cControllerOutput, err error) {
	v := in.Version
	if gstr.IsNumeric(in.Version) {
		v = "v" + in.Version
	}
	apiPath := filepath.Join(in.SrcFolder, in.Module, v)
	if !in.Rest {
		return nil, genController(ctx, apiPath, in, v)
	}
	prefix := []string{"GetList", "GetOne", "Create", "Update", "Delete"}
	name := util.CaseCamel(in.Name)
	wg := sync.WaitGroup{}
	for _, i2 := range prefix {
		wg.Add(1)
		in.Name = i2 + name
		go func(ctx context.Context, apiPath string, in cControllerInput, v string) {
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					g.Log().Error(ctx, err.(error).Error())
				}
			}()
			err := genController(ctx, apiPath, in, v)
			if err != nil {
				panic(err)
			}
		}(ctx, apiPath, in, v)
	}
	wg.Wait()
	if err = util.Run("gf gen ctrl"); err != nil {
		return nil, err
	}
	util.Success()
	return nil, nil
}
func genController(ctx context.Context, apiPath string, in cControllerInput, v string) error {
	g.Log().Infof(gctx.GetInitCtx(), "正在生成%s", in.Name)
	content, err := g.View().ParseContent(ctx, consts.ControllerTag, g.MapStrAny{
		"path":   util.AddPrefixIfNotExist(in.Path, "/"+gstr.ToLower(in.Module)+"/"),
		"method": in.Method,
	})
	if err != nil {
		return err
	}

	c, err := g.View().ParseContent(ctx, consts.ApiTemplate, g.MapStrAny{
		"packageName": v,
		"name":        in.Name,
		"tag":         content,
		"ignoreJSON":  "`json:\"-\"`",
	})
	if err != nil {
		return err

	}

	if err = gfile.PutContents(apiPath+"/"+fmt.Sprintf("%s.go", util.CaseSnakeFirstUpper(in.Name)), c); err != nil {
		return err

	}

	//g.Dump(in)
	return nil
}
