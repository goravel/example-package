package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

//Signature The name and signature of the console command.
func (receiver *Hello) Signature() string {
	return "hello"
}

//Description The console command description.
func (receiver *Hello) Description() string {
	return "Hello"
}

//Extend The console command extend.
func (receiver *Hello) Extend() command.Extend {
	return command.Extend{}
}

//Handle Execute the console command.
func (receiver *Hello) Handle(ctx console.Context) error {
	fmt.Println("Run Hello command")

	return nil
}
