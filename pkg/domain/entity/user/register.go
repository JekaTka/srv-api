package user

import "go.uber.org/dig"

func Register(c *dig.Container) {
	c.Provide(NewPGRepoRepository)
	c.Provide(NewService)
}
