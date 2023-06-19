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
	app.Publishes("github.com/goravel/example-package", map[string]string{
		"config/hello.go": app.ConfigPath("hello.go"),
	})
	app.Commands([]console.Command{
		commands.NewHello(),
	})
}
