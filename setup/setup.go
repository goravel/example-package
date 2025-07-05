package main

import (
	"os"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/match"
	"github.com/goravel/framework/packages/modify"
	supportfile "github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/path"
)

func main() {
	// The config file will be published to the project's config directory when installing by this way.
	// You can also publish this config file manually: ./artisan vendor:publish --package=github.com/goravel/example-package
	config, err := supportfile.GetPackageContent(packages.GetModulePath(), "setup/config/hello.go")
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
