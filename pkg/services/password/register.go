package password

import (
	"go.uber.org/dig"
)

func Register(c *dig.Container) {
	c.Provide(New)
	c.Provide(func(p Passwrder) Generate { return p.Generate })
	c.Provide(func(p Passwrder) Compare { return p.Compare })
}
