package context

import (
	"github.com/golobby/container"
)


type ApplicationContext struct{}

func(appContext ApplicationContext) Resolve(toResolve interface{}) {
	container.Make(toResolve)
}
