# MitzIT - Go Modules

MitzIT microservices are built on top of one or multiple modules. Each module is an independent piece and has its own domain. This package allows registering modules to a single entry point and having it all started with one call.

## Usage

```go
package main

import (
	modules "github.com/mitz-it/golang-modules"
)

func main() {
	builder := modules.NewHostBuilder()
	host := builder.Build()

	host.Run()
}
```