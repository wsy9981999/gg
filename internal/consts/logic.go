package consts

const LogicTemplate = `
package {{ .name | CaseCamelLower }}

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"{{ .servicePath }}"
)

func init() {
	service.Register{{ .name | CaseCamel }}({{ if .constructor }} news{{ .name | CaseCamel }}() {{ else }} &s{{ .name | CaseCamel }}{}  {{ end }})
}


type s{{ .name | CaseCamel }} struct {
	{{ if .cfg }}
	cfg *{{ .name | CaseCamelLower }}Config
 	{{ end }}
}
{{ if .cfg }}
type {{ .name | CaseCamelLower }}Config struct {
}
{{ end }}
{{ if .constructor }}
func news{{ .name | CaseCamel }}() *s{{ .name | CaseCamel }} {
	{{ if .cfg }}
	ctx := gctx.GetInitCtx()
	get, err := g.Cfg().Get(ctx, "{{ .name | CaseCamelLower }}")
	if err != nil {
		panic(err)
	}

	p := &s{{ .name | CaseCamel }}{}
	err = get.Scan(&p.cfg)
	if err != nil {
		panic(err)
	}
	return p
	{{ else }}
	return &s{{ .name | CaseCamel }}{}
	{{ end }}
}

{{ end }}

`
