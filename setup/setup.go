package main

import (
	"os"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/modify"
	"github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/path"
)

func main() {
	setup := packages.Setup(os.Args)

	// The config file will be published to the project's config directory automatically when installing by this way.
	// You can also publish this config file manually: ./artisan vendor:publish --package=github.com/goravel/example-package
	config, err := file.GetPackageContent(setup.Paths().Module().String(), "setup/config/hello.go")
	if err != nil {
		panic(err)
	}

	serviceProvider := "&hello.ServiceProvider{}"
	moduleImport := setup.Paths().Module().Import()

	setup.Install(
		// Add the service provider to the providers slice in bootstrap/providers.go
		modify.AddProviderApply(moduleImport, serviceProvider),

		// Add the config file to the config directory
		modify.File(path.Config("hello.go")).Overwrite(config),
	).Uninstall(
		// Remove the config file from the config directory
		modify.File(path.Config("hello.go")).Remove(),

		// Remove the service provider from the providers slice in bootstrap/providers.go
		modify.RemoveProviderApply(moduleImport, serviceProvider),
	).Execute()
}
