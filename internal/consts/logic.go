package consts

const LogicTemplate = `
package {name}

import "{importPath}"

func init(){
	service.Register{Name}(&s{Name}{})
}

type s{Name} struct{}

`
