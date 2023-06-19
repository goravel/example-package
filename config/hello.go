package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("hello", map[string]any{
		"name": "Package",
	})
}
