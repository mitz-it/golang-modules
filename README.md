# MitzIT - Go Modules

MitzIT microservices are built on top of one or multiple modules. Each module is an independent piece and has its own domain. This package allows registering modules to a single entry point and having it all started with one call.

## Modules

This package provides two kinds of modules. API modules, which will have all HTTP endpoints registered to a root router group, and Worker modules, which will be responsible for loop processing in the background.

### API modules

API modules must implement the `IApiModule` interface, and have a constructor function that receives a `*dig.Container` as parameters.

```go
package awesomeapimodule

import (
 "github.com/gin-gonic/gin"
 modules "github.com/mitz-it/golang-modules"
 "go.uber.org/dig"
)

type AwesomeApiModule struct {
 container *dig.Container
}

func (module *AwesomeApiModule) Register(group *gin.RouterGroup) {
 group.GET("/awesome-people", func(ctx *gin.Context) {
  // ...
 })
 group.POST("/awesome-people", func(ctx *gin.Context) {
  // ...
 })
}

func NewAwesomeModule(container *dig.Container) modules.IApiModule {
 return &AwesomeApiModule{
  container: container,
 }
}
```
