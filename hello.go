package hello

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
)

type Hello struct {
	config config.Config
}

func NewHello(config config.Config) *Hello {
	return &Hello{config: config}
}

func (s *Hello) World() string {
	return fmt.Sprintf("Welcome To Goravel %s", s.config.GetString("hello.name"))
}
