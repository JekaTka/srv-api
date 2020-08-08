package db

import (
	"go.uber.org/dig"
)

func Register(c *dig.Container) {
	c.Provide(NewConnection)
}
