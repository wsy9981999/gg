package cmd

import (
	"context"
	"gf-generate/internal/consts"
	"gf-generate/internal/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"path/filepath"
)

var Service = cService{}

type cService struct {
	g.Meta `name:"service"`
}
type cServiceInput struct {
	g.Meta      `name:"service" config:"gfcli.gen.service"`
	Name        string `name:"name" arg:"true" v:"required"`
	Constructor bool   `name:"constructor" short:"i" orphan:"true" d:"false"`
	Config      bool   `name:"cfg" short:"c" orphan:"true" d:"false"`
	SrcFolder   string `name:"srcFolder" short:"s" d:"internal/logic"`
	DstFolder   string `name:"dstFolder" short:"d" d:"internal/service"`
}
type cServiceOutput struct {
}

func (receiver *cService) Index(ctx context.Context, in cServiceInput) (out *cServiceOutput, err error) {
	if in.Config {
		in.Constructor = true
	}
	content, err := g.View().ParseContent(ctx, consts.LogicTemplate, g.MapStrAny{
		"name":        in.Name,
		"servicePath": util.GetImportPath(in.DstFolder),
		"constructor": in.Constructor,
		"cfg":         in.Config,
	})
	if err != nil {
		return nil, err
	}
	fileName := filepath.Join(in.SrcFolder, util.CaseCamelLower(in.Name), util.CaseCamelLower(in.Name)+".go")
	if err = gfile.PutContents(fileName, content); err != nil {
		return nil, err
	}
	if err = util.Run("gf gen service"); err != nil {
		return nil, err
	}
	if err = util.Format(fileName); err != nil {
		return nil, err
	}
	util.Success()
	return
}
