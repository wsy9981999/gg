package consts

const ControllerTag = "`path:\"{PATH}\" method:\"{METHOD}\"` "
const ApiTemplate = `
package {PackageName}

import "github.com/gogf/gf/v2/frame/g"

type {Name}Req struct {
	g.Meta {TAG}
}
type {Name}Res struct {
}
`
