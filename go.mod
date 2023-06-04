module github.com/goravel/example-package

go 1.18

require github.com/goravel/framework v1.11.2 // indirect

replace (
	github.com/goravel/framework v1.11.2 => ../goravel/framework
)
