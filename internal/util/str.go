package util

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/text/gstr"
)

var m = gmap.New(false)

const (
	CaseSnakeFirstUpperFunc = "CaseSnakeFirstUpper"
	CaseCamelFunc           = "CaseCamel"
	CaseCamelLowerFunc      = "CaseCamelLower"
)

func CaseSnakeFirstUpper(str string) string {
	cacheName := str + "|" + CaseSnakeFirstUpperFunc
	m.SetIfNotExistFunc(cacheName, func() interface{} {
		return gstr.CaseSnakeFirstUpper(str)
	})
	return m.Get(cacheName).(string)
}

func CaseCamel(str string) string {
	cacheName := str + "|" + CaseCamelFunc
	m.SetIfNotExistFunc(cacheName, func() interface{} {
		return gstr.CaseCamel(str)
	})
	return m.Get(cacheName).(string)
}
func CaseCamelLower(str string) string {
	cacheName := str + "|" + CaseCamelLowerFunc
	m.SetIfNotExistFunc(cacheName, func() interface{} {
		return gstr.CaseCamelLower(str)
	})
	return m.Get(cacheName).(string)
}
