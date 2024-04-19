package cmd

import (
	"context"
	"gf-generate/internal/consts"
	"gf-generate/internal/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
)

var Service = cService{}

type cService struct {
	g.Meta `name:"service"`
}
type cServiceInput struct {
	g.Meta `name:"service"`
	Name   string `name:"name" arg:"true" v:"required"`
}
type cServiceOutput struct {
}

func (receiver *cService) Index(ctx context.Context, in cServiceInput) (out *cServiceOutput, err error) {
	snakeStr := gstr.CaseSnakeFirstUpper(in.Name)
	camelStr := gstr.CaseCamel(in.Name)
	name := "internal/logic/" + snakeStr + "/" + snakeStr + ".go"
	gfile.PutContents(name, gstr.ReplaceByMap(consts.LogicTemplate, g.MapStrStr{
		"{name}":       snakeStr,
		"{Name}":       camelStr,
		"{importPath}": gstr.CaseSnake(util.GetImportPath("internal/service")),
	}))
	if err := gproc.ShellRun(ctx, "gf gen service"); err != nil {
		return nil, err
	}
	return
}
