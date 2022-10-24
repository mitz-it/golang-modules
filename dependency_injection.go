package modules

import (
	config "github.com/mitz-it/golang-config"
	logging "github.com/mitz-it/golang-logging"
	"go.uber.org/dig"
)

type IDependencyInjectionContainer interface {
	Setup()
}

func SetupConfig(container *dig.Container, path, prefix string) {
	container.Provide(config.NewConfig)
	container.Provide(func() config.StartConfig {
		return config.StartConfig{ConfigPath: path, Prefix: prefix}
	})
}

func SetupLogging(container *dig.Container) {
	container.Provide(logging.NewLogger)
}
