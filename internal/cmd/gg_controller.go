package cmd

import (
	"context"
	"fmt"
	"gf-generate/internal/consts"
	"gf-generate/internal/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
)

var Controller = cController{}

type cController struct {
	g.Meta `name:"controller"`
}
type cControllerInput struct {
	g.Meta  `name:"controller"`
	Module  string `name:"module" arg:"true" v:"required"`
	Name    string `name:"name" arg:"true" v:"required"`
	Path    string `name:"path" short:"p" d:"/" `
	Version string `name:"version" short:"v" d:"v1"`
	Method  string `name:"method" short:"m" d:"get"`
}
type cControllerOutput struct {
}

func (receiver cController) Controller(ctx context.Context, in cControllerInput) (out *cControllerOutput, err error) {
	v := in.Version
	if gstr.IsNumeric(in.Version) {
		v = "v" + in.Version
	}

	apiPath := fmt.Sprintf("api/%s/%s", gstr.CaseSnakeFirstUpper(in.Module), v)

	if err = gfile.PutContents(apiPath+"/"+fmt.Sprintf("%s.go", gstr.CaseSnakeFirstUpper(in.Name)), gstr.ReplaceByMap(consts.ApiTemplate, g.MapStrStr{
		"{PackageName}": v,
		"{Name}":        gstr.CaseCamel(in.Name),
		"{TAG}": gstr.ReplaceByMap(consts.ControllerTag, g.MapStrStr{
			"{PATH}":   util.AddPrefixIfNotExist(in.Path, "/"),
			"{METHOD}": gstr.ToUpper(in.Method),
		}),
	})); err != nil {
		return nil, err
	}
	if err := gproc.ShellRun(ctx, "gf gen ctrl"); err != nil {
		return nil, err
	}
	g.Log().Info(ctx, "生成成功")
	return
}
