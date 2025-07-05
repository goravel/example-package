package main

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/match"
	"github.com/goravel/framework/packages/modify"
	supportfile "github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/path"
)

func main() {
	config, err := GetFrameworkContent("config/hello.go")
	if err != nil {
		panic(err)
	}

	packages.Setup(os.Args).
		Install(
			modify.GoFile(path.Config("app.go")).
				Find(match.Imports()).Modify(modify.AddImport(packages.GetModulePath(), "examplepackage")).
				Find(match.Providers()).Modify(modify.Register("&examplepackage.ServiceProvider{}")),
			modify.File(path.Config("hello.go")).Overwrite(config),
		).
		Uninstall(
			modify.GoFile(path.Config("app.go")).
				Find(match.Providers()).Modify(modify.Unregister("&examplepackage.ServiceProvider{}")).
				Find(match.Imports()).Modify(modify.RemoveImport(packages.GetModulePath(), "examplepackage")),
			modify.File(path.Config("hello.go")).Remove(),
		).
		Execute()
}

func GetFrameworkContent(file string) (string, error) {
	packageName := "github.com/goravel/example-package"
	pkg, err := build.Import(packageName, "", build.FindOnly)
	if err != nil {
		return "", err
	}

	paths := strings.Split(file, "/")
	paths = append([]string{pkg.Dir}, paths...)

	return supportfile.GetContent(filepath.Join(paths...))
}
