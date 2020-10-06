package authv1

import "go.uber.org/dig"

func Register(c *dig.Container) {
	c.Provide(makeRegistrationHandler)
}
