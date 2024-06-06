package cmd

import (
	"context"
	"fmt"
	"gf-generate/internal/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gzuidhof/tygo/tygo"
	"path/filepath"
)

var Ts = cTs{}

type cTs struct {
	g.Meta `name:"ts" `
}

type cTsInput struct {
	g.Meta     `name:"ts" config:"gfcli.gen.ts"`
	Src        string `name:"src" short:"s" d:"api"`
	Dist       string `name:"dist" short:"d" d:"frontend"`
	ModuleName string `name:"moduleName" short:"m"`
}
type cTsOutput struct {
}

func (receiver cTs) Index(ctx context.Context, in cTsInput) (*cTsOutput, error) {
	var path string
	//
	if in.ModuleName == "" {
		path = util.GetImportPath(in.Src)
	} else {
		path = gfile.Join(in.ModuleName)
	}
	dir, err := gfile.ScanDir(in.Src, "*", false)
	if err != nil {
		return nil, err
	}
	for _, d := range dir {
		m := gfile.Name(d)
		glob, err := gfile.ScanDirFunc(d, "*", true, func(path string) string {
			if gstr.HasSuffix(path, ".go") {
				return ""
			}
			return path
		})
		if err != nil {
			return nil, err
		}
		for _, p := range glob {
			rel, err := filepath.Rel(gfile.Join(gfile.Pwd(), in.Src), p)
			if err != nil {
				return nil, err
			}
			err = tygo.New(&tygo.Config{
				Packages: []*tygo.PackageConfig{
					{
						Path:         gfile.Join(path, rel),
						OutputPath:   fmt.Sprintf("frontend/%s.ts", m),
						FallbackType: "any",
					},
				},
			}).Generate()

		}

	}
	//g.Dump(in)

	//if err != nil {
	//	panic(err)
	//}
	return nil, nil
}
