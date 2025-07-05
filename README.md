# An Example For A Goravel Extend Package

## Install

Run the command below in your project to install the package automatically:

```bash
./artisan package:install github.com/goravel/example-package
``` 

Or check [the setup file](./setup/setup.go) to install the package manually.

## Testing

```
// main.go
import examplefacades "github.com/goravel/example-package/facades"

fmt.Println(examplefacades.Hello().World())
```

The console will print `Welcome To Goravel Package`.
