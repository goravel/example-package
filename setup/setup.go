package main

import (
	"embed"
	"io/fs"
	"os"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/match"
	"github.com/goravel/framework/packages/modify"
	"github.com/goravel/framework/support/path"
)

//go:embed config/hello.go
var files embed.FS

func main() {
	config, err := fs.ReadFile(files, "config/hello.go")
	if err != nil {
		panic(err)
	}

	packages.Setup(os.Args).
		Install(
			modify.GoFile(path.Config("app.go")).
				Find(match.Imports()).Modify(modify.AddImport(packages.GetModulePath(), "examplepackage")).
				Find(match.Providers()).Modify(modify.Register("&examplepackage.ServiceProvider{}")),
			modify.File(path.Config("hello.go")).Overwrite(string(config)),
		).
		Uninstall(
			modify.GoFile(path.Config("app.go")).
				Find(match.Providers()).Modify(modify.Unregister("&examplepackage.ServiceProvider{}")).
				Find(match.Imports()).Modify(modify.RemoveImport(packages.GetModulePath(), "examplepackage")),
			modify.File(path.Config("hello.go")).Remove(),
		).
		Execute()
}
