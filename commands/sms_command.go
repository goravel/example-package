package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type SmsCommand struct{}

func NewSmsCommand() *SmsCommand {
	return &SmsCommand{}
}

//Signature The name and signature of the console command.
func (receiver *SmsCommand) Signature() string {
	return "sms"
}

//Description The console command description.
func (receiver *SmsCommand) Description() string {
	return "Sms"
}

//Extend The console command extend.
func (receiver *SmsCommand) Extend() command.Extend {
	return command.Extend{}
}

//Handle Execute the console command.
func (receiver *SmsCommand) Handle(ctx console.Context) error {
	fmt.Println("Run sms command")

	return nil
}
