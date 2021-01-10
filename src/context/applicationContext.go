package context

import (
	"github.com/golobby/container"
)


type ApplicationContext struct{}

func(appContext ApplicationContext) Bind(bindFunction interface{}) {
	container.Singleton(bindFunction)
}

func(appContext ApplicationContext) Resolve(toResolve interface{}) {
	container.Make(toResolve)
}

func(appContext ApplicationContext) Reset() {
	container.Reset()
}
