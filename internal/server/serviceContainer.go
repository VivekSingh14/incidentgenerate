package server

import "sync"

type IServiceContainer interface{}

type Container struct{}

var (
	c             *Container
	containerOnce sync.Once
)

func ServiceContainer() *Container {
	if c == nil {
		containerOnce.Do(func() {
			c = &Container{}
		})
	}
	return c
}
