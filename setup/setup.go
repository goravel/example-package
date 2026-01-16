package main

import (
	"os"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/modify"
	supportfile "github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/path"
)

func main() {
	setup := packages.Setup(os.Args)

	// The config file will be published to the project's config directory when installing by this way.
	// You can also publish this config file manually: ./artisan vendor:publish --package=github.com/goravel/example-package
	config, err := supportfile.GetPackageContent(setup.Paths().Module().String(), "setup/config/hello.go")
	if err != nil {
		panic(err)
	}

	serviceProvider := "&hello.ServiceProvider{}"
	moduleImport := setup.Paths().Module().Import()

	setup.Install(
		modify.AddProviderApply(moduleImport, serviceProvider),

		modify.File(path.Config("hello.go")).Overwrite(config),
	).Uninstall(
		modify.File(path.Config("hello.go")).Remove(),

		modify.RemoveProviderApply(moduleImport, serviceProvider),
	).Execute()
}
