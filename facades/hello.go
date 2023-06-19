package facades

import (
	"log"

	hello "github.com/goravel/example-package"
	"github.com/goravel/example-package/contracts"
)

func Hello() contracts.Hello {
	instance, err := hello.App.Make(hello.Binding)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(contracts.Hello)
}
