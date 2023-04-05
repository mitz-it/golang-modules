package modules

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Module struct {
	name        string
	workers     []IWorker
	controllers []IController
	initCalls   []InitCall
	container   *dig.Container
}

func (module *Module) appendControler(controller IController) {
	module.controllers = append(module.controllers, controller)
}

func (module *Module) appendWorker(worker IWorker) {
	module.workers = append(module.workers, worker)
}

func (module *Module) appendInitCall(initCall InitCall) {
	module.initCalls = append(module.initCalls, initCall)
}

func (module *Module) registerControllers(api *API) {
	var group *gin.RouterGroup

	if module.name != "" {
		group = api.rootGroup.Group(module.name)
	} else {
		group = api.rootGroup
	}

	for _, controller := range module.controllers {
		controller.Register(group)
	}
}

func (module *Module) startWorkers() {
	for _, worker := range module.workers {
		go worker.Run()
	}
}

func (module *Module) invokeInitCalls() {
	for _, initCall := range module.initCalls {
		initCall(module.container)
	}
}

func newModule(name string, container *dig.Container) *Module {
	return &Module{
		name:        name,
		container:   container,
		workers:     make([]IWorker, 0),
		controllers: make([]IController, 0),
		initCalls:   make([]InitCall, 0),
	}
}
