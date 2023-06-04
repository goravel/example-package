package facades

import (
	"log"

	"github.com/goravel/example-package"
	"github.com/goravel/example-package/contracts"
)

func Sms() contracts.Sms {
	instance, err := sms.App.Make(sms.Binding)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(contracts.Sms)
}
