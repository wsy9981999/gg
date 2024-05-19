package consts

const ControllerTag = "`path:\"{{ .path }}\" method:\"{{ .method }}\" json:\"-\"`"
const ApiTemplate = `
package {{ .packageName }}

import "github.com/gogf/gf/v2/frame/g"

type {{ .name | CaseCamel }}Req struct {
	g.Meta {{ .tag }}
}
type {{ .name | CaseCamel }}Res struct {
	g.Meta {{ .ignoreJSON }}
}
`
