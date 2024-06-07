package frontend

import (
	"bytes"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Type string

const (
	Resolver Type = "resolver"
	Imports  Type = "imports"
)

type subItem struct {
	Type    Type
	Import  string
	Content string
	Named   bool
}

type autoImport struct {
	subItems []subItem
}

func newAutoImport() *autoImport {

	return &autoImport{
		subItems: make([]subItem, 0),
	}
}

var imports = gmap.New()

func (receiver *autoImport) Handle(plugins []string) {
	for _, p := range plugins {
		if imports.Contains(p) {
			continue
		}
		item := imports.Get(p).(subItem)
		receiver.subItems = append(receiver.subItems, item)
	}
}
func (receiver *autoImport) Render() {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("import autoImport from 'unplugin-auto-import/vite'\n")
	for _, item := range receiver.subItems {
		if item.Import != "" {
			buffer.WriteString(item.Import + "\n")
		}
	}
	json := gjson.New(nil)
	json.MustSet("dts", "types/auto-imports.d.ts")
	json.MustSet("dirs", []string{"src/api/**", "src/stores/**", "src/hook/**"})
	json.MustSet("eslintrc.enabled", true)
	r := 0
	im := 0
	for i, item := range receiver.subItems {
		if item.Type == Resolver {
			json.MustSet("eslintrc.enabled", true)

			r++
		}

	}

	//		 dts: 'types/auto-imports.d.ts',
	//  dirs:['src/api/**','','src/hook/**'],
	//  eslintrc: {
	//    enabled: true
	//  }
}
