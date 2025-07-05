package hello

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"

	"github.com/goravel/example-package/commands"
)

const Binding = "hello"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewHello(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	// The config file will be published to the project's config directory when installing by setup/setup.go.
	// You can also publish this config file manually, uncomment the following code and run ./artisan vendor:publish --package=github.com/goravel/example-package

	// app.Publishes("github.com/goravel/example-package", map[string]string{
	// 	"config/hello.go": app.ConfigPath("hello.go"),
	// })

	app.Commands([]console.Command{
		commands.NewHello(),
	})
}
