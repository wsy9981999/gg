package util

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

func GetImportPath(filePath string) string {

	// If `filePath` does not exist, create it firstly to find the import path.
	var realPath = gfile.RealPath(filePath)
	if realPath == "" {
		_ = gfile.Mkdir(filePath)
		realPath = gfile.RealPath(filePath)
	}

	const goModName = "go.mod"
	var (
		newDir     = gfile.Dir(realPath)
		oldDir     string
		suffix     string
		goModPath  string
		importPath string
	)

	if gfile.IsDir(filePath) {
		suffix = gfile.Basename(filePath)
	}
	for {
		goModPath = gfile.Join(newDir, goModName)
		if gfile.Exists(goModPath) {
			match, _ := gregex.MatchString(`^module\s+(.+)\s*`, gfile.GetContents(goModPath))
			importPath = gstr.Trim(match[1]) + "/" + suffix
			importPath = gstr.Replace(importPath, `\`, `/`)
			importPath = gstr.TrimRight(importPath, `/`)
			return importPath
		}
		oldDir = newDir
		newDir = gfile.Dir(oldDir)
		if newDir == oldDir {
			return ""
		}
		suffix = gfile.Basename(oldDir) + "/" + suffix
	}
}
func AddPrefixIfNotExist(p string, prefix string) string {
	if gstr.HasPrefix(p, prefix) {
		return p
	}
	return prefix + p
}
